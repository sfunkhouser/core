// Code generated by ent, DO NOT EDIT.

package generated

import (
	"context"
	"errors"
	"fmt"
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/dialect/sql/sqljson"
	"entgo.io/ent/schema/field"
	"github.com/theopenlane/core/internal/ent/generated/internalpolicyhistory"
	"github.com/theopenlane/core/internal/ent/generated/predicate"

	"github.com/theopenlane/core/internal/ent/generated/internal"
)

// InternalPolicyHistoryUpdate is the builder for updating InternalPolicyHistory entities.
type InternalPolicyHistoryUpdate struct {
	config
	hooks     []Hook
	mutation  *InternalPolicyHistoryMutation
	modifiers []func(*sql.UpdateBuilder)
}

// Where appends a list predicates to the InternalPolicyHistoryUpdate builder.
func (iphu *InternalPolicyHistoryUpdate) Where(ps ...predicate.InternalPolicyHistory) *InternalPolicyHistoryUpdate {
	iphu.mutation.Where(ps...)
	return iphu
}

// SetUpdatedAt sets the "updated_at" field.
func (iphu *InternalPolicyHistoryUpdate) SetUpdatedAt(t time.Time) *InternalPolicyHistoryUpdate {
	iphu.mutation.SetUpdatedAt(t)
	return iphu
}

// ClearUpdatedAt clears the value of the "updated_at" field.
func (iphu *InternalPolicyHistoryUpdate) ClearUpdatedAt() *InternalPolicyHistoryUpdate {
	iphu.mutation.ClearUpdatedAt()
	return iphu
}

// SetUpdatedBy sets the "updated_by" field.
func (iphu *InternalPolicyHistoryUpdate) SetUpdatedBy(s string) *InternalPolicyHistoryUpdate {
	iphu.mutation.SetUpdatedBy(s)
	return iphu
}

// SetNillableUpdatedBy sets the "updated_by" field if the given value is not nil.
func (iphu *InternalPolicyHistoryUpdate) SetNillableUpdatedBy(s *string) *InternalPolicyHistoryUpdate {
	if s != nil {
		iphu.SetUpdatedBy(*s)
	}
	return iphu
}

// ClearUpdatedBy clears the value of the "updated_by" field.
func (iphu *InternalPolicyHistoryUpdate) ClearUpdatedBy() *InternalPolicyHistoryUpdate {
	iphu.mutation.ClearUpdatedBy()
	return iphu
}

// SetDeletedAt sets the "deleted_at" field.
func (iphu *InternalPolicyHistoryUpdate) SetDeletedAt(t time.Time) *InternalPolicyHistoryUpdate {
	iphu.mutation.SetDeletedAt(t)
	return iphu
}

// SetNillableDeletedAt sets the "deleted_at" field if the given value is not nil.
func (iphu *InternalPolicyHistoryUpdate) SetNillableDeletedAt(t *time.Time) *InternalPolicyHistoryUpdate {
	if t != nil {
		iphu.SetDeletedAt(*t)
	}
	return iphu
}

// ClearDeletedAt clears the value of the "deleted_at" field.
func (iphu *InternalPolicyHistoryUpdate) ClearDeletedAt() *InternalPolicyHistoryUpdate {
	iphu.mutation.ClearDeletedAt()
	return iphu
}

// SetDeletedBy sets the "deleted_by" field.
func (iphu *InternalPolicyHistoryUpdate) SetDeletedBy(s string) *InternalPolicyHistoryUpdate {
	iphu.mutation.SetDeletedBy(s)
	return iphu
}

// SetNillableDeletedBy sets the "deleted_by" field if the given value is not nil.
func (iphu *InternalPolicyHistoryUpdate) SetNillableDeletedBy(s *string) *InternalPolicyHistoryUpdate {
	if s != nil {
		iphu.SetDeletedBy(*s)
	}
	return iphu
}

// ClearDeletedBy clears the value of the "deleted_by" field.
func (iphu *InternalPolicyHistoryUpdate) ClearDeletedBy() *InternalPolicyHistoryUpdate {
	iphu.mutation.ClearDeletedBy()
	return iphu
}

// SetTags sets the "tags" field.
func (iphu *InternalPolicyHistoryUpdate) SetTags(s []string) *InternalPolicyHistoryUpdate {
	iphu.mutation.SetTags(s)
	return iphu
}

// AppendTags appends s to the "tags" field.
func (iphu *InternalPolicyHistoryUpdate) AppendTags(s []string) *InternalPolicyHistoryUpdate {
	iphu.mutation.AppendTags(s)
	return iphu
}

// ClearTags clears the value of the "tags" field.
func (iphu *InternalPolicyHistoryUpdate) ClearTags() *InternalPolicyHistoryUpdate {
	iphu.mutation.ClearTags()
	return iphu
}

// SetName sets the "name" field.
func (iphu *InternalPolicyHistoryUpdate) SetName(s string) *InternalPolicyHistoryUpdate {
	iphu.mutation.SetName(s)
	return iphu
}

// SetNillableName sets the "name" field if the given value is not nil.
func (iphu *InternalPolicyHistoryUpdate) SetNillableName(s *string) *InternalPolicyHistoryUpdate {
	if s != nil {
		iphu.SetName(*s)
	}
	return iphu
}

// SetDescription sets the "description" field.
func (iphu *InternalPolicyHistoryUpdate) SetDescription(s string) *InternalPolicyHistoryUpdate {
	iphu.mutation.SetDescription(s)
	return iphu
}

// SetNillableDescription sets the "description" field if the given value is not nil.
func (iphu *InternalPolicyHistoryUpdate) SetNillableDescription(s *string) *InternalPolicyHistoryUpdate {
	if s != nil {
		iphu.SetDescription(*s)
	}
	return iphu
}

// SetStatus sets the "status" field.
func (iphu *InternalPolicyHistoryUpdate) SetStatus(s string) *InternalPolicyHistoryUpdate {
	iphu.mutation.SetStatus(s)
	return iphu
}

// SetNillableStatus sets the "status" field if the given value is not nil.
func (iphu *InternalPolicyHistoryUpdate) SetNillableStatus(s *string) *InternalPolicyHistoryUpdate {
	if s != nil {
		iphu.SetStatus(*s)
	}
	return iphu
}

// ClearStatus clears the value of the "status" field.
func (iphu *InternalPolicyHistoryUpdate) ClearStatus() *InternalPolicyHistoryUpdate {
	iphu.mutation.ClearStatus()
	return iphu
}

// SetPolicyType sets the "policy_type" field.
func (iphu *InternalPolicyHistoryUpdate) SetPolicyType(s string) *InternalPolicyHistoryUpdate {
	iphu.mutation.SetPolicyType(s)
	return iphu
}

// SetNillablePolicyType sets the "policy_type" field if the given value is not nil.
func (iphu *InternalPolicyHistoryUpdate) SetNillablePolicyType(s *string) *InternalPolicyHistoryUpdate {
	if s != nil {
		iphu.SetPolicyType(*s)
	}
	return iphu
}

// ClearPolicyType clears the value of the "policy_type" field.
func (iphu *InternalPolicyHistoryUpdate) ClearPolicyType() *InternalPolicyHistoryUpdate {
	iphu.mutation.ClearPolicyType()
	return iphu
}

// SetVersion sets the "version" field.
func (iphu *InternalPolicyHistoryUpdate) SetVersion(s string) *InternalPolicyHistoryUpdate {
	iphu.mutation.SetVersion(s)
	return iphu
}

// SetNillableVersion sets the "version" field if the given value is not nil.
func (iphu *InternalPolicyHistoryUpdate) SetNillableVersion(s *string) *InternalPolicyHistoryUpdate {
	if s != nil {
		iphu.SetVersion(*s)
	}
	return iphu
}

// ClearVersion clears the value of the "version" field.
func (iphu *InternalPolicyHistoryUpdate) ClearVersion() *InternalPolicyHistoryUpdate {
	iphu.mutation.ClearVersion()
	return iphu
}

// SetPurposeAndScope sets the "purpose_and_scope" field.
func (iphu *InternalPolicyHistoryUpdate) SetPurposeAndScope(s string) *InternalPolicyHistoryUpdate {
	iphu.mutation.SetPurposeAndScope(s)
	return iphu
}

// SetNillablePurposeAndScope sets the "purpose_and_scope" field if the given value is not nil.
func (iphu *InternalPolicyHistoryUpdate) SetNillablePurposeAndScope(s *string) *InternalPolicyHistoryUpdate {
	if s != nil {
		iphu.SetPurposeAndScope(*s)
	}
	return iphu
}

// ClearPurposeAndScope clears the value of the "purpose_and_scope" field.
func (iphu *InternalPolicyHistoryUpdate) ClearPurposeAndScope() *InternalPolicyHistoryUpdate {
	iphu.mutation.ClearPurposeAndScope()
	return iphu
}

// SetBackground sets the "background" field.
func (iphu *InternalPolicyHistoryUpdate) SetBackground(s string) *InternalPolicyHistoryUpdate {
	iphu.mutation.SetBackground(s)
	return iphu
}

// SetNillableBackground sets the "background" field if the given value is not nil.
func (iphu *InternalPolicyHistoryUpdate) SetNillableBackground(s *string) *InternalPolicyHistoryUpdate {
	if s != nil {
		iphu.SetBackground(*s)
	}
	return iphu
}

// ClearBackground clears the value of the "background" field.
func (iphu *InternalPolicyHistoryUpdate) ClearBackground() *InternalPolicyHistoryUpdate {
	iphu.mutation.ClearBackground()
	return iphu
}

// SetDetails sets the "details" field.
func (iphu *InternalPolicyHistoryUpdate) SetDetails(m map[string]interface{}) *InternalPolicyHistoryUpdate {
	iphu.mutation.SetDetails(m)
	return iphu
}

// ClearDetails clears the value of the "details" field.
func (iphu *InternalPolicyHistoryUpdate) ClearDetails() *InternalPolicyHistoryUpdate {
	iphu.mutation.ClearDetails()
	return iphu
}

// Mutation returns the InternalPolicyHistoryMutation object of the builder.
func (iphu *InternalPolicyHistoryUpdate) Mutation() *InternalPolicyHistoryMutation {
	return iphu.mutation
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (iphu *InternalPolicyHistoryUpdate) Save(ctx context.Context) (int, error) {
	iphu.defaults()
	return withHooks(ctx, iphu.sqlSave, iphu.mutation, iphu.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (iphu *InternalPolicyHistoryUpdate) SaveX(ctx context.Context) int {
	affected, err := iphu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (iphu *InternalPolicyHistoryUpdate) Exec(ctx context.Context) error {
	_, err := iphu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (iphu *InternalPolicyHistoryUpdate) ExecX(ctx context.Context) {
	if err := iphu.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (iphu *InternalPolicyHistoryUpdate) defaults() {
	if _, ok := iphu.mutation.UpdatedAt(); !ok && !iphu.mutation.UpdatedAtCleared() {
		v := internalpolicyhistory.UpdateDefaultUpdatedAt()
		iphu.mutation.SetUpdatedAt(v)
	}
}

// Modify adds a statement modifier for attaching custom logic to the UPDATE statement.
func (iphu *InternalPolicyHistoryUpdate) Modify(modifiers ...func(u *sql.UpdateBuilder)) *InternalPolicyHistoryUpdate {
	iphu.modifiers = append(iphu.modifiers, modifiers...)
	return iphu
}

func (iphu *InternalPolicyHistoryUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := sqlgraph.NewUpdateSpec(internalpolicyhistory.Table, internalpolicyhistory.Columns, sqlgraph.NewFieldSpec(internalpolicyhistory.FieldID, field.TypeString))
	if ps := iphu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if iphu.mutation.RefCleared() {
		_spec.ClearField(internalpolicyhistory.FieldRef, field.TypeString)
	}
	if iphu.mutation.CreatedAtCleared() {
		_spec.ClearField(internalpolicyhistory.FieldCreatedAt, field.TypeTime)
	}
	if value, ok := iphu.mutation.UpdatedAt(); ok {
		_spec.SetField(internalpolicyhistory.FieldUpdatedAt, field.TypeTime, value)
	}
	if iphu.mutation.UpdatedAtCleared() {
		_spec.ClearField(internalpolicyhistory.FieldUpdatedAt, field.TypeTime)
	}
	if iphu.mutation.CreatedByCleared() {
		_spec.ClearField(internalpolicyhistory.FieldCreatedBy, field.TypeString)
	}
	if value, ok := iphu.mutation.UpdatedBy(); ok {
		_spec.SetField(internalpolicyhistory.FieldUpdatedBy, field.TypeString, value)
	}
	if iphu.mutation.UpdatedByCleared() {
		_spec.ClearField(internalpolicyhistory.FieldUpdatedBy, field.TypeString)
	}
	if value, ok := iphu.mutation.DeletedAt(); ok {
		_spec.SetField(internalpolicyhistory.FieldDeletedAt, field.TypeTime, value)
	}
	if iphu.mutation.DeletedAtCleared() {
		_spec.ClearField(internalpolicyhistory.FieldDeletedAt, field.TypeTime)
	}
	if value, ok := iphu.mutation.DeletedBy(); ok {
		_spec.SetField(internalpolicyhistory.FieldDeletedBy, field.TypeString, value)
	}
	if iphu.mutation.DeletedByCleared() {
		_spec.ClearField(internalpolicyhistory.FieldDeletedBy, field.TypeString)
	}
	if value, ok := iphu.mutation.Tags(); ok {
		_spec.SetField(internalpolicyhistory.FieldTags, field.TypeJSON, value)
	}
	if value, ok := iphu.mutation.AppendedTags(); ok {
		_spec.AddModifier(func(u *sql.UpdateBuilder) {
			sqljson.Append(u, internalpolicyhistory.FieldTags, value)
		})
	}
	if iphu.mutation.TagsCleared() {
		_spec.ClearField(internalpolicyhistory.FieldTags, field.TypeJSON)
	}
	if value, ok := iphu.mutation.Name(); ok {
		_spec.SetField(internalpolicyhistory.FieldName, field.TypeString, value)
	}
	if value, ok := iphu.mutation.Description(); ok {
		_spec.SetField(internalpolicyhistory.FieldDescription, field.TypeString, value)
	}
	if value, ok := iphu.mutation.Status(); ok {
		_spec.SetField(internalpolicyhistory.FieldStatus, field.TypeString, value)
	}
	if iphu.mutation.StatusCleared() {
		_spec.ClearField(internalpolicyhistory.FieldStatus, field.TypeString)
	}
	if value, ok := iphu.mutation.PolicyType(); ok {
		_spec.SetField(internalpolicyhistory.FieldPolicyType, field.TypeString, value)
	}
	if iphu.mutation.PolicyTypeCleared() {
		_spec.ClearField(internalpolicyhistory.FieldPolicyType, field.TypeString)
	}
	if value, ok := iphu.mutation.Version(); ok {
		_spec.SetField(internalpolicyhistory.FieldVersion, field.TypeString, value)
	}
	if iphu.mutation.VersionCleared() {
		_spec.ClearField(internalpolicyhistory.FieldVersion, field.TypeString)
	}
	if value, ok := iphu.mutation.PurposeAndScope(); ok {
		_spec.SetField(internalpolicyhistory.FieldPurposeAndScope, field.TypeString, value)
	}
	if iphu.mutation.PurposeAndScopeCleared() {
		_spec.ClearField(internalpolicyhistory.FieldPurposeAndScope, field.TypeString)
	}
	if value, ok := iphu.mutation.Background(); ok {
		_spec.SetField(internalpolicyhistory.FieldBackground, field.TypeString, value)
	}
	if iphu.mutation.BackgroundCleared() {
		_spec.ClearField(internalpolicyhistory.FieldBackground, field.TypeString)
	}
	if value, ok := iphu.mutation.Details(); ok {
		_spec.SetField(internalpolicyhistory.FieldDetails, field.TypeJSON, value)
	}
	if iphu.mutation.DetailsCleared() {
		_spec.ClearField(internalpolicyhistory.FieldDetails, field.TypeJSON)
	}
	_spec.Node.Schema = iphu.schemaConfig.InternalPolicyHistory
	ctx = internal.NewSchemaConfigContext(ctx, iphu.schemaConfig)
	_spec.AddModifiers(iphu.modifiers...)
	if n, err = sqlgraph.UpdateNodes(ctx, iphu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{internalpolicyhistory.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	iphu.mutation.done = true
	return n, nil
}

// InternalPolicyHistoryUpdateOne is the builder for updating a single InternalPolicyHistory entity.
type InternalPolicyHistoryUpdateOne struct {
	config
	fields    []string
	hooks     []Hook
	mutation  *InternalPolicyHistoryMutation
	modifiers []func(*sql.UpdateBuilder)
}

// SetUpdatedAt sets the "updated_at" field.
func (iphuo *InternalPolicyHistoryUpdateOne) SetUpdatedAt(t time.Time) *InternalPolicyHistoryUpdateOne {
	iphuo.mutation.SetUpdatedAt(t)
	return iphuo
}

// ClearUpdatedAt clears the value of the "updated_at" field.
func (iphuo *InternalPolicyHistoryUpdateOne) ClearUpdatedAt() *InternalPolicyHistoryUpdateOne {
	iphuo.mutation.ClearUpdatedAt()
	return iphuo
}

// SetUpdatedBy sets the "updated_by" field.
func (iphuo *InternalPolicyHistoryUpdateOne) SetUpdatedBy(s string) *InternalPolicyHistoryUpdateOne {
	iphuo.mutation.SetUpdatedBy(s)
	return iphuo
}

// SetNillableUpdatedBy sets the "updated_by" field if the given value is not nil.
func (iphuo *InternalPolicyHistoryUpdateOne) SetNillableUpdatedBy(s *string) *InternalPolicyHistoryUpdateOne {
	if s != nil {
		iphuo.SetUpdatedBy(*s)
	}
	return iphuo
}

// ClearUpdatedBy clears the value of the "updated_by" field.
func (iphuo *InternalPolicyHistoryUpdateOne) ClearUpdatedBy() *InternalPolicyHistoryUpdateOne {
	iphuo.mutation.ClearUpdatedBy()
	return iphuo
}

// SetDeletedAt sets the "deleted_at" field.
func (iphuo *InternalPolicyHistoryUpdateOne) SetDeletedAt(t time.Time) *InternalPolicyHistoryUpdateOne {
	iphuo.mutation.SetDeletedAt(t)
	return iphuo
}

// SetNillableDeletedAt sets the "deleted_at" field if the given value is not nil.
func (iphuo *InternalPolicyHistoryUpdateOne) SetNillableDeletedAt(t *time.Time) *InternalPolicyHistoryUpdateOne {
	if t != nil {
		iphuo.SetDeletedAt(*t)
	}
	return iphuo
}

// ClearDeletedAt clears the value of the "deleted_at" field.
func (iphuo *InternalPolicyHistoryUpdateOne) ClearDeletedAt() *InternalPolicyHistoryUpdateOne {
	iphuo.mutation.ClearDeletedAt()
	return iphuo
}

// SetDeletedBy sets the "deleted_by" field.
func (iphuo *InternalPolicyHistoryUpdateOne) SetDeletedBy(s string) *InternalPolicyHistoryUpdateOne {
	iphuo.mutation.SetDeletedBy(s)
	return iphuo
}

// SetNillableDeletedBy sets the "deleted_by" field if the given value is not nil.
func (iphuo *InternalPolicyHistoryUpdateOne) SetNillableDeletedBy(s *string) *InternalPolicyHistoryUpdateOne {
	if s != nil {
		iphuo.SetDeletedBy(*s)
	}
	return iphuo
}

// ClearDeletedBy clears the value of the "deleted_by" field.
func (iphuo *InternalPolicyHistoryUpdateOne) ClearDeletedBy() *InternalPolicyHistoryUpdateOne {
	iphuo.mutation.ClearDeletedBy()
	return iphuo
}

// SetTags sets the "tags" field.
func (iphuo *InternalPolicyHistoryUpdateOne) SetTags(s []string) *InternalPolicyHistoryUpdateOne {
	iphuo.mutation.SetTags(s)
	return iphuo
}

// AppendTags appends s to the "tags" field.
func (iphuo *InternalPolicyHistoryUpdateOne) AppendTags(s []string) *InternalPolicyHistoryUpdateOne {
	iphuo.mutation.AppendTags(s)
	return iphuo
}

// ClearTags clears the value of the "tags" field.
func (iphuo *InternalPolicyHistoryUpdateOne) ClearTags() *InternalPolicyHistoryUpdateOne {
	iphuo.mutation.ClearTags()
	return iphuo
}

// SetName sets the "name" field.
func (iphuo *InternalPolicyHistoryUpdateOne) SetName(s string) *InternalPolicyHistoryUpdateOne {
	iphuo.mutation.SetName(s)
	return iphuo
}

// SetNillableName sets the "name" field if the given value is not nil.
func (iphuo *InternalPolicyHistoryUpdateOne) SetNillableName(s *string) *InternalPolicyHistoryUpdateOne {
	if s != nil {
		iphuo.SetName(*s)
	}
	return iphuo
}

// SetDescription sets the "description" field.
func (iphuo *InternalPolicyHistoryUpdateOne) SetDescription(s string) *InternalPolicyHistoryUpdateOne {
	iphuo.mutation.SetDescription(s)
	return iphuo
}

// SetNillableDescription sets the "description" field if the given value is not nil.
func (iphuo *InternalPolicyHistoryUpdateOne) SetNillableDescription(s *string) *InternalPolicyHistoryUpdateOne {
	if s != nil {
		iphuo.SetDescription(*s)
	}
	return iphuo
}

// SetStatus sets the "status" field.
func (iphuo *InternalPolicyHistoryUpdateOne) SetStatus(s string) *InternalPolicyHistoryUpdateOne {
	iphuo.mutation.SetStatus(s)
	return iphuo
}

// SetNillableStatus sets the "status" field if the given value is not nil.
func (iphuo *InternalPolicyHistoryUpdateOne) SetNillableStatus(s *string) *InternalPolicyHistoryUpdateOne {
	if s != nil {
		iphuo.SetStatus(*s)
	}
	return iphuo
}

// ClearStatus clears the value of the "status" field.
func (iphuo *InternalPolicyHistoryUpdateOne) ClearStatus() *InternalPolicyHistoryUpdateOne {
	iphuo.mutation.ClearStatus()
	return iphuo
}

// SetPolicyType sets the "policy_type" field.
func (iphuo *InternalPolicyHistoryUpdateOne) SetPolicyType(s string) *InternalPolicyHistoryUpdateOne {
	iphuo.mutation.SetPolicyType(s)
	return iphuo
}

// SetNillablePolicyType sets the "policy_type" field if the given value is not nil.
func (iphuo *InternalPolicyHistoryUpdateOne) SetNillablePolicyType(s *string) *InternalPolicyHistoryUpdateOne {
	if s != nil {
		iphuo.SetPolicyType(*s)
	}
	return iphuo
}

// ClearPolicyType clears the value of the "policy_type" field.
func (iphuo *InternalPolicyHistoryUpdateOne) ClearPolicyType() *InternalPolicyHistoryUpdateOne {
	iphuo.mutation.ClearPolicyType()
	return iphuo
}

// SetVersion sets the "version" field.
func (iphuo *InternalPolicyHistoryUpdateOne) SetVersion(s string) *InternalPolicyHistoryUpdateOne {
	iphuo.mutation.SetVersion(s)
	return iphuo
}

// SetNillableVersion sets the "version" field if the given value is not nil.
func (iphuo *InternalPolicyHistoryUpdateOne) SetNillableVersion(s *string) *InternalPolicyHistoryUpdateOne {
	if s != nil {
		iphuo.SetVersion(*s)
	}
	return iphuo
}

// ClearVersion clears the value of the "version" field.
func (iphuo *InternalPolicyHistoryUpdateOne) ClearVersion() *InternalPolicyHistoryUpdateOne {
	iphuo.mutation.ClearVersion()
	return iphuo
}

// SetPurposeAndScope sets the "purpose_and_scope" field.
func (iphuo *InternalPolicyHistoryUpdateOne) SetPurposeAndScope(s string) *InternalPolicyHistoryUpdateOne {
	iphuo.mutation.SetPurposeAndScope(s)
	return iphuo
}

// SetNillablePurposeAndScope sets the "purpose_and_scope" field if the given value is not nil.
func (iphuo *InternalPolicyHistoryUpdateOne) SetNillablePurposeAndScope(s *string) *InternalPolicyHistoryUpdateOne {
	if s != nil {
		iphuo.SetPurposeAndScope(*s)
	}
	return iphuo
}

// ClearPurposeAndScope clears the value of the "purpose_and_scope" field.
func (iphuo *InternalPolicyHistoryUpdateOne) ClearPurposeAndScope() *InternalPolicyHistoryUpdateOne {
	iphuo.mutation.ClearPurposeAndScope()
	return iphuo
}

// SetBackground sets the "background" field.
func (iphuo *InternalPolicyHistoryUpdateOne) SetBackground(s string) *InternalPolicyHistoryUpdateOne {
	iphuo.mutation.SetBackground(s)
	return iphuo
}

// SetNillableBackground sets the "background" field if the given value is not nil.
func (iphuo *InternalPolicyHistoryUpdateOne) SetNillableBackground(s *string) *InternalPolicyHistoryUpdateOne {
	if s != nil {
		iphuo.SetBackground(*s)
	}
	return iphuo
}

// ClearBackground clears the value of the "background" field.
func (iphuo *InternalPolicyHistoryUpdateOne) ClearBackground() *InternalPolicyHistoryUpdateOne {
	iphuo.mutation.ClearBackground()
	return iphuo
}

// SetDetails sets the "details" field.
func (iphuo *InternalPolicyHistoryUpdateOne) SetDetails(m map[string]interface{}) *InternalPolicyHistoryUpdateOne {
	iphuo.mutation.SetDetails(m)
	return iphuo
}

// ClearDetails clears the value of the "details" field.
func (iphuo *InternalPolicyHistoryUpdateOne) ClearDetails() *InternalPolicyHistoryUpdateOne {
	iphuo.mutation.ClearDetails()
	return iphuo
}

// Mutation returns the InternalPolicyHistoryMutation object of the builder.
func (iphuo *InternalPolicyHistoryUpdateOne) Mutation() *InternalPolicyHistoryMutation {
	return iphuo.mutation
}

// Where appends a list predicates to the InternalPolicyHistoryUpdate builder.
func (iphuo *InternalPolicyHistoryUpdateOne) Where(ps ...predicate.InternalPolicyHistory) *InternalPolicyHistoryUpdateOne {
	iphuo.mutation.Where(ps...)
	return iphuo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (iphuo *InternalPolicyHistoryUpdateOne) Select(field string, fields ...string) *InternalPolicyHistoryUpdateOne {
	iphuo.fields = append([]string{field}, fields...)
	return iphuo
}

// Save executes the query and returns the updated InternalPolicyHistory entity.
func (iphuo *InternalPolicyHistoryUpdateOne) Save(ctx context.Context) (*InternalPolicyHistory, error) {
	iphuo.defaults()
	return withHooks(ctx, iphuo.sqlSave, iphuo.mutation, iphuo.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (iphuo *InternalPolicyHistoryUpdateOne) SaveX(ctx context.Context) *InternalPolicyHistory {
	node, err := iphuo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (iphuo *InternalPolicyHistoryUpdateOne) Exec(ctx context.Context) error {
	_, err := iphuo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (iphuo *InternalPolicyHistoryUpdateOne) ExecX(ctx context.Context) {
	if err := iphuo.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (iphuo *InternalPolicyHistoryUpdateOne) defaults() {
	if _, ok := iphuo.mutation.UpdatedAt(); !ok && !iphuo.mutation.UpdatedAtCleared() {
		v := internalpolicyhistory.UpdateDefaultUpdatedAt()
		iphuo.mutation.SetUpdatedAt(v)
	}
}

// Modify adds a statement modifier for attaching custom logic to the UPDATE statement.
func (iphuo *InternalPolicyHistoryUpdateOne) Modify(modifiers ...func(u *sql.UpdateBuilder)) *InternalPolicyHistoryUpdateOne {
	iphuo.modifiers = append(iphuo.modifiers, modifiers...)
	return iphuo
}

func (iphuo *InternalPolicyHistoryUpdateOne) sqlSave(ctx context.Context) (_node *InternalPolicyHistory, err error) {
	_spec := sqlgraph.NewUpdateSpec(internalpolicyhistory.Table, internalpolicyhistory.Columns, sqlgraph.NewFieldSpec(internalpolicyhistory.FieldID, field.TypeString))
	id, ok := iphuo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`generated: missing "InternalPolicyHistory.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := iphuo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, internalpolicyhistory.FieldID)
		for _, f := range fields {
			if !internalpolicyhistory.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("generated: invalid field %q for query", f)}
			}
			if f != internalpolicyhistory.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := iphuo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if iphuo.mutation.RefCleared() {
		_spec.ClearField(internalpolicyhistory.FieldRef, field.TypeString)
	}
	if iphuo.mutation.CreatedAtCleared() {
		_spec.ClearField(internalpolicyhistory.FieldCreatedAt, field.TypeTime)
	}
	if value, ok := iphuo.mutation.UpdatedAt(); ok {
		_spec.SetField(internalpolicyhistory.FieldUpdatedAt, field.TypeTime, value)
	}
	if iphuo.mutation.UpdatedAtCleared() {
		_spec.ClearField(internalpolicyhistory.FieldUpdatedAt, field.TypeTime)
	}
	if iphuo.mutation.CreatedByCleared() {
		_spec.ClearField(internalpolicyhistory.FieldCreatedBy, field.TypeString)
	}
	if value, ok := iphuo.mutation.UpdatedBy(); ok {
		_spec.SetField(internalpolicyhistory.FieldUpdatedBy, field.TypeString, value)
	}
	if iphuo.mutation.UpdatedByCleared() {
		_spec.ClearField(internalpolicyhistory.FieldUpdatedBy, field.TypeString)
	}
	if value, ok := iphuo.mutation.DeletedAt(); ok {
		_spec.SetField(internalpolicyhistory.FieldDeletedAt, field.TypeTime, value)
	}
	if iphuo.mutation.DeletedAtCleared() {
		_spec.ClearField(internalpolicyhistory.FieldDeletedAt, field.TypeTime)
	}
	if value, ok := iphuo.mutation.DeletedBy(); ok {
		_spec.SetField(internalpolicyhistory.FieldDeletedBy, field.TypeString, value)
	}
	if iphuo.mutation.DeletedByCleared() {
		_spec.ClearField(internalpolicyhistory.FieldDeletedBy, field.TypeString)
	}
	if value, ok := iphuo.mutation.Tags(); ok {
		_spec.SetField(internalpolicyhistory.FieldTags, field.TypeJSON, value)
	}
	if value, ok := iphuo.mutation.AppendedTags(); ok {
		_spec.AddModifier(func(u *sql.UpdateBuilder) {
			sqljson.Append(u, internalpolicyhistory.FieldTags, value)
		})
	}
	if iphuo.mutation.TagsCleared() {
		_spec.ClearField(internalpolicyhistory.FieldTags, field.TypeJSON)
	}
	if value, ok := iphuo.mutation.Name(); ok {
		_spec.SetField(internalpolicyhistory.FieldName, field.TypeString, value)
	}
	if value, ok := iphuo.mutation.Description(); ok {
		_spec.SetField(internalpolicyhistory.FieldDescription, field.TypeString, value)
	}
	if value, ok := iphuo.mutation.Status(); ok {
		_spec.SetField(internalpolicyhistory.FieldStatus, field.TypeString, value)
	}
	if iphuo.mutation.StatusCleared() {
		_spec.ClearField(internalpolicyhistory.FieldStatus, field.TypeString)
	}
	if value, ok := iphuo.mutation.PolicyType(); ok {
		_spec.SetField(internalpolicyhistory.FieldPolicyType, field.TypeString, value)
	}
	if iphuo.mutation.PolicyTypeCleared() {
		_spec.ClearField(internalpolicyhistory.FieldPolicyType, field.TypeString)
	}
	if value, ok := iphuo.mutation.Version(); ok {
		_spec.SetField(internalpolicyhistory.FieldVersion, field.TypeString, value)
	}
	if iphuo.mutation.VersionCleared() {
		_spec.ClearField(internalpolicyhistory.FieldVersion, field.TypeString)
	}
	if value, ok := iphuo.mutation.PurposeAndScope(); ok {
		_spec.SetField(internalpolicyhistory.FieldPurposeAndScope, field.TypeString, value)
	}
	if iphuo.mutation.PurposeAndScopeCleared() {
		_spec.ClearField(internalpolicyhistory.FieldPurposeAndScope, field.TypeString)
	}
	if value, ok := iphuo.mutation.Background(); ok {
		_spec.SetField(internalpolicyhistory.FieldBackground, field.TypeString, value)
	}
	if iphuo.mutation.BackgroundCleared() {
		_spec.ClearField(internalpolicyhistory.FieldBackground, field.TypeString)
	}
	if value, ok := iphuo.mutation.Details(); ok {
		_spec.SetField(internalpolicyhistory.FieldDetails, field.TypeJSON, value)
	}
	if iphuo.mutation.DetailsCleared() {
		_spec.ClearField(internalpolicyhistory.FieldDetails, field.TypeJSON)
	}
	_spec.Node.Schema = iphuo.schemaConfig.InternalPolicyHistory
	ctx = internal.NewSchemaConfigContext(ctx, iphuo.schemaConfig)
	_spec.AddModifiers(iphuo.modifiers...)
	_node = &InternalPolicyHistory{config: iphuo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, iphuo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{internalpolicyhistory.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	iphuo.mutation.done = true
	return _node, nil
}
