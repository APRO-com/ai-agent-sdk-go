// Package signers
package signers

import (
	"ai-agent-go-sdk/attps/core/auth"
	"context"
)

// NULLSigner NULL digital signature generator
type NULLSigner struct{}

func (s *NULLSigner) Sign(_ context.Context, message string, timestamp int64) (*auth.SignatureResult, error) {
	return &auth.SignatureResult{AgentId: "", Signature: ""}, nil
}

func (s *NULLSigner) Algorithm() string {
	return "Null"
}
