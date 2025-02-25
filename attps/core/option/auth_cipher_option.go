package option

import (
	"github.com/APRO-com/ATTPs-sdk-go/attps/core"
	"github.com/APRO-com/ATTPs-sdk-go/attps/core/auth/signers"
)

type withAuthCipherOption struct{ settings core.DialSettings }

// Apply Set Signer, Validator, and Cipher for core.DialSettings
func (w withAuthCipherOption) Apply(o *core.DialSettings) error {
	o.Signer = w.settings.Signer
	return nil
}

// WithNullAuthCipher
func WithNullAuthCipher() core.ClientOption {
	return withAuthCipherOption{
		settings: core.DialSettings{
			Signer: &signers.NULLSigner{},
		},
	}
}

// WithIdentifierAuthCipher
func WithIdentifierAuthCipher(identifier string) core.ClientOption {
	return withAuthCipherOption{
		settings: core.DialSettings{
			Signer: &signers.IdentifierSigner{
				Identifier: identifier,
			},
		},
	}
}
