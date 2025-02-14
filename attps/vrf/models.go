package vrf

import (
	"ai-agent-go-sdk/attps/service"
)

// VRFRequest Random Request
type VRFRequest struct {
	Version          *int64  `json:"version"`
	TargetAgentID    *string `json:"target_agent_id"`
	ClientSeed       *string `json:"client_seed"`
	KeyHash          *string `json:"key_hash"`
	RequestTimestamp *int64  `json:"request_timestamp"`
	RequestID        *string `json:"request_id"`
	CallbackURI      *string `json:"callback_uri"`
}

// VRFResponse Random Response
type VRFResponse struct {
	service.BaseResponse
	Result *string `json:"result"`
}
