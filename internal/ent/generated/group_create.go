// Code generated by ent, DO NOT EDIT.

package generated

import (
	"context"
	"errors"
	"fmt"
	"time"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/theopenlane/core/internal/ent/generated/event"
	"github.com/theopenlane/core/internal/ent/generated/file"
	"github.com/theopenlane/core/internal/ent/generated/group"
	"github.com/theopenlane/core/internal/ent/generated/groupmembership"
	"github.com/theopenlane/core/internal/ent/generated/groupsetting"
	"github.com/theopenlane/core/internal/ent/generated/integration"
	"github.com/theopenlane/core/internal/ent/generated/organization"
	"github.com/theopenlane/core/internal/ent/generated/task"
	"github.com/theopenlane/core/internal/ent/generated/user"
)

// GroupCreate is the builder for creating a Group entity.
type GroupCreate struct {
	config
	mutation *GroupMutation
	hooks    []Hook
}

// SetCreatedAt sets the "created_at" field.
func (gc *GroupCreate) SetCreatedAt(t time.Time) *GroupCreate {
	gc.mutation.SetCreatedAt(t)
	return gc
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (gc *GroupCreate) SetNillableCreatedAt(t *time.Time) *GroupCreate {
	if t != nil {
		gc.SetCreatedAt(*t)
	}
	return gc
}

// SetUpdatedAt sets the "updated_at" field.
func (gc *GroupCreate) SetUpdatedAt(t time.Time) *GroupCreate {
	gc.mutation.SetUpdatedAt(t)
	return gc
}

// SetNillableUpdatedAt sets the "updated_at" field if the given value is not nil.
func (gc *GroupCreate) SetNillableUpdatedAt(t *time.Time) *GroupCreate {
	if t != nil {
		gc.SetUpdatedAt(*t)
	}
	return gc
}

// SetCreatedBy sets the "created_by" field.
func (gc *GroupCreate) SetCreatedBy(s string) *GroupCreate {
	gc.mutation.SetCreatedBy(s)
	return gc
}

// SetNillableCreatedBy sets the "created_by" field if the given value is not nil.
func (gc *GroupCreate) SetNillableCreatedBy(s *string) *GroupCreate {
	if s != nil {
		gc.SetCreatedBy(*s)
	}
	return gc
}

// SetUpdatedBy sets the "updated_by" field.
func (gc *GroupCreate) SetUpdatedBy(s string) *GroupCreate {
	gc.mutation.SetUpdatedBy(s)
	return gc
}

// SetNillableUpdatedBy sets the "updated_by" field if the given value is not nil.
func (gc *GroupCreate) SetNillableUpdatedBy(s *string) *GroupCreate {
	if s != nil {
		gc.SetUpdatedBy(*s)
	}
	return gc
}

// SetDeletedAt sets the "deleted_at" field.
func (gc *GroupCreate) SetDeletedAt(t time.Time) *GroupCreate {
	gc.mutation.SetDeletedAt(t)
	return gc
}

// SetNillableDeletedAt sets the "deleted_at" field if the given value is not nil.
func (gc *GroupCreate) SetNillableDeletedAt(t *time.Time) *GroupCreate {
	if t != nil {
		gc.SetDeletedAt(*t)
	}
	return gc
}

// SetDeletedBy sets the "deleted_by" field.
func (gc *GroupCreate) SetDeletedBy(s string) *GroupCreate {
	gc.mutation.SetDeletedBy(s)
	return gc
}

// SetNillableDeletedBy sets the "deleted_by" field if the given value is not nil.
func (gc *GroupCreate) SetNillableDeletedBy(s *string) *GroupCreate {
	if s != nil {
		gc.SetDeletedBy(*s)
	}
	return gc
}

// SetMappingID sets the "mapping_id" field.
func (gc *GroupCreate) SetMappingID(s string) *GroupCreate {
	gc.mutation.SetMappingID(s)
	return gc
}

// SetNillableMappingID sets the "mapping_id" field if the given value is not nil.
func (gc *GroupCreate) SetNillableMappingID(s *string) *GroupCreate {
	if s != nil {
		gc.SetMappingID(*s)
	}
	return gc
}

// SetTags sets the "tags" field.
func (gc *GroupCreate) SetTags(s []string) *GroupCreate {
	gc.mutation.SetTags(s)
	return gc
}

// SetOwnerID sets the "owner_id" field.
func (gc *GroupCreate) SetOwnerID(s string) *GroupCreate {
	gc.mutation.SetOwnerID(s)
	return gc
}

// SetNillableOwnerID sets the "owner_id" field if the given value is not nil.
func (gc *GroupCreate) SetNillableOwnerID(s *string) *GroupCreate {
	if s != nil {
		gc.SetOwnerID(*s)
	}
	return gc
}

// SetName sets the "name" field.
func (gc *GroupCreate) SetName(s string) *GroupCreate {
	gc.mutation.SetName(s)
	return gc
}

// SetDescription sets the "description" field.
func (gc *GroupCreate) SetDescription(s string) *GroupCreate {
	gc.mutation.SetDescription(s)
	return gc
}

// SetNillableDescription sets the "description" field if the given value is not nil.
func (gc *GroupCreate) SetNillableDescription(s *string) *GroupCreate {
	if s != nil {
		gc.SetDescription(*s)
	}
	return gc
}

// SetGravatarLogoURL sets the "gravatar_logo_url" field.
func (gc *GroupCreate) SetGravatarLogoURL(s string) *GroupCreate {
	gc.mutation.SetGravatarLogoURL(s)
	return gc
}

// SetNillableGravatarLogoURL sets the "gravatar_logo_url" field if the given value is not nil.
func (gc *GroupCreate) SetNillableGravatarLogoURL(s *string) *GroupCreate {
	if s != nil {
		gc.SetGravatarLogoURL(*s)
	}
	return gc
}

// SetLogoURL sets the "logo_url" field.
func (gc *GroupCreate) SetLogoURL(s string) *GroupCreate {
	gc.mutation.SetLogoURL(s)
	return gc
}

// SetNillableLogoURL sets the "logo_url" field if the given value is not nil.
func (gc *GroupCreate) SetNillableLogoURL(s *string) *GroupCreate {
	if s != nil {
		gc.SetLogoURL(*s)
	}
	return gc
}

// SetDisplayName sets the "display_name" field.
func (gc *GroupCreate) SetDisplayName(s string) *GroupCreate {
	gc.mutation.SetDisplayName(s)
	return gc
}

// SetNillableDisplayName sets the "display_name" field if the given value is not nil.
func (gc *GroupCreate) SetNillableDisplayName(s *string) *GroupCreate {
	if s != nil {
		gc.SetDisplayName(*s)
	}
	return gc
}

// SetID sets the "id" field.
func (gc *GroupCreate) SetID(s string) *GroupCreate {
	gc.mutation.SetID(s)
	return gc
}

// SetNillableID sets the "id" field if the given value is not nil.
func (gc *GroupCreate) SetNillableID(s *string) *GroupCreate {
	if s != nil {
		gc.SetID(*s)
	}
	return gc
}

// SetOwner sets the "owner" edge to the Organization entity.
func (gc *GroupCreate) SetOwner(o *Organization) *GroupCreate {
	return gc.SetOwnerID(o.ID)
}

// SetSettingID sets the "setting" edge to the GroupSetting entity by ID.
func (gc *GroupCreate) SetSettingID(id string) *GroupCreate {
	gc.mutation.SetSettingID(id)
	return gc
}

// SetSetting sets the "setting" edge to the GroupSetting entity.
func (gc *GroupCreate) SetSetting(g *GroupSetting) *GroupCreate {
	return gc.SetSettingID(g.ID)
}

// AddUserIDs adds the "users" edge to the User entity by IDs.
func (gc *GroupCreate) AddUserIDs(ids ...string) *GroupCreate {
	gc.mutation.AddUserIDs(ids...)
	return gc
}

// AddUsers adds the "users" edges to the User entity.
func (gc *GroupCreate) AddUsers(u ...*User) *GroupCreate {
	ids := make([]string, len(u))
	for i := range u {
		ids[i] = u[i].ID
	}
	return gc.AddUserIDs(ids...)
}

// AddEventIDs adds the "events" edge to the Event entity by IDs.
func (gc *GroupCreate) AddEventIDs(ids ...string) *GroupCreate {
	gc.mutation.AddEventIDs(ids...)
	return gc
}

// AddEvents adds the "events" edges to the Event entity.
func (gc *GroupCreate) AddEvents(e ...*Event) *GroupCreate {
	ids := make([]string, len(e))
	for i := range e {
		ids[i] = e[i].ID
	}
	return gc.AddEventIDs(ids...)
}

// AddIntegrationIDs adds the "integrations" edge to the Integration entity by IDs.
func (gc *GroupCreate) AddIntegrationIDs(ids ...string) *GroupCreate {
	gc.mutation.AddIntegrationIDs(ids...)
	return gc
}

// AddIntegrations adds the "integrations" edges to the Integration entity.
func (gc *GroupCreate) AddIntegrations(i ...*Integration) *GroupCreate {
	ids := make([]string, len(i))
	for j := range i {
		ids[j] = i[j].ID
	}
	return gc.AddIntegrationIDs(ids...)
}

// AddFileIDs adds the "files" edge to the File entity by IDs.
func (gc *GroupCreate) AddFileIDs(ids ...string) *GroupCreate {
	gc.mutation.AddFileIDs(ids...)
	return gc
}

// AddFiles adds the "files" edges to the File entity.
func (gc *GroupCreate) AddFiles(f ...*File) *GroupCreate {
	ids := make([]string, len(f))
	for i := range f {
		ids[i] = f[i].ID
	}
	return gc.AddFileIDs(ids...)
}

// AddTaskIDs adds the "tasks" edge to the Task entity by IDs.
func (gc *GroupCreate) AddTaskIDs(ids ...string) *GroupCreate {
	gc.mutation.AddTaskIDs(ids...)
	return gc
}

// AddTasks adds the "tasks" edges to the Task entity.
func (gc *GroupCreate) AddTasks(t ...*Task) *GroupCreate {
	ids := make([]string, len(t))
	for i := range t {
		ids[i] = t[i].ID
	}
	return gc.AddTaskIDs(ids...)
}

// AddMemberIDs adds the "members" edge to the GroupMembership entity by IDs.
func (gc *GroupCreate) AddMemberIDs(ids ...string) *GroupCreate {
	gc.mutation.AddMemberIDs(ids...)
	return gc
}

// AddMembers adds the "members" edges to the GroupMembership entity.
func (gc *GroupCreate) AddMembers(g ...*GroupMembership) *GroupCreate {
	ids := make([]string, len(g))
	for i := range g {
		ids[i] = g[i].ID
	}
	return gc.AddMemberIDs(ids...)
}

// Mutation returns the GroupMutation object of the builder.
func (gc *GroupCreate) Mutation() *GroupMutation {
	return gc.mutation
}

// Save creates the Group in the database.
func (gc *GroupCreate) Save(ctx context.Context) (*Group, error) {
	if err := gc.defaults(); err != nil {
		return nil, err
	}
	return withHooks(ctx, gc.sqlSave, gc.mutation, gc.hooks)
}

// SaveX calls Save and panics if Save returns an error.
func (gc *GroupCreate) SaveX(ctx context.Context) *Group {
	v, err := gc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (gc *GroupCreate) Exec(ctx context.Context) error {
	_, err := gc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (gc *GroupCreate) ExecX(ctx context.Context) {
	if err := gc.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (gc *GroupCreate) defaults() error {
	if _, ok := gc.mutation.CreatedAt(); !ok {
		if group.DefaultCreatedAt == nil {
			return fmt.Errorf("generated: uninitialized group.DefaultCreatedAt (forgotten import generated/runtime?)")
		}
		v := group.DefaultCreatedAt()
		gc.mutation.SetCreatedAt(v)
	}
	if _, ok := gc.mutation.UpdatedAt(); !ok {
		if group.DefaultUpdatedAt == nil {
			return fmt.Errorf("generated: uninitialized group.DefaultUpdatedAt (forgotten import generated/runtime?)")
		}
		v := group.DefaultUpdatedAt()
		gc.mutation.SetUpdatedAt(v)
	}
	if _, ok := gc.mutation.MappingID(); !ok {
		if group.DefaultMappingID == nil {
			return fmt.Errorf("generated: uninitialized group.DefaultMappingID (forgotten import generated/runtime?)")
		}
		v := group.DefaultMappingID()
		gc.mutation.SetMappingID(v)
	}
	if _, ok := gc.mutation.Tags(); !ok {
		v := group.DefaultTags
		gc.mutation.SetTags(v)
	}
	if _, ok := gc.mutation.DisplayName(); !ok {
		v := group.DefaultDisplayName
		gc.mutation.SetDisplayName(v)
	}
	if _, ok := gc.mutation.ID(); !ok {
		if group.DefaultID == nil {
			return fmt.Errorf("generated: uninitialized group.DefaultID (forgotten import generated/runtime?)")
		}
		v := group.DefaultID()
		gc.mutation.SetID(v)
	}
	return nil
}

// check runs all checks and user-defined validators on the builder.
func (gc *GroupCreate) check() error {
	if _, ok := gc.mutation.MappingID(); !ok {
		return &ValidationError{Name: "mapping_id", err: errors.New(`generated: missing required field "Group.mapping_id"`)}
	}
	if v, ok := gc.mutation.OwnerID(); ok {
		if err := group.OwnerIDValidator(v); err != nil {
			return &ValidationError{Name: "owner_id", err: fmt.Errorf(`generated: validator failed for field "Group.owner_id": %w`, err)}
		}
	}
	if _, ok := gc.mutation.Name(); !ok {
		return &ValidationError{Name: "name", err: errors.New(`generated: missing required field "Group.name"`)}
	}
	if v, ok := gc.mutation.Name(); ok {
		if err := group.NameValidator(v); err != nil {
			return &ValidationError{Name: "name", err: fmt.Errorf(`generated: validator failed for field "Group.name": %w`, err)}
		}
	}
	if _, ok := gc.mutation.DisplayName(); !ok {
		return &ValidationError{Name: "display_name", err: errors.New(`generated: missing required field "Group.display_name"`)}
	}
	if v, ok := gc.mutation.DisplayName(); ok {
		if err := group.DisplayNameValidator(v); err != nil {
			return &ValidationError{Name: "display_name", err: fmt.Errorf(`generated: validator failed for field "Group.display_name": %w`, err)}
		}
	}
	if len(gc.mutation.SettingIDs()) == 0 {
		return &ValidationError{Name: "setting", err: errors.New(`generated: missing required edge "Group.setting"`)}
	}
	return nil
}

func (gc *GroupCreate) sqlSave(ctx context.Context) (*Group, error) {
	if err := gc.check(); err != nil {
		return nil, err
	}
	_node, _spec := gc.createSpec()
	if err := sqlgraph.CreateNode(ctx, gc.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	if _spec.ID.Value != nil {
		if id, ok := _spec.ID.Value.(string); ok {
			_node.ID = id
		} else {
			return nil, fmt.Errorf("unexpected Group.ID type: %T", _spec.ID.Value)
		}
	}
	gc.mutation.id = &_node.ID
	gc.mutation.done = true
	return _node, nil
}

func (gc *GroupCreate) createSpec() (*Group, *sqlgraph.CreateSpec) {
	var (
		_node = &Group{config: gc.config}
		_spec = sqlgraph.NewCreateSpec(group.Table, sqlgraph.NewFieldSpec(group.FieldID, field.TypeString))
	)
	_spec.Schema = gc.schemaConfig.Group
	if id, ok := gc.mutation.ID(); ok {
		_node.ID = id
		_spec.ID.Value = id
	}
	if value, ok := gc.mutation.CreatedAt(); ok {
		_spec.SetField(group.FieldCreatedAt, field.TypeTime, value)
		_node.CreatedAt = value
	}
	if value, ok := gc.mutation.UpdatedAt(); ok {
		_spec.SetField(group.FieldUpdatedAt, field.TypeTime, value)
		_node.UpdatedAt = value
	}
	if value, ok := gc.mutation.CreatedBy(); ok {
		_spec.SetField(group.FieldCreatedBy, field.TypeString, value)
		_node.CreatedBy = value
	}
	if value, ok := gc.mutation.UpdatedBy(); ok {
		_spec.SetField(group.FieldUpdatedBy, field.TypeString, value)
		_node.UpdatedBy = value
	}
	if value, ok := gc.mutation.DeletedAt(); ok {
		_spec.SetField(group.FieldDeletedAt, field.TypeTime, value)
		_node.DeletedAt = value
	}
	if value, ok := gc.mutation.DeletedBy(); ok {
		_spec.SetField(group.FieldDeletedBy, field.TypeString, value)
		_node.DeletedBy = value
	}
	if value, ok := gc.mutation.MappingID(); ok {
		_spec.SetField(group.FieldMappingID, field.TypeString, value)
		_node.MappingID = value
	}
	if value, ok := gc.mutation.Tags(); ok {
		_spec.SetField(group.FieldTags, field.TypeJSON, value)
		_node.Tags = value
	}
	if value, ok := gc.mutation.Name(); ok {
		_spec.SetField(group.FieldName, field.TypeString, value)
		_node.Name = value
	}
	if value, ok := gc.mutation.Description(); ok {
		_spec.SetField(group.FieldDescription, field.TypeString, value)
		_node.Description = value
	}
	if value, ok := gc.mutation.GravatarLogoURL(); ok {
		_spec.SetField(group.FieldGravatarLogoURL, field.TypeString, value)
		_node.GravatarLogoURL = value
	}
	if value, ok := gc.mutation.LogoURL(); ok {
		_spec.SetField(group.FieldLogoURL, field.TypeString, value)
		_node.LogoURL = value
	}
	if value, ok := gc.mutation.DisplayName(); ok {
		_spec.SetField(group.FieldDisplayName, field.TypeString, value)
		_node.DisplayName = value
	}
	if nodes := gc.mutation.OwnerIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   group.OwnerTable,
			Columns: []string{group.OwnerColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(organization.FieldID, field.TypeString),
			},
		}
		edge.Schema = gc.schemaConfig.Group
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_node.OwnerID = nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := gc.mutation.SettingIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2O,
			Inverse: false,
			Table:   group.SettingTable,
			Columns: []string{group.SettingColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(groupsetting.FieldID, field.TypeString),
			},
		}
		edge.Schema = gc.schemaConfig.GroupSetting
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := gc.mutation.UsersIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   group.UsersTable,
			Columns: group.UsersPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(user.FieldID, field.TypeString),
			},
		}
		edge.Schema = gc.schemaConfig.GroupMembership
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		createE := &GroupMembershipCreate{config: gc.config, mutation: newGroupMembershipMutation(gc.config, OpCreate)}
		_ = createE.defaults()
		_, specE := createE.createSpec()
		edge.Target.Fields = specE.Fields
		if specE.ID.Value != nil {
			edge.Target.Fields = append(edge.Target.Fields, specE.ID)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := gc.mutation.EventsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   group.EventsTable,
			Columns: group.EventsPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(event.FieldID, field.TypeString),
			},
		}
		edge.Schema = gc.schemaConfig.GroupEvents
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := gc.mutation.IntegrationsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   group.IntegrationsTable,
			Columns: []string{group.IntegrationsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(integration.FieldID, field.TypeString),
			},
		}
		edge.Schema = gc.schemaConfig.Integration
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := gc.mutation.FilesIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   group.FilesTable,
			Columns: group.FilesPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(file.FieldID, field.TypeString),
			},
		}
		edge.Schema = gc.schemaConfig.GroupFiles
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := gc.mutation.TasksIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   group.TasksTable,
			Columns: group.TasksPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(task.FieldID, field.TypeString),
			},
		}
		edge.Schema = gc.schemaConfig.GroupTasks
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := gc.mutation.MembersIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: true,
			Table:   group.MembersTable,
			Columns: []string{group.MembersColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(groupmembership.FieldID, field.TypeString),
			},
		}
		edge.Schema = gc.schemaConfig.GroupMembership
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	return _node, _spec
}

// GroupCreateBulk is the builder for creating many Group entities in bulk.
type GroupCreateBulk struct {
	config
	err      error
	builders []*GroupCreate
}

// Save creates the Group entities in the database.
func (gcb *GroupCreateBulk) Save(ctx context.Context) ([]*Group, error) {
	if gcb.err != nil {
		return nil, gcb.err
	}
	specs := make([]*sqlgraph.CreateSpec, len(gcb.builders))
	nodes := make([]*Group, len(gcb.builders))
	mutators := make([]Mutator, len(gcb.builders))
	for i := range gcb.builders {
		func(i int, root context.Context) {
			builder := gcb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*GroupMutation)
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
					_, err = mutators[i+1].Mutate(root, gcb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, gcb.driver, spec); err != nil {
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
		if _, err := mutators[0].Mutate(ctx, gcb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (gcb *GroupCreateBulk) SaveX(ctx context.Context) []*Group {
	v, err := gcb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (gcb *GroupCreateBulk) Exec(ctx context.Context) error {
	_, err := gcb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (gcb *GroupCreateBulk) ExecX(ctx context.Context) {
	if err := gcb.Exec(ctx); err != nil {
		panic(err)
	}
}
