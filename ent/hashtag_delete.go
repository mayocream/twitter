// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/mayocream/twitter/ent/hashtag"
	"github.com/mayocream/twitter/ent/predicate"
)

// HashtagDelete is the builder for deleting a Hashtag entity.
type HashtagDelete struct {
	config
	hooks    []Hook
	mutation *HashtagMutation
}

// Where appends a list predicates to the HashtagDelete builder.
func (hd *HashtagDelete) Where(ps ...predicate.Hashtag) *HashtagDelete {
	hd.mutation.Where(ps...)
	return hd
}

// Exec executes the deletion query and returns how many vertices were deleted.
func (hd *HashtagDelete) Exec(ctx context.Context) (int, error) {
	return withHooks(ctx, hd.sqlExec, hd.mutation, hd.hooks)
}

// ExecX is like Exec, but panics if an error occurs.
func (hd *HashtagDelete) ExecX(ctx context.Context) int {
	n, err := hd.Exec(ctx)
	if err != nil {
		panic(err)
	}
	return n
}

func (hd *HashtagDelete) sqlExec(ctx context.Context) (int, error) {
	_spec := sqlgraph.NewDeleteSpec(hashtag.Table, sqlgraph.NewFieldSpec(hashtag.FieldID, field.TypeInt))
	if ps := hd.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	affected, err := sqlgraph.DeleteNodes(ctx, hd.driver, _spec)
	if err != nil && sqlgraph.IsConstraintError(err) {
		err = &ConstraintError{msg: err.Error(), wrap: err}
	}
	hd.mutation.done = true
	return affected, err
}

// HashtagDeleteOne is the builder for deleting a single Hashtag entity.
type HashtagDeleteOne struct {
	hd *HashtagDelete
}

// Where appends a list predicates to the HashtagDelete builder.
func (hdo *HashtagDeleteOne) Where(ps ...predicate.Hashtag) *HashtagDeleteOne {
	hdo.hd.mutation.Where(ps...)
	return hdo
}

// Exec executes the deletion query.
func (hdo *HashtagDeleteOne) Exec(ctx context.Context) error {
	n, err := hdo.hd.Exec(ctx)
	switch {
	case err != nil:
		return err
	case n == 0:
		return &NotFoundError{hashtag.Label}
	default:
		return nil
	}
}

// ExecX is like Exec, but panics if an error occurs.
func (hdo *HashtagDeleteOne) ExecX(ctx context.Context) {
	if err := hdo.Exec(ctx); err != nil {
		panic(err)
	}
}
