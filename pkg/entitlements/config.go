package entitlements

import (
	ripe "github.com/theopenlane/core/pkg/entitlements/stripe"
)

type Config struct {
	Enabled bool `json:"enabled" koanf:"enabled" default:"false"`
	// StripeConfig is the configuration for the stripe service
	StripeConfig ripe.Config
}
