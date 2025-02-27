// Code generated by ent, DO NOT EDIT.

package generated

import (
	"context"
	"errors"
	"fmt"
	"time"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/theopenlane/core/internal/ent/generated/entitlement"
	"github.com/theopenlane/core/internal/ent/generated/entitlementplan"
	"github.com/theopenlane/core/internal/ent/generated/entitlementplanfeature"
	"github.com/theopenlane/core/internal/ent/generated/event"
	"github.com/theopenlane/core/internal/ent/generated/feature"
	"github.com/theopenlane/core/internal/ent/generated/organization"
)

// EntitlementPlanCreate is the builder for creating a EntitlementPlan entity.
type EntitlementPlanCreate struct {
	config
	mutation *EntitlementPlanMutation
	hooks    []Hook
}

// SetCreatedAt sets the "created_at" field.
func (epc *EntitlementPlanCreate) SetCreatedAt(t time.Time) *EntitlementPlanCreate {
	epc.mutation.SetCreatedAt(t)
	return epc
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (epc *EntitlementPlanCreate) SetNillableCreatedAt(t *time.Time) *EntitlementPlanCreate {
	if t != nil {
		epc.SetCreatedAt(*t)
	}
	return epc
}

// SetUpdatedAt sets the "updated_at" field.
func (epc *EntitlementPlanCreate) SetUpdatedAt(t time.Time) *EntitlementPlanCreate {
	epc.mutation.SetUpdatedAt(t)
	return epc
}

// SetNillableUpdatedAt sets the "updated_at" field if the given value is not nil.
func (epc *EntitlementPlanCreate) SetNillableUpdatedAt(t *time.Time) *EntitlementPlanCreate {
	if t != nil {
		epc.SetUpdatedAt(*t)
	}
	return epc
}

// SetCreatedBy sets the "created_by" field.
func (epc *EntitlementPlanCreate) SetCreatedBy(s string) *EntitlementPlanCreate {
	epc.mutation.SetCreatedBy(s)
	return epc
}

// SetNillableCreatedBy sets the "created_by" field if the given value is not nil.
func (epc *EntitlementPlanCreate) SetNillableCreatedBy(s *string) *EntitlementPlanCreate {
	if s != nil {
		epc.SetCreatedBy(*s)
	}
	return epc
}

// SetUpdatedBy sets the "updated_by" field.
func (epc *EntitlementPlanCreate) SetUpdatedBy(s string) *EntitlementPlanCreate {
	epc.mutation.SetUpdatedBy(s)
	return epc
}

// SetNillableUpdatedBy sets the "updated_by" field if the given value is not nil.
func (epc *EntitlementPlanCreate) SetNillableUpdatedBy(s *string) *EntitlementPlanCreate {
	if s != nil {
		epc.SetUpdatedBy(*s)
	}
	return epc
}

// SetMappingID sets the "mapping_id" field.
func (epc *EntitlementPlanCreate) SetMappingID(s string) *EntitlementPlanCreate {
	epc.mutation.SetMappingID(s)
	return epc
}

// SetNillableMappingID sets the "mapping_id" field if the given value is not nil.
func (epc *EntitlementPlanCreate) SetNillableMappingID(s *string) *EntitlementPlanCreate {
	if s != nil {
		epc.SetMappingID(*s)
	}
	return epc
}

// SetDeletedAt sets the "deleted_at" field.
func (epc *EntitlementPlanCreate) SetDeletedAt(t time.Time) *EntitlementPlanCreate {
	epc.mutation.SetDeletedAt(t)
	return epc
}

// SetNillableDeletedAt sets the "deleted_at" field if the given value is not nil.
func (epc *EntitlementPlanCreate) SetNillableDeletedAt(t *time.Time) *EntitlementPlanCreate {
	if t != nil {
		epc.SetDeletedAt(*t)
	}
	return epc
}

// SetDeletedBy sets the "deleted_by" field.
func (epc *EntitlementPlanCreate) SetDeletedBy(s string) *EntitlementPlanCreate {
	epc.mutation.SetDeletedBy(s)
	return epc
}

// SetNillableDeletedBy sets the "deleted_by" field if the given value is not nil.
func (epc *EntitlementPlanCreate) SetNillableDeletedBy(s *string) *EntitlementPlanCreate {
	if s != nil {
		epc.SetDeletedBy(*s)
	}
	return epc
}

// SetTags sets the "tags" field.
func (epc *EntitlementPlanCreate) SetTags(s []string) *EntitlementPlanCreate {
	epc.mutation.SetTags(s)
	return epc
}

// SetOwnerID sets the "owner_id" field.
func (epc *EntitlementPlanCreate) SetOwnerID(s string) *EntitlementPlanCreate {
	epc.mutation.SetOwnerID(s)
	return epc
}

// SetNillableOwnerID sets the "owner_id" field if the given value is not nil.
func (epc *EntitlementPlanCreate) SetNillableOwnerID(s *string) *EntitlementPlanCreate {
	if s != nil {
		epc.SetOwnerID(*s)
	}
	return epc
}

// SetDisplayName sets the "display_name" field.
func (epc *EntitlementPlanCreate) SetDisplayName(s string) *EntitlementPlanCreate {
	epc.mutation.SetDisplayName(s)
	return epc
}

// SetNillableDisplayName sets the "display_name" field if the given value is not nil.
func (epc *EntitlementPlanCreate) SetNillableDisplayName(s *string) *EntitlementPlanCreate {
	if s != nil {
		epc.SetDisplayName(*s)
	}
	return epc
}

// SetName sets the "name" field.
func (epc *EntitlementPlanCreate) SetName(s string) *EntitlementPlanCreate {
	epc.mutation.SetName(s)
	return epc
}

// SetDescription sets the "description" field.
func (epc *EntitlementPlanCreate) SetDescription(s string) *EntitlementPlanCreate {
	epc.mutation.SetDescription(s)
	return epc
}

// SetNillableDescription sets the "description" field if the given value is not nil.
func (epc *EntitlementPlanCreate) SetNillableDescription(s *string) *EntitlementPlanCreate {
	if s != nil {
		epc.SetDescription(*s)
	}
	return epc
}

// SetVersion sets the "version" field.
func (epc *EntitlementPlanCreate) SetVersion(s string) *EntitlementPlanCreate {
	epc.mutation.SetVersion(s)
	return epc
}

// SetMetadata sets the "metadata" field.
func (epc *EntitlementPlanCreate) SetMetadata(m map[string]interface{}) *EntitlementPlanCreate {
	epc.mutation.SetMetadata(m)
	return epc
}

// SetStripeProductID sets the "stripe_product_id" field.
func (epc *EntitlementPlanCreate) SetStripeProductID(s string) *EntitlementPlanCreate {
	epc.mutation.SetStripeProductID(s)
	return epc
}

// SetNillableStripeProductID sets the "stripe_product_id" field if the given value is not nil.
func (epc *EntitlementPlanCreate) SetNillableStripeProductID(s *string) *EntitlementPlanCreate {
	if s != nil {
		epc.SetStripeProductID(*s)
	}
	return epc
}

// SetStripePriceID sets the "stripe_price_id" field.
func (epc *EntitlementPlanCreate) SetStripePriceID(s string) *EntitlementPlanCreate {
	epc.mutation.SetStripePriceID(s)
	return epc
}

// SetNillableStripePriceID sets the "stripe_price_id" field if the given value is not nil.
func (epc *EntitlementPlanCreate) SetNillableStripePriceID(s *string) *EntitlementPlanCreate {
	if s != nil {
		epc.SetStripePriceID(*s)
	}
	return epc
}

// SetID sets the "id" field.
func (epc *EntitlementPlanCreate) SetID(s string) *EntitlementPlanCreate {
	epc.mutation.SetID(s)
	return epc
}

// SetNillableID sets the "id" field if the given value is not nil.
func (epc *EntitlementPlanCreate) SetNillableID(s *string) *EntitlementPlanCreate {
	if s != nil {
		epc.SetID(*s)
	}
	return epc
}

// SetOwner sets the "owner" edge to the Organization entity.
func (epc *EntitlementPlanCreate) SetOwner(o *Organization) *EntitlementPlanCreate {
	return epc.SetOwnerID(o.ID)
}

// AddEntitlementIDs adds the "entitlements" edge to the Entitlement entity by IDs.
func (epc *EntitlementPlanCreate) AddEntitlementIDs(ids ...string) *EntitlementPlanCreate {
	epc.mutation.AddEntitlementIDs(ids...)
	return epc
}

// AddEntitlements adds the "entitlements" edges to the Entitlement entity.
func (epc *EntitlementPlanCreate) AddEntitlements(e ...*Entitlement) *EntitlementPlanCreate {
	ids := make([]string, len(e))
	for i := range e {
		ids[i] = e[i].ID
	}
	return epc.AddEntitlementIDs(ids...)
}

// AddBaseFeatureIDs adds the "base_features" edge to the Feature entity by IDs.
func (epc *EntitlementPlanCreate) AddBaseFeatureIDs(ids ...string) *EntitlementPlanCreate {
	epc.mutation.AddBaseFeatureIDs(ids...)
	return epc
}

// AddBaseFeatures adds the "base_features" edges to the Feature entity.
func (epc *EntitlementPlanCreate) AddBaseFeatures(f ...*Feature) *EntitlementPlanCreate {
	ids := make([]string, len(f))
	for i := range f {
		ids[i] = f[i].ID
	}
	return epc.AddBaseFeatureIDs(ids...)
}

// AddEventIDs adds the "events" edge to the Event entity by IDs.
func (epc *EntitlementPlanCreate) AddEventIDs(ids ...string) *EntitlementPlanCreate {
	epc.mutation.AddEventIDs(ids...)
	return epc
}

// AddEvents adds the "events" edges to the Event entity.
func (epc *EntitlementPlanCreate) AddEvents(e ...*Event) *EntitlementPlanCreate {
	ids := make([]string, len(e))
	for i := range e {
		ids[i] = e[i].ID
	}
	return epc.AddEventIDs(ids...)
}

// AddFeatureIDs adds the "features" edge to the EntitlementPlanFeature entity by IDs.
func (epc *EntitlementPlanCreate) AddFeatureIDs(ids ...string) *EntitlementPlanCreate {
	epc.mutation.AddFeatureIDs(ids...)
	return epc
}

// AddFeatures adds the "features" edges to the EntitlementPlanFeature entity.
func (epc *EntitlementPlanCreate) AddFeatures(e ...*EntitlementPlanFeature) *EntitlementPlanCreate {
	ids := make([]string, len(e))
	for i := range e {
		ids[i] = e[i].ID
	}
	return epc.AddFeatureIDs(ids...)
}

// Mutation returns the EntitlementPlanMutation object of the builder.
func (epc *EntitlementPlanCreate) Mutation() *EntitlementPlanMutation {
	return epc.mutation
}

// Save creates the EntitlementPlan in the database.
func (epc *EntitlementPlanCreate) Save(ctx context.Context) (*EntitlementPlan, error) {
	if err := epc.defaults(); err != nil {
		return nil, err
	}
	return withHooks(ctx, epc.sqlSave, epc.mutation, epc.hooks)
}

// SaveX calls Save and panics if Save returns an error.
func (epc *EntitlementPlanCreate) SaveX(ctx context.Context) *EntitlementPlan {
	v, err := epc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (epc *EntitlementPlanCreate) Exec(ctx context.Context) error {
	_, err := epc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (epc *EntitlementPlanCreate) ExecX(ctx context.Context) {
	if err := epc.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (epc *EntitlementPlanCreate) defaults() error {
	if _, ok := epc.mutation.CreatedAt(); !ok {
		if entitlementplan.DefaultCreatedAt == nil {
			return fmt.Errorf("generated: uninitialized entitlementplan.DefaultCreatedAt (forgotten import generated/runtime?)")
		}
		v := entitlementplan.DefaultCreatedAt()
		epc.mutation.SetCreatedAt(v)
	}
	if _, ok := epc.mutation.UpdatedAt(); !ok {
		if entitlementplan.DefaultUpdatedAt == nil {
			return fmt.Errorf("generated: uninitialized entitlementplan.DefaultUpdatedAt (forgotten import generated/runtime?)")
		}
		v := entitlementplan.DefaultUpdatedAt()
		epc.mutation.SetUpdatedAt(v)
	}
	if _, ok := epc.mutation.MappingID(); !ok {
		if entitlementplan.DefaultMappingID == nil {
			return fmt.Errorf("generated: uninitialized entitlementplan.DefaultMappingID (forgotten import generated/runtime?)")
		}
		v := entitlementplan.DefaultMappingID()
		epc.mutation.SetMappingID(v)
	}
	if _, ok := epc.mutation.Tags(); !ok {
		v := entitlementplan.DefaultTags
		epc.mutation.SetTags(v)
	}
	if _, ok := epc.mutation.ID(); !ok {
		if entitlementplan.DefaultID == nil {
			return fmt.Errorf("generated: uninitialized entitlementplan.DefaultID (forgotten import generated/runtime?)")
		}
		v := entitlementplan.DefaultID()
		epc.mutation.SetID(v)
	}
	return nil
}

// check runs all checks and user-defined validators on the builder.
func (epc *EntitlementPlanCreate) check() error {
	if _, ok := epc.mutation.MappingID(); !ok {
		return &ValidationError{Name: "mapping_id", err: errors.New(`generated: missing required field "EntitlementPlan.mapping_id"`)}
	}
	if v, ok := epc.mutation.OwnerID(); ok {
		if err := entitlementplan.OwnerIDValidator(v); err != nil {
			return &ValidationError{Name: "owner_id", err: fmt.Errorf(`generated: validator failed for field "EntitlementPlan.owner_id": %w`, err)}
		}
	}
	if _, ok := epc.mutation.Name(); !ok {
		return &ValidationError{Name: "name", err: errors.New(`generated: missing required field "EntitlementPlan.name"`)}
	}
	if v, ok := epc.mutation.Name(); ok {
		if err := entitlementplan.NameValidator(v); err != nil {
			return &ValidationError{Name: "name", err: fmt.Errorf(`generated: validator failed for field "EntitlementPlan.name": %w`, err)}
		}
	}
	if _, ok := epc.mutation.Version(); !ok {
		return &ValidationError{Name: "version", err: errors.New(`generated: missing required field "EntitlementPlan.version"`)}
	}
	if v, ok := epc.mutation.Version(); ok {
		if err := entitlementplan.VersionValidator(v); err != nil {
			return &ValidationError{Name: "version", err: fmt.Errorf(`generated: validator failed for field "EntitlementPlan.version": %w`, err)}
		}
	}
	return nil
}

func (epc *EntitlementPlanCreate) sqlSave(ctx context.Context) (*EntitlementPlan, error) {
	if err := epc.check(); err != nil {
		return nil, err
	}
	_node, _spec := epc.createSpec()
	if err := sqlgraph.CreateNode(ctx, epc.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	if _spec.ID.Value != nil {
		if id, ok := _spec.ID.Value.(string); ok {
			_node.ID = id
		} else {
			return nil, fmt.Errorf("unexpected EntitlementPlan.ID type: %T", _spec.ID.Value)
		}
	}
	epc.mutation.id = &_node.ID
	epc.mutation.done = true
	return _node, nil
}

func (epc *EntitlementPlanCreate) createSpec() (*EntitlementPlan, *sqlgraph.CreateSpec) {
	var (
		_node = &EntitlementPlan{config: epc.config}
		_spec = sqlgraph.NewCreateSpec(entitlementplan.Table, sqlgraph.NewFieldSpec(entitlementplan.FieldID, field.TypeString))
	)
	_spec.Schema = epc.schemaConfig.EntitlementPlan
	if id, ok := epc.mutation.ID(); ok {
		_node.ID = id
		_spec.ID.Value = id
	}
	if value, ok := epc.mutation.CreatedAt(); ok {
		_spec.SetField(entitlementplan.FieldCreatedAt, field.TypeTime, value)
		_node.CreatedAt = value
	}
	if value, ok := epc.mutation.UpdatedAt(); ok {
		_spec.SetField(entitlementplan.FieldUpdatedAt, field.TypeTime, value)
		_node.UpdatedAt = value
	}
	if value, ok := epc.mutation.CreatedBy(); ok {
		_spec.SetField(entitlementplan.FieldCreatedBy, field.TypeString, value)
		_node.CreatedBy = value
	}
	if value, ok := epc.mutation.UpdatedBy(); ok {
		_spec.SetField(entitlementplan.FieldUpdatedBy, field.TypeString, value)
		_node.UpdatedBy = value
	}
	if value, ok := epc.mutation.MappingID(); ok {
		_spec.SetField(entitlementplan.FieldMappingID, field.TypeString, value)
		_node.MappingID = value
	}
	if value, ok := epc.mutation.DeletedAt(); ok {
		_spec.SetField(entitlementplan.FieldDeletedAt, field.TypeTime, value)
		_node.DeletedAt = value
	}
	if value, ok := epc.mutation.DeletedBy(); ok {
		_spec.SetField(entitlementplan.FieldDeletedBy, field.TypeString, value)
		_node.DeletedBy = value
	}
	if value, ok := epc.mutation.Tags(); ok {
		_spec.SetField(entitlementplan.FieldTags, field.TypeJSON, value)
		_node.Tags = value
	}
	if value, ok := epc.mutation.DisplayName(); ok {
		_spec.SetField(entitlementplan.FieldDisplayName, field.TypeString, value)
		_node.DisplayName = value
	}
	if value, ok := epc.mutation.Name(); ok {
		_spec.SetField(entitlementplan.FieldName, field.TypeString, value)
		_node.Name = value
	}
	if value, ok := epc.mutation.Description(); ok {
		_spec.SetField(entitlementplan.FieldDescription, field.TypeString, value)
		_node.Description = value
	}
	if value, ok := epc.mutation.Version(); ok {
		_spec.SetField(entitlementplan.FieldVersion, field.TypeString, value)
		_node.Version = value
	}
	if value, ok := epc.mutation.Metadata(); ok {
		_spec.SetField(entitlementplan.FieldMetadata, field.TypeJSON, value)
		_node.Metadata = value
	}
	if value, ok := epc.mutation.StripeProductID(); ok {
		_spec.SetField(entitlementplan.FieldStripeProductID, field.TypeString, value)
		_node.StripeProductID = value
	}
	if value, ok := epc.mutation.StripePriceID(); ok {
		_spec.SetField(entitlementplan.FieldStripePriceID, field.TypeString, value)
		_node.StripePriceID = value
	}
	if nodes := epc.mutation.OwnerIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   entitlementplan.OwnerTable,
			Columns: []string{entitlementplan.OwnerColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(organization.FieldID, field.TypeString),
			},
		}
		edge.Schema = epc.schemaConfig.EntitlementPlan
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_node.OwnerID = nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := epc.mutation.EntitlementsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   entitlementplan.EntitlementsTable,
			Columns: []string{entitlementplan.EntitlementsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(entitlement.FieldID, field.TypeString),
			},
		}
		edge.Schema = epc.schemaConfig.Entitlement
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := epc.mutation.BaseFeaturesIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   entitlementplan.BaseFeaturesTable,
			Columns: entitlementplan.BaseFeaturesPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(feature.FieldID, field.TypeString),
			},
		}
		edge.Schema = epc.schemaConfig.EntitlementPlanFeature
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		createE := &EntitlementPlanFeatureCreate{config: epc.config, mutation: newEntitlementPlanFeatureMutation(epc.config, OpCreate)}
		_ = createE.defaults()
		_, specE := createE.createSpec()
		edge.Target.Fields = specE.Fields
		if specE.ID.Value != nil {
			edge.Target.Fields = append(edge.Target.Fields, specE.ID)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := epc.mutation.EventsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   entitlementplan.EventsTable,
			Columns: entitlementplan.EventsPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(event.FieldID, field.TypeString),
			},
		}
		edge.Schema = epc.schemaConfig.EntitlementPlanEvents
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := epc.mutation.FeaturesIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: true,
			Table:   entitlementplan.FeaturesTable,
			Columns: []string{entitlementplan.FeaturesColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(entitlementplanfeature.FieldID, field.TypeString),
			},
		}
		edge.Schema = epc.schemaConfig.EntitlementPlanFeature
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	return _node, _spec
}

// EntitlementPlanCreateBulk is the builder for creating many EntitlementPlan entities in bulk.
type EntitlementPlanCreateBulk struct {
	config
	err      error
	builders []*EntitlementPlanCreate
}

// Save creates the EntitlementPlan entities in the database.
func (epcb *EntitlementPlanCreateBulk) Save(ctx context.Context) ([]*EntitlementPlan, error) {
	if epcb.err != nil {
		return nil, epcb.err
	}
	specs := make([]*sqlgraph.CreateSpec, len(epcb.builders))
	nodes := make([]*EntitlementPlan, len(epcb.builders))
	mutators := make([]Mutator, len(epcb.builders))
	for i := range epcb.builders {
		func(i int, root context.Context) {
			builder := epcb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*EntitlementPlanMutation)
				if !ok {
					return nil, fmt.Errorf("unexpected mutation type %T", m)
				}
				if err := builder.check(); err != nil {
					return nil, err
				}
				builder.mutation = mutation
				var err error
				nodes[i], specs[i] = builder.createSpec()
				if i < len(mutators)-1 {
					_, err = mutators[i+1].Mutate(root, epcb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, epcb.driver, spec); err != nil {
						if sqlgraph.IsConstraintError(err) {
							err = &ConstraintError{msg: err.Error(), wrap: err}
						}
					}
				}
				if err != nil {
					return nil, err
				}
				mutation.id = &nodes[i].ID
				mutation.done = true
				return nodes[i], nil
			})
			for i := len(builder.hooks) - 1; i >= 0; i-- {
				mut = builder.hooks[i](mut)
			}
			mutators[i] = mut
		}(i, ctx)
	}
	if len(mutators) > 0 {
		if _, err := mutators[0].Mutate(ctx, epcb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (epcb *EntitlementPlanCreateBulk) SaveX(ctx context.Context) []*EntitlementPlan {
	v, err := epcb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (epcb *EntitlementPlanCreateBulk) Exec(ctx context.Context) error {
	_, err := epcb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (epcb *EntitlementPlanCreateBulk) ExecX(ctx context.Context) {
	if err := epcb.Exec(ctx); err != nil {
		panic(err)
	}
}
