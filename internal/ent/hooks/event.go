package hooks

import (
	"context"
	"fmt"

	"entgo.io/ent"

	"github.com/rs/zerolog/log"

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

			event := NewEntEvent(fmt.Sprintf("%s.%s", typ, op), mutation)
			for _, field := range fields {
				value, exists := mutation.Field(field)
				if exists {
					event.Properties()[field] = value
				}

				log.Info().Msgf("Field: %s, Value: %v", field, value)
			}

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
	OrganizationCreated = "Organization.OpCreate"
)

func RegisterListeners(pool *soiree.EventPool) {
	id, err := pool.On(OrganizationCreated, handleCustomerCreate)
	if err != nil {
		log.Err(err).Msgf("Failed to register listener for event: %s", id)
	}

	log.Info().Msgf("Registered listener for event: %s", id)
	log.Info().Msgf("Registered listener for event type: %s", OrganizationCreated)
}

func handleCustomerCreate(event soiree.Event) error {
	log.Info().Msg("Handling organization create event")
	props := event.Properties()
	for k, v := range props {
		log.Info().Msgf("Key: %s, Value: %v", k, v)
	}

	// log.Info().Msgf("Organization Name: %s", name)
	// log.Info().Msgf("Event Topic: %s", event.Topic())
	//	mutation.Client().OrganizationSetting.Query().Where(organizationsetting.OrganizationID(id)).Only(context.Background())

	//	customerParams := &stripe.CustomerParams{
	//        Email: mutation.setting.
	//		        Name:  &name,
	//	}
	//	_, err := scust.New(customerParams)
	//	if err != nil {
	//		log.Printf("Failed to create Stripe customer: %v", err)
	//	}

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
