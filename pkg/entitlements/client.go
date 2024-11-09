package entitlements

import (
	"github.com/stripe/stripe-go/v81"
	"github.com/stripe/stripe-go/v81/client"
)

type StripeClient struct {
	Client   *client.API
	apikey   string
	Products []stripe.Product
	Prices   []stripe.Price
}

func NewStripeClient(opts ...StripeOptions) *StripeClient {
	sc := &StripeClient{}
	for _, opt := range opts {
		opt(sc)
	}

	sc.Client = client.New(sc.apikey, nil)

	return sc
}

type StripeOptions func(*StripeClient)

func WithAPIKey(apiKey string) StripeOptions {
	return func(sc *StripeClient) {
		sc.apikey = apiKey
	}
}
