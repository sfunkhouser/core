package payments

import "github.com/stripe/stripe-go/v81/client"

type PaymentProvider interface {
	Charge() error
}

type Payment struct {
	Amount int
}

func NewPayment(amount int) *Payment {
	return &Payment{Amount: amount}
}

func (p *Payment) Charge(provider PaymentProvider) error {
	return provider.Charge()
}

type StripeProvider struct {
	Client *client.API
}

func (s *StripeProvider) Charge() error {
	return nil
}

func NewStripeProvider(client *client.API) *StripeProvider {
	return &StripeProvider{Client: client}
}

func NewStripeClient(opts ...StripeOption) *client.API {
	options := NewStripeOptions(opts...)
	if options.SecretKey == "" {
		panic("missing secret key")
	}

	return client.New(options.SecretKey, nil)
}

type StripeOptions struct {
	SecretKey            string
	PublishableKey       string
	WebhookSigningSecret string
}

type StripeOption func(*StripeOptions)

func WithSecretKey(secretKey string) StripeOption {
	return func(s *StripeOptions) {
		s.SecretKey = secretKey
	}
}

func WithPublishableKey(publishableKey string) StripeOption {
	return func(s *StripeOptions) {
		s.PublishableKey = publishableKey
	}
}

func WithWebhookSigningSecret(webhookSigningSecret string) StripeOption {
	return func(s *StripeOptions) {
		s.WebhookSigningSecret = webhookSigningSecret
	}
}

func NewStripeOptions(opts ...StripeOption) *StripeOptions {
	options := &StripeOptions{}
	for _, opt := range opts {
		opt(options)
	}

	return options
}
