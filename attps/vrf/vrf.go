package vrf

import (
	"ai-agent-go-sdk/attps/core"
	"ai-agent-go-sdk/attps/core/consts"
	"ai-agent-go-sdk/attps/service"
	"ai-agent-go-sdk/util"
	"context"
	"encoding/hex"
	"errors"
	"fmt"
	"github.com/ethereum/go-ethereum/crypto"
	nethttp "net/http"
	neturl "net/url"
)

type VRF service.Service

// Request get a vrf random number and verify the returned proof data
func (v *VRF) Request(ctx context.Context, req VRFRequest) (random string, err error) {
	var (
		localVarHTTPMethod   = nethttp.MethodPost
		localVarPostBody     interface{}
		localVarQueryParams  neturl.Values
		localVarHeaderParams = nethttp.Header{}
	)

	localVarPath := consts.APIBaseServer + "/api/vrf/request"

	// Setup Body Params and check
	localVarPostBody = &req

	if err = v.checkRequestParams(req); err != nil {
		return "", err
	}

	// Determine the Content-Type Header
	localVarHTTPContentTypes := []string{"application/json"}
	// Setup Content-Type
	localVarHTTPContentType := core.SelectHeaderContentType(localVarHTTPContentTypes)

	// Perform Http Request
	result, err := v.Client.Request(ctx, localVarHTTPMethod, localVarPath, localVarHeaderParams, localVarQueryParams, localVarPostBody, localVarHTTPContentType)
	if err != nil {
		return "", err
	}

	resp := new(VRFResponse)
	err = core.UnMarshalResponse(result.Response, resp)
	if err != nil {
		return "", err
	}

	if resp.BaseResponse.Code != 0 {
		return "", errors.New(*resp.Message)
	}

	return *resp.Result, nil
}

func (v *VRF) CalRequestID(version int64, targetAgentId string, customerFeed string, requestTimestamp int64, callbackUri string) (string, error) {
	// Concatenate byte arrays
	t := append(util.LongToBytes(version), []byte(targetAgentId)...)
	t = append(t, *core.HexBytes(customerFeed)...)
	t = append(t, util.LongToBytes(requestTimestamp)...)
	t = append(t, []byte(callbackUri)...)

	// Compute Keccak256 hash
	hash := crypto.Keccak256(t)
	requestId := hex.EncodeToString(hash[:])

	return requestId, nil
}

func (v *VRF) checkRequestParams(req VRFRequest) error {
	if *req.Version != consts.VRFVersion {
		return errors.New(fmt.Sprintf("vrf version mismatch, must %d", consts.VRFVersion))
	}
	if !util.IsUUIDv4(*req.TargetAgentID) {
		return errors.New(fmt.Sprintf("invalid target agent id: %s", *req.TargetAgentID))
	}
	if !util.IsHexString(*req.ClientFeed) {
		return errors.New(fmt.Sprintf("invalid client feed: %s", *req.ClientFeed))
	}
	if !util.IsHexString(*req.KeyHash) {
		return errors.New(fmt.Sprintf("invalid key hash: %s", *req.KeyHash))
	}
	//if !util.IsValid13DigitTimestamp(big.NewInt(*req.RequestTimestamp)) {
	//	return errors.New(fmt.Sprintf("invalid request timestamp: %d", *req.RequestTimestamp))
	//}

	requestID, err := v.CalRequestID(*req.Version, *req.TargetAgentID, *req.ClientFeed, *req.RequestTimestamp, *req.CallbackURI)
	if err != nil {
		return err
	}
	if *req.RequestID != requestID {
		return errors.New(fmt.Sprintf("invalid request ID: %s", *req.RequestID))
	}

	return nil
}
