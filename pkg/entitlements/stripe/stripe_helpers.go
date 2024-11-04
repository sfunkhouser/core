package ripe

import (
	"github.com/stripe/stripe-go/v81"
	scust "github.com/stripe/stripe-go/v81/customer"
	splan "github.com/stripe/stripe-go/v81/plan"
	sprice "github.com/stripe/stripe-go/v81/price"
	sprod "github.com/stripe/stripe-go/v81/product"
)

// CreateCustomer creates a new Stripe customer
func (sc *StripeClient) CreateCustomer(params *stripe.CustomerParams) (*stripe.Customer, error) {
	return scust.New(params)
}

// UpdateCustomer updates an existing Stripe customer
func (sc *StripeClient) UpdateCustomer(id string, params *stripe.CustomerParams) (*stripe.Customer, error) {
	return scust.Update(id, params)
}

// CreateProduct creates a new Stripe product
func (sc *StripeClient) CreateProduct(params *stripe.ProductParams) (*stripe.Product, error) {
	return sprod.New(params)
}

// UpdateProduct updates an existing Stripe product
func (sc *StripeClient) UpdateProduct(id string, params *stripe.ProductParams) (*stripe.Product, error) {
	return sprod.Update(id, params)
}

// CreatePrice creates a new Stripe price
func (sc *StripeClient) CreatePrice(params *stripe.PriceParams) (*stripe.Price, error) {
	return sprice.New(params)
}

// UpdatePrice updates an existing Stripe price
func (sc *StripeClient) UpdatePrice(id string, params *stripe.PriceParams) (*stripe.Price, error) {
	return sprice.Update(id, params)
}

// CreatePlan creates a new Stripe plan
func (sc *StripeClient) CreatePlan(params *stripe.PlanParams) (*stripe.Plan, error) {
	return splan.New(params)
}

// UpdatePlan updates an existing Stripe plan
func (sc *StripeClient) UpdatePlan(id string, params *stripe.PlanParams) (*stripe.Plan, error) {
	return splan.Update(id, params)
}
