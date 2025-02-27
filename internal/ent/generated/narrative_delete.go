// Code generated by ent, DO NOT EDIT.

package generated

import (
	"context"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/theopenlane/core/internal/ent/generated/predicate"

	"github.com/theopenlane/core/internal/ent/generated/internal"
	"github.com/theopenlane/core/internal/ent/generated/narrative"
)

// NarrativeDelete is the builder for deleting a Narrative entity.
type NarrativeDelete struct {
	config
	hooks    []Hook
	mutation *NarrativeMutation
}

// Where appends a list predicates to the NarrativeDelete builder.
func (nd *NarrativeDelete) Where(ps ...predicate.Narrative) *NarrativeDelete {
	nd.mutation.Where(ps...)
	return nd
}

// Exec executes the deletion query and returns how many vertices were deleted.
func (nd *NarrativeDelete) Exec(ctx context.Context) (int, error) {
	return withHooks(ctx, nd.sqlExec, nd.mutation, nd.hooks)
}

// ExecX is like Exec, but panics if an error occurs.
func (nd *NarrativeDelete) ExecX(ctx context.Context) int {
	n, err := nd.Exec(ctx)
	if err != nil {
		panic(err)
	}
	return n
}

func (nd *NarrativeDelete) sqlExec(ctx context.Context) (int, error) {
	_spec := sqlgraph.NewDeleteSpec(narrative.Table, sqlgraph.NewFieldSpec(narrative.FieldID, field.TypeString))
	_spec.Node.Schema = nd.schemaConfig.Narrative
	ctx = internal.NewSchemaConfigContext(ctx, nd.schemaConfig)
	if ps := nd.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	affected, err := sqlgraph.DeleteNodes(ctx, nd.driver, _spec)
	if err != nil && sqlgraph.IsConstraintError(err) {
		err = &ConstraintError{msg: err.Error(), wrap: err}
	}
	nd.mutation.done = true
	return affected, err
}

// NarrativeDeleteOne is the builder for deleting a single Narrative entity.
type NarrativeDeleteOne struct {
	nd *NarrativeDelete
}

// Where appends a list predicates to the NarrativeDelete builder.
func (ndo *NarrativeDeleteOne) Where(ps ...predicate.Narrative) *NarrativeDeleteOne {
	ndo.nd.mutation.Where(ps...)
	return ndo
}

// Exec executes the deletion query.
func (ndo *NarrativeDeleteOne) Exec(ctx context.Context) error {
	n, err := ndo.nd.Exec(ctx)
	switch {
	case err != nil:
		return err
	case n == 0:
		return &NotFoundError{narrative.Label}
	default:
		return nil
	}
}

// ExecX is like Exec, but panics if an error occurs.
func (ndo *NarrativeDeleteOne) ExecX(ctx context.Context) {
	if err := ndo.Exec(ctx); err != nil {
		panic(err)
	}
}
