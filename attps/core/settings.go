package core

import (
	"fmt"
	"net/http"

	"ai-agent-go-sdk/attps/core/auth"
)

// DialSettings core.Client Configuration information required
type DialSettings struct {
	HTTPClient *http.Client
	Signer     auth.Signer
}

// Validate checks whether the request configuration is valid
func (ds *DialSettings) Validate() error {
	if ds.Signer == nil {
		return fmt.Errorf("signer is required for Client")
	}
	return nil
}
