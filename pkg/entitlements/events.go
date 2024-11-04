package entitlements

import (
	"context"
	"log"

	"github.com/stripe/stripe-go/v81"
	scust "github.com/stripe/stripe-go/v81/customer"

	ent "github.com/theopenlane/core/internal/ent/generated"
	"github.com/theopenlane/core/internal/ent/generated/organizationsetting"
	"github.com/theopenlane/core/pkg/events/soiree"
)

func init() {
	// hooks.EntEventPool.On("Customer_create", handleCustomerCreate)
}

func handleCustomerCreate(event soiree.Event) error {
	mutation, ok := event.Payload().(*ent.OrganizationMutation)
	if !ok {
		log.Printf("Failed to cast payload to CustomerMutation")
		return nil
	}

	id, _ := mutation.ID()
	mutation.Client().OrganizationSetting.Query().Where(organizationsetting.OrganizationID(id)).Only(context.Background())

	customerParams := &stripe.CustomerParams{
		//        Email: mutation.setting.
		//        Name:  mutation.Name,
	}
	_, err := scust.New(customerParams)
	if err != nil {
		log.Printf("Failed to create Stripe customer: %v", err)
	}
	return err
}
