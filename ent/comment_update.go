// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
	"github.com/k0kishima/golang-realworld-example-app/ent/comment"
	"github.com/k0kishima/golang-realworld-example-app/ent/predicate"
)

// CommentUpdate is the builder for updating Comment entities.
type CommentUpdate struct {
	config
	hooks    []Hook
	mutation *CommentMutation
}

// Where appends a list predicates to the CommentUpdate builder.
func (cu *CommentUpdate) Where(ps ...predicate.Comment) *CommentUpdate {
	cu.mutation.Where(ps...)
	return cu
}

// SetAuthorID sets the "author_id" field.
func (cu *CommentUpdate) SetAuthorID(u uuid.UUID) *CommentUpdate {
	cu.mutation.SetAuthorID(u)
	return cu
}

// SetNillableAuthorID sets the "author_id" field if the given value is not nil.
func (cu *CommentUpdate) SetNillableAuthorID(u *uuid.UUID) *CommentUpdate {
	if u != nil {
		cu.SetAuthorID(*u)
	}
	return cu
}

// SetArticleID sets the "article_id" field.
func (cu *CommentUpdate) SetArticleID(u uuid.UUID) *CommentUpdate {
	cu.mutation.SetArticleID(u)
	return cu
}

// SetNillableArticleID sets the "article_id" field if the given value is not nil.
func (cu *CommentUpdate) SetNillableArticleID(u *uuid.UUID) *CommentUpdate {
	if u != nil {
		cu.SetArticleID(*u)
	}
	return cu
}

// SetBody sets the "body" field.
func (cu *CommentUpdate) SetBody(s string) *CommentUpdate {
	cu.mutation.SetBody(s)
	return cu
}

// SetNillableBody sets the "body" field if the given value is not nil.
func (cu *CommentUpdate) SetNillableBody(s *string) *CommentUpdate {
	if s != nil {
		cu.SetBody(*s)
	}
	return cu
}

// SetUpdatedAt sets the "updated_at" field.
func (cu *CommentUpdate) SetUpdatedAt(t time.Time) *CommentUpdate {
	cu.mutation.SetUpdatedAt(t)
	return cu
}

// Mutation returns the CommentMutation object of the builder.
func (cu *CommentUpdate) Mutation() *CommentMutation {
	return cu.mutation
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (cu *CommentUpdate) Save(ctx context.Context) (int, error) {
	cu.defaults()
	return withHooks(ctx, cu.sqlSave, cu.mutation, cu.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (cu *CommentUpdate) SaveX(ctx context.Context) int {
	affected, err := cu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (cu *CommentUpdate) Exec(ctx context.Context) error {
	_, err := cu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (cu *CommentUpdate) ExecX(ctx context.Context) {
	if err := cu.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (cu *CommentUpdate) defaults() {
	if _, ok := cu.mutation.UpdatedAt(); !ok {
		v := comment.UpdateDefaultUpdatedAt()
		cu.mutation.SetUpdatedAt(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (cu *CommentUpdate) check() error {
	if v, ok := cu.mutation.Body(); ok {
		if err := comment.BodyValidator(v); err != nil {
			return &ValidationError{Name: "body", err: fmt.Errorf(`ent: validator failed for field "Comment.body": %w`, err)}
		}
	}
	return nil
}

func (cu *CommentUpdate) sqlSave(ctx context.Context) (n int, err error) {
	if err := cu.check(); err != nil {
		return n, err
	}
	_spec := sqlgraph.NewUpdateSpec(comment.Table, comment.Columns, sqlgraph.NewFieldSpec(comment.FieldID, field.TypeUUID))
	if ps := cu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := cu.mutation.AuthorID(); ok {
		_spec.SetField(comment.FieldAuthorID, field.TypeUUID, value)
	}
	if value, ok := cu.mutation.ArticleID(); ok {
		_spec.SetField(comment.FieldArticleID, field.TypeUUID, value)
	}
	if value, ok := cu.mutation.Body(); ok {
		_spec.SetField(comment.FieldBody, field.TypeString, value)
	}
	if value, ok := cu.mutation.UpdatedAt(); ok {
		_spec.SetField(comment.FieldUpdatedAt, field.TypeTime, value)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, cu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{comment.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	cu.mutation.done = true
	return n, nil
}

// CommentUpdateOne is the builder for updating a single Comment entity.
type CommentUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *CommentMutation
}

// SetAuthorID sets the "author_id" field.
func (cuo *CommentUpdateOne) SetAuthorID(u uuid.UUID) *CommentUpdateOne {
	cuo.mutation.SetAuthorID(u)
	return cuo
}

// SetNillableAuthorID sets the "author_id" field if the given value is not nil.
func (cuo *CommentUpdateOne) SetNillableAuthorID(u *uuid.UUID) *CommentUpdateOne {
	if u != nil {
		cuo.SetAuthorID(*u)
	}
	return cuo
}

// SetArticleID sets the "article_id" field.
func (cuo *CommentUpdateOne) SetArticleID(u uuid.UUID) *CommentUpdateOne {
	cuo.mutation.SetArticleID(u)
	return cuo
}

// SetNillableArticleID sets the "article_id" field if the given value is not nil.
func (cuo *CommentUpdateOne) SetNillableArticleID(u *uuid.UUID) *CommentUpdateOne {
	if u != nil {
		cuo.SetArticleID(*u)
	}
	return cuo
}

// SetBody sets the "body" field.
func (cuo *CommentUpdateOne) SetBody(s string) *CommentUpdateOne {
	cuo.mutation.SetBody(s)
	return cuo
}

// SetNillableBody sets the "body" field if the given value is not nil.
func (cuo *CommentUpdateOne) SetNillableBody(s *string) *CommentUpdateOne {
	if s != nil {
		cuo.SetBody(*s)
	}
	return cuo
}

// SetUpdatedAt sets the "updated_at" field.
func (cuo *CommentUpdateOne) SetUpdatedAt(t time.Time) *CommentUpdateOne {
	cuo.mutation.SetUpdatedAt(t)
	return cuo
}

// Mutation returns the CommentMutation object of the builder.
func (cuo *CommentUpdateOne) Mutation() *CommentMutation {
	return cuo.mutation
}

// Where appends a list predicates to the CommentUpdate builder.
func (cuo *CommentUpdateOne) Where(ps ...predicate.Comment) *CommentUpdateOne {
	cuo.mutation.Where(ps...)
	return cuo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (cuo *CommentUpdateOne) Select(field string, fields ...string) *CommentUpdateOne {
	cuo.fields = append([]string{field}, fields...)
	return cuo
}

// Save executes the query and returns the updated Comment entity.
func (cuo *CommentUpdateOne) Save(ctx context.Context) (*Comment, error) {
	cuo.defaults()
	return withHooks(ctx, cuo.sqlSave, cuo.mutation, cuo.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (cuo *CommentUpdateOne) SaveX(ctx context.Context) *Comment {
	node, err := cuo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (cuo *CommentUpdateOne) Exec(ctx context.Context) error {
	_, err := cuo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (cuo *CommentUpdateOne) ExecX(ctx context.Context) {
	if err := cuo.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (cuo *CommentUpdateOne) defaults() {
	if _, ok := cuo.mutation.UpdatedAt(); !ok {
		v := comment.UpdateDefaultUpdatedAt()
		cuo.mutation.SetUpdatedAt(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (cuo *CommentUpdateOne) check() error {
	if v, ok := cuo.mutation.Body(); ok {
		if err := comment.BodyValidator(v); err != nil {
			return &ValidationError{Name: "body", err: fmt.Errorf(`ent: validator failed for field "Comment.body": %w`, err)}
		}
	}
	return nil
}

func (cuo *CommentUpdateOne) sqlSave(ctx context.Context) (_node *Comment, err error) {
	if err := cuo.check(); err != nil {
		return _node, err
	}
	_spec := sqlgraph.NewUpdateSpec(comment.Table, comment.Columns, sqlgraph.NewFieldSpec(comment.FieldID, field.TypeUUID))
	id, ok := cuo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "Comment.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := cuo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, comment.FieldID)
		for _, f := range fields {
			if !comment.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != comment.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := cuo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := cuo.mutation.AuthorID(); ok {
		_spec.SetField(comment.FieldAuthorID, field.TypeUUID, value)
	}
	if value, ok := cuo.mutation.ArticleID(); ok {
		_spec.SetField(comment.FieldArticleID, field.TypeUUID, value)
	}
	if value, ok := cuo.mutation.Body(); ok {
		_spec.SetField(comment.FieldBody, field.TypeString, value)
	}
	if value, ok := cuo.mutation.UpdatedAt(); ok {
		_spec.SetField(comment.FieldUpdatedAt, field.TypeTime, value)
	}
	_node = &Comment{config: cuo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, cuo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{comment.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	cuo.mutation.done = true
	return _node, nil
}
