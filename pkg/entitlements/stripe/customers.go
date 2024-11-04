package ripe

import (
	"os"

	"github.com/stripe/stripe-go/v81"
	"github.com/stripe/stripe-go/v81/client"
	scust "github.com/stripe/stripe-go/v81/customer"
	"github.com/stripe/stripe-go/v81/product"
)

// Customer is a wrapper for the stripe customer object
type Customer struct {
	*stripe.Customer
}

// NewCustomer creates a new customer
func NewCustomer() *Customer {
	return &Customer{&stripe.Customer{}}
}

type Stripe interface {
	CreateCustomer() (*Customer, error)
}

type stripeClient struct {
}

// CreateCustomer implements Stripe.
func (s *stripeClient) CreateCustomer() (*Customer, error) {
	panic("unimplemented")
}

func NewStripe() Stripe {
	return &stripeClient{}
}

type StripeWrapper struct {
	CustomerParams stripe.CustomerParams
	Customer       stripe.Customer
	Product        stripe.Product
	ProductParams  stripe.ProductParams
	Price          stripe.Price
	PriceParams    stripe.PriceParams
}

func NewStripeCustomer(custparams stripe.CustomerParams) error {
	_, err := scust.New(&custparams)

	return err
}

func NewStripeProduct(productparams stripe.ProductParams) error {
	product.New(&productparams)

	return nil
}

type StripeClient struct {
	Client *client.API
}

func NewStripeClient() *StripeClient {
	apiKey := os.Getenv("STRIPE_API_KEY")
	if apiKey == "" {
		panic("STRIPE_API_KEY environment variable is not set")
	}

	sc := &client.API{}
	sc.Init(apiKey, nil)

	return &StripeClient{Client: sc}
}
