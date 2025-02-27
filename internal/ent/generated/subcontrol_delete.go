// Code generated by ent, DO NOT EDIT.

package generated

import (
	"context"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/theopenlane/core/internal/ent/generated/predicate"

	"github.com/theopenlane/core/internal/ent/generated/internal"
	"github.com/theopenlane/core/internal/ent/generated/subcontrol"
)

// SubcontrolDelete is the builder for deleting a Subcontrol entity.
type SubcontrolDelete struct {
	config
	hooks    []Hook
	mutation *SubcontrolMutation
}

// Where appends a list predicates to the SubcontrolDelete builder.
func (sd *SubcontrolDelete) Where(ps ...predicate.Subcontrol) *SubcontrolDelete {
	sd.mutation.Where(ps...)
	return sd
}

// Exec executes the deletion query and returns how many vertices were deleted.
func (sd *SubcontrolDelete) Exec(ctx context.Context) (int, error) {
	return withHooks(ctx, sd.sqlExec, sd.mutation, sd.hooks)
}

// ExecX is like Exec, but panics if an error occurs.
func (sd *SubcontrolDelete) ExecX(ctx context.Context) int {
	n, err := sd.Exec(ctx)
	if err != nil {
		panic(err)
	}
	return n
}

func (sd *SubcontrolDelete) sqlExec(ctx context.Context) (int, error) {
	_spec := sqlgraph.NewDeleteSpec(subcontrol.Table, sqlgraph.NewFieldSpec(subcontrol.FieldID, field.TypeString))
	_spec.Node.Schema = sd.schemaConfig.Subcontrol
	ctx = internal.NewSchemaConfigContext(ctx, sd.schemaConfig)
	if ps := sd.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	affected, err := sqlgraph.DeleteNodes(ctx, sd.driver, _spec)
	if err != nil && sqlgraph.IsConstraintError(err) {
		err = &ConstraintError{msg: err.Error(), wrap: err}
	}
	sd.mutation.done = true
	return affected, err
}

// SubcontrolDeleteOne is the builder for deleting a single Subcontrol entity.
type SubcontrolDeleteOne struct {
	sd *SubcontrolDelete
}

// Where appends a list predicates to the SubcontrolDelete builder.
func (sdo *SubcontrolDeleteOne) Where(ps ...predicate.Subcontrol) *SubcontrolDeleteOne {
	sdo.sd.mutation.Where(ps...)
	return sdo
}

// Exec executes the deletion query.
func (sdo *SubcontrolDeleteOne) Exec(ctx context.Context) error {
	n, err := sdo.sd.Exec(ctx)
	switch {
	case err != nil:
		return err
	case n == 0:
		return &NotFoundError{subcontrol.Label}
	default:
		return nil
	}
}

// ExecX is like Exec, but panics if an error occurs.
func (sdo *SubcontrolDeleteOne) ExecX(ctx context.Context) {
	if err := sdo.Exec(ctx); err != nil {
		panic(err)
	}
}
