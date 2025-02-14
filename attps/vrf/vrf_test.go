package vrf

import (
	"ai-agent-go-sdk/attps/core"
	"ai-agent-go-sdk/attps/core/option"
	"ai-agent-go-sdk/util"
	"context"
	"log"
	"testing"
	"time"
)

func TestVRF_Request(t *testing.T) {
	ctx := context.Background()

	opts := []core.ClientOption{
		option.WithNullAuthCipher(),
	}
	client, err := core.NewClient(ctx, opts...)
	if err != nil {
		log.Printf("new client err:%s", err)
	}

	svc := VRF{Client: client}
	version := int64(1)
	targetAgentID := "f2464336-fbcf-4603-bda5-ce65c0318fb6"
	customerFeed := util.SecureRandomString(4)
	keyHash := "d004ced906bcdb3ebcbf706dd9a284367b6d3e25a91c91b5a430af2593886eb9"
	callbackUri := "http://127.0.0.1:8888/api/vrf/proof"

	requestTimestamp := time.Now().Unix()
	requestID, err := new(VRF).CalRequestID(version, targetAgentID, customerFeed, requestTimestamp, callbackUri)
	if err != nil {
		t.Errorf("call RequestID err:%s", err)
		return
	}
	random, err := svc.Request(ctx,
		VRFRequest{
			Version:          core.Int64(version),
			TargetAgentID:    core.String(targetAgentID),
			ClientSeed:       core.String(customerFeed),
			KeyHash:          core.String(keyHash),
			RequestTimestamp: core.Int64(requestTimestamp),
			RequestID:        core.String(requestID),
			CallbackURI:      core.String(callbackUri),
		},
	)

	if err != nil {
		t.Errorf("call Request err:%s", err)
		return
	} else {
		t.Log(random)
	}
}

func TestVRF_CalRequestID(t *testing.T) {
	benchmarkRequestID := "6f71619f1e6ea42616c9bbdc8fe001511e0c37b72373dc259857b29c1e61597c"
	version := int64(1)
	targetAgentID := "f2464336-fbcf-4603-bda5-ce65c0318fb6"
	customerFeed := "1234"
	callbackUri := "http://127.0.0.1:8888/api/vrf/proof"
	requestTimestamp := int64(1739265192)

	requestID, err := new(VRF).CalRequestID(version, targetAgentID, customerFeed, requestTimestamp, callbackUri)
	if err != nil {
		t.Errorf("call RequestID err:%s", err)
		return
	}
	if requestID != benchmarkRequestID {
		t.Errorf("requestID:%s != %s", requestID, benchmarkRequestID)
		return
	}
}
