// Package credentials Go SDK Request header Authorization information generator
package credentials

import (
	"context"
	"fmt"
	"github.com/APRO-com/ai-agent-sdk-go/attps/core/auth"
	"time"
)

// Credentials Request Header Authorization Information Generator
type Credentials struct {
	Signer auth.Signer // 数字签名生成器
}

func (c *Credentials) GenerateAuthorization(
	ctx context.Context, signBody string,
) (string, error) {
	if c.Signer == nil {
		return "", fmt.Errorf("you must init Credentials with signer")
	}

	timestamp := time.Now().Unix()
	signatureResult, err := c.Signer.Sign(ctx, signBody, timestamp)
	if err != nil {
		return "", err
	}

	authorization := signatureResult.Signature

	return authorization, nil
}
