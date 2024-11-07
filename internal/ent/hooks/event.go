package hooks

import (
	"context"
	"encoding/json"
	"fmt"

	"entgo.io/ent"

	"github.com/rs/zerolog/log"
	"github.com/stripe/stripe-go/v81"
	scust "github.com/stripe/stripe-go/v81/customer"

	entgen "github.com/theopenlane/core/internal/ent/generated"
	"github.com/theopenlane/core/pkg/events/soiree"
)

type EntEvent struct {
	*soiree.BaseEvent
}

func NewEntEvent(topic string, payload interface{}) *EntEvent {
	return &EntEvent{
		BaseEvent: soiree.NewBaseEvent(topic, payload),
	}
}

type Output struct {
	ID string `json:"id"`
}

func EmitEventHook(pool *soiree.EventPool) ent.Hook {
	return func(next ent.Mutator) ent.Mutator {
		return ent.MutateFunc(func(ctx context.Context, mutation ent.Mutation) (ent.Value, error) {
			retVal, err := next.Mutate(ctx, mutation)
			if err != nil {
				return nil, err
			}

			op := mutation.Op()
			typ := mutation.Type()

			fields := mutation.Fields()

			out, err := json.Marshal(retVal)
			if err != nil {
				log.Err(err).Msg("Failed to marshal return value")
			}

			other := Output{}
			json.Unmarshal(out, &other)

			event := NewEntEvent(fmt.Sprintf("%s.%s", typ, op), mutation)

			event.Properties().Set("ID", other.ID)

			for _, field := range fields {
				value, exists := mutation.Field(field)
				if exists {
					event.Properties().Set(field, value)
				}

				log.Info().Msgf("Field: %s, Value: %v", field, value)
			}

			event.SetContext(ctx)

			pool.Emit(event.Topic(), event)
			log.Info().Msgf("Emitted event: %s", event.Topic())
			log.Info().Msgf("Event properties: %v", event.Properties())

			return retVal, err
		})
	}
}

func RegisterGlobalHooks(client *entgen.Client, pool *soiree.EventPool) {
	client.Use(EmitEventHook(pool))
}

var EntEventPool *soiree.EventPool

func InitEventPool() *soiree.EventPool {
	customErrorHandler := func(event soiree.Event, err error) error {
		log.Printf("Error encountered during event '%s': %v, with payload: %v", event.Topic(), err, event.Payload())

		return nil // Returning nil to indicate that the error has been swallowed
	}

	entpool := soiree.NewNamedPondPool(100, "ent_event_pool") // nolint:mnd

	EntEventPool = soiree.NewEventPool(soiree.WithPool(entpool), soiree.WithErrorHandler(customErrorHandler))

	log.Info().Msg("Initialized EntEventPool")

	return EntEventPool
}

const (
	OrganizationCreated          = "Organization.OpCreate"
	OrganizationSettingCreated   = "OrganizationSetting.OpCreate"
	OrganizationSettingUpdated   = "OrganizationSetting.OpUpdate"
	OrganizationSettingUpdateOne = "OrganizationSetting.OpUpdateOne"
)

func RegisterListeners(pool *soiree.EventPool) {
	id, err := pool.On(OrganizationCreated, handleCustomerCreate)
	if err != nil {
		log.Err(err).Msgf("Failed to register listener for event: %s", id)
	}

	log.Info().Msgf("Registered listener for event type: %s", OrganizationCreated)

	id, err = pool.On(OrganizationSettingCreated, handleCustomerCreate)
	if err != nil {
		log.Err(err).Msgf("Failed to register listener for event: %s", id)
	}

	log.Info().Msgf("Registered listener for event type: %s", OrganizationSettingCreated)

	id, err = pool.On(OrganizationSettingUpdated, handleCustomerCreate)
	if err != nil {
		log.Err(err).Msgf("Failed to register listener for event: %s", id)
	}

	log.Info().Msgf("Registered listener for event type: %s", OrganizationSettingUpdated)

	id, err = pool.On(OrganizationSettingUpdateOne, handleCustomerCreate)
	if err != nil {
		log.Err(err).Msgf("Failed to register listener for event: %s", id)
	}

	log.Info().Msgf("Registered listener for event type: %s", OrganizationSettingUpdateOne)
}

func handleCustomerCreate(event soiree.Event) error {
	log.Info().Msg("Handling organization create event")

	stripe.Key = ""

	props := event.Properties()
	// this map is empty?
	log.Info().Msgf("Event properties from listener: %v", props)

	orgsettingID, exists := props["ID"]
	if !exists {
		log.Info().Msg("OrganizationSetting ID not found in event properties")
		return nil
	}

	log.Info().Msgf("OrganizationSetting ID: %s", orgsettingID)

	billingEmail, exists := props["billing_email"]

	if exists && billingEmail != "" {
		log.Info().Msgf("Billing email: %s", billingEmail)
		email := billingEmail.(string)
		customerParams := &stripe.CustomerListParams{
			Email: &email,
		}

		i := scust.List(customerParams)

		if !i.Next() {
			log.Info().Msgf("Attmpting to create Stripe customer with email %s", email)
			customerParams := &stripe.CustomerParams{
				Email: &email,
			}

			customer, err := scust.New(customerParams)
			if err != nil {
				log.Err(err).Msg("Failed to create Stripe customer")
				return err
			}

			log.Info().Msgf("Created Stripe customer with ID: %s", customer.ID)

			if err := updateOrganizationSettingWithCustomerID(orgsettingID.(string), customer.ID, entgen.FromContext(event.Context())); err != nil {
				log.Err(err).Msg("Failed to update OrganizationSetting with Stripe customer ID")
				return err
			}

			log.Info().Msgf("Updated OrganizationSetting with Stripe customer ID: %s", customer.ID)
		}

		log.Info().Msgf("Stripe customer with email %s already exists", email)
	}

	return nil
}

func updateOrganizationSettingWithCustomerID(orgID, customerID string, dbclient *entgen.Client) error {
	_, err := dbclient.OrganizationSetting.UpdateOneID(orgID).SetStripeID(customerID).Save(context.Background())
	if err != nil {
		log.Err(err).Msgf("Failed to update OrganizationSetting ID %s with Stripe customer ID %s", orgID, customerID)
		return err
	}

	log.Info().Msgf("Updated OrganizationSetting ID %s with Stripe customer ID %s", orgID, customerID)

	return nil
}

func CreateStripeCustomerIfNotExists(name, email string) error {
	customerParams := &stripe.CustomerParams{
		Email: &email,
		Name:  &name,
	}

	_, err := scust.New(customerParams)
	if err != nil {
		log.Printf("Failed to create Stripe customer: %v", err)
	}

	return nil
}

func MapMutationFieldsToProperties(mutation ent.Mutation) soiree.Properties {
	properties := soiree.Properties{}

	for _, field := range mutation.Fields() {
		value, exists := mutation.Field(field)
		if exists {
			properties[field] = value
		}
	}

	return properties
}
