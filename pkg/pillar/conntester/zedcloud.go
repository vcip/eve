// Copyright (c) 2022 Zededa, Inc.
// SPDX-License-Identifier: Apache-2.0

package conntester

import (
	"crypto/tls"
	"fmt"
	"io/ioutil"
	"strings"
	"time"

	"github.com/lf-edge/eve/pkg/pillar/base"
	"github.com/lf-edge/eve/pkg/pillar/devicenetwork"
	"github.com/lf-edge/eve/pkg/pillar/hardware"
	"github.com/lf-edge/eve/pkg/pillar/types"
	"github.com/lf-edge/eve/pkg/pillar/zedcloud"
	uuid "github.com/satori/go.uuid"
)

// Hard-coded at 1 for now; at least one interface needs to work.
const requiredSuccessCount uint = 1

var nilUUID = uuid.UUID{} // used as a constant

// ZedcloudConnectivityTester implements external connectivity testing using
// the "/api/v2/edgeDevice/ping" endpoint provided by the zedcloud.
type ZedcloudConnectivityTester struct {
	// Exported attributes below should be injected.
	Log         *base.LogObject
	AgentName   string
	TestTimeout time.Duration // can be changed in run-time
	Metrics     *zedcloud.AgentMetrics

	iteration     int
	prevTLSConfig *tls.Config
}

// TestConnectivity uses VerifyAllIntf from the zedcloud package, which
// tries to call the "ping" API of the controller.
func (t *ZedcloudConnectivityTester) TestConnectivity(
	dns types.DeviceNetworkStatus) (types.IntfStatusMap, error) {

	t.iteration++
	intfStatusMap := *types.NewIntfStatusMap()
	t.Log.Tracef("TestConnectivity() requiredSuccessCount %d, iteration %d",
		requiredSuccessCount, t.iteration)

	server, err := ioutil.ReadFile(types.ServerFileName)
	if err != nil {
		t.Log.Fatal(err)
	}
	serverNameAndPort := strings.TrimSpace(string(server))
	serverName := strings.Split(serverNameAndPort, ":")[0]

	zedcloudCtx := zedcloud.NewContext(t.Log, zedcloud.ContextOptions{
		DevNetworkStatus: &dns,
		Timeout:          uint32(t.TestTimeout.Seconds()),
		AgentMetrics:     t.Metrics,
		Serial:           hardware.GetProductSerial(t.Log),
		SoftSerial:       hardware.GetSoftSerial(t.Log),
		AgentName:        t.AgentName,
	})
	t.Log.Functionf("TestConnectivity: Use V2 API %v\n", zedcloud.UseV2API())
	testURL := zedcloud.URLPathString(serverNameAndPort, zedcloudCtx.V2API, nilUUID, "ping")

	tlsConfig, err := zedcloud.GetTlsConfig(zedcloudCtx.DeviceNetworkStatus, serverName,
		nil, &zedcloudCtx)
	if err != nil {
		t.Log.Functionf("TestConnectivity: " +
			"Device certificate not found, looking for Onboarding certificate")
		onboardingCert, err := tls.LoadX509KeyPair(types.OnboardCertName,
			types.OnboardKeyName)
		if err != nil {
			err = fmt.Errorf("onboarding certificate cannot be loaded: %v", err)
			t.Log.Functionf("TestConnectivity: %v\n", err)
			return intfStatusMap, err
		}
		clientCert := &onboardingCert
		tlsConfig, err = zedcloud.GetTlsConfig(zedcloudCtx.DeviceNetworkStatus,
			serverName, clientCert, &zedcloudCtx)
		if err != nil {
			err = fmt.Errorf("failed to load TLS config for talking to Zedcloud: %v", err)
			t.Log.Functionf("TestConnectivity: %v", err)
			return intfStatusMap, err
		}
	}

	if t.prevTLSConfig != nil {
		tlsConfig.ClientSessionCache = t.prevTLSConfig.ClientSessionCache
	}
	zedcloudCtx.TlsConfig = tlsConfig
	for ix := range dns.Ports {
		err = devicenetwork.CheckAndGetNetworkProxy(t.Log, &dns.Ports[ix], t.Metrics)
		if err != nil {
			ifName := dns.Ports[ix].IfName
			err = fmt.Errorf("failed to get network proxy for interface %s: %v",
				ifName, err)
			t.Log.Errorf("TestConnectivity: %v", err)
			intfStatusMap.RecordFailure(ifName, err.Error())
			return intfStatusMap, err
		}
	}
	cloudReachable, rtf, tempIntfStatusMap, err := zedcloud.VerifyAllIntf(
		&zedcloudCtx, testURL, requiredSuccessCount, t.iteration)
	intfStatusMap.SetOrUpdateFromMap(tempIntfStatusMap)
	t.Log.Tracef("TestConnectivity: intfStatusMap = %+v", intfStatusMap)
	if err != nil {
		if rtf {
			err = &RemoteTemporaryFailure{
				Endpoint:   serverNameAndPort,
				WrappedErr: err,
			}
		}
		t.Log.Errorf("TestConnectivity: %v", err)
		return intfStatusMap, err
	}

	t.prevTLSConfig = zedcloudCtx.TlsConfig
	if cloudReachable {
		t.Log.Functionf("TestConnectivity: uplink test SUCCEEDED for URL: %s", testURL)
		return intfStatusMap, nil
	}
	err = fmt.Errorf("uplink test FAILED for URL: %s", testURL)
	t.Log.Errorf("TestConnectivity: %v, intfStatusMap: %+v", err, intfStatusMap)
	return intfStatusMap, err
}
