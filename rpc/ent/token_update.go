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
	uuid "github.com/gofrs/uuid/v5"
	"github.com/toutmost/admin-core/rpc/ent/predicate"
	"github.com/toutmost/admin-core/rpc/ent/token"
)

// TokenUpdate is the builder for updating Token entities.
type TokenUpdate struct {
	config
	hooks    []Hook
	mutation *TokenMutation
}

// Where appends a list predicates to the TokenUpdate builder.
func (tu *TokenUpdate) Where(ps ...predicate.Token) *TokenUpdate {
	tu.mutation.Where(ps...)
	return tu
}

// SetUpdatedAt sets the "updated_at" field.
func (tu *TokenUpdate) SetUpdatedAt(t time.Time) *TokenUpdate {
	tu.mutation.SetUpdatedAt(t)
	return tu
}

// SetStatus sets the "status" field.
func (tu *TokenUpdate) SetStatus(u uint8) *TokenUpdate {
	tu.mutation.ResetStatus()
	tu.mutation.SetStatus(u)
	return tu
}

// SetNillableStatus sets the "status" field if the given value is not nil.
func (tu *TokenUpdate) SetNillableStatus(u *uint8) *TokenUpdate {
	if u != nil {
		tu.SetStatus(*u)
	}
	return tu
}

// AddStatus adds u to the "status" field.
func (tu *TokenUpdate) AddStatus(u int8) *TokenUpdate {
	tu.mutation.AddStatus(u)
	return tu
}

// ClearStatus clears the value of the "status" field.
func (tu *TokenUpdate) ClearStatus() *TokenUpdate {
	tu.mutation.ClearStatus()
	return tu
}

// SetUUID sets the "uuid" field.
func (tu *TokenUpdate) SetUUID(u uuid.UUID) *TokenUpdate {
	tu.mutation.SetUUID(u)
	return tu
}

// SetNillableUUID sets the "uuid" field if the given value is not nil.
func (tu *TokenUpdate) SetNillableUUID(u *uuid.UUID) *TokenUpdate {
	if u != nil {
		tu.SetUUID(*u)
	}
	return tu
}

// SetUsername sets the "username" field.
func (tu *TokenUpdate) SetUsername(s string) *TokenUpdate {
	tu.mutation.SetUsername(s)
	return tu
}

// SetNillableUsername sets the "username" field if the given value is not nil.
func (tu *TokenUpdate) SetNillableUsername(s *string) *TokenUpdate {
	if s != nil {
		tu.SetUsername(*s)
	}
	return tu
}

// SetToken sets the "token" field.
func (tu *TokenUpdate) SetToken(s string) *TokenUpdate {
	tu.mutation.SetToken(s)
	return tu
}

// SetNillableToken sets the "token" field if the given value is not nil.
func (tu *TokenUpdate) SetNillableToken(s *string) *TokenUpdate {
	if s != nil {
		tu.SetToken(*s)
	}
	return tu
}

// SetSource sets the "source" field.
func (tu *TokenUpdate) SetSource(s string) *TokenUpdate {
	tu.mutation.SetSource(s)
	return tu
}

// SetNillableSource sets the "source" field if the given value is not nil.
func (tu *TokenUpdate) SetNillableSource(s *string) *TokenUpdate {
	if s != nil {
		tu.SetSource(*s)
	}
	return tu
}

// SetExpiredAt sets the "expired_at" field.
func (tu *TokenUpdate) SetExpiredAt(t time.Time) *TokenUpdate {
	tu.mutation.SetExpiredAt(t)
	return tu
}

// SetNillableExpiredAt sets the "expired_at" field if the given value is not nil.
func (tu *TokenUpdate) SetNillableExpiredAt(t *time.Time) *TokenUpdate {
	if t != nil {
		tu.SetExpiredAt(*t)
	}
	return tu
}

// Mutation returns the TokenMutation object of the builder.
func (tu *TokenUpdate) Mutation() *TokenMutation {
	return tu.mutation
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (tu *TokenUpdate) Save(ctx context.Context) (int, error) {
	tu.defaults()
	return withHooks(ctx, tu.sqlSave, tu.mutation, tu.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (tu *TokenUpdate) SaveX(ctx context.Context) int {
	affected, err := tu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (tu *TokenUpdate) Exec(ctx context.Context) error {
	_, err := tu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (tu *TokenUpdate) ExecX(ctx context.Context) {
	if err := tu.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (tu *TokenUpdate) defaults() {
	if _, ok := tu.mutation.UpdatedAt(); !ok {
		v := token.UpdateDefaultUpdatedAt()
		tu.mutation.SetUpdatedAt(v)
	}
}

func (tu *TokenUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := sqlgraph.NewUpdateSpec(token.Table, token.Columns, sqlgraph.NewFieldSpec(token.FieldID, field.TypeUUID))
	if ps := tu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := tu.mutation.UpdatedAt(); ok {
		_spec.SetField(token.FieldUpdatedAt, field.TypeTime, value)
	}
	if value, ok := tu.mutation.Status(); ok {
		_spec.SetField(token.FieldStatus, field.TypeUint8, value)
	}
	if value, ok := tu.mutation.AddedStatus(); ok {
		_spec.AddField(token.FieldStatus, field.TypeUint8, value)
	}
	if tu.mutation.StatusCleared() {
		_spec.ClearField(token.FieldStatus, field.TypeUint8)
	}
	if value, ok := tu.mutation.UUID(); ok {
		_spec.SetField(token.FieldUUID, field.TypeUUID, value)
	}
	if value, ok := tu.mutation.Username(); ok {
		_spec.SetField(token.FieldUsername, field.TypeString, value)
	}
	if value, ok := tu.mutation.Token(); ok {
		_spec.SetField(token.FieldToken, field.TypeString, value)
	}
	if value, ok := tu.mutation.Source(); ok {
		_spec.SetField(token.FieldSource, field.TypeString, value)
	}
	if value, ok := tu.mutation.ExpiredAt(); ok {
		_spec.SetField(token.FieldExpiredAt, field.TypeTime, value)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, tu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{token.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	tu.mutation.done = true
	return n, nil
}

// TokenUpdateOne is the builder for updating a single Token entity.
type TokenUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *TokenMutation
}

// SetUpdatedAt sets the "updated_at" field.
func (tuo *TokenUpdateOne) SetUpdatedAt(t time.Time) *TokenUpdateOne {
	tuo.mutation.SetUpdatedAt(t)
	return tuo
}

// SetStatus sets the "status" field.
func (tuo *TokenUpdateOne) SetStatus(u uint8) *TokenUpdateOne {
	tuo.mutation.ResetStatus()
	tuo.mutation.SetStatus(u)
	return tuo
}

// SetNillableStatus sets the "status" field if the given value is not nil.
func (tuo *TokenUpdateOne) SetNillableStatus(u *uint8) *TokenUpdateOne {
	if u != nil {
		tuo.SetStatus(*u)
	}
	return tuo
}

// AddStatus adds u to the "status" field.
func (tuo *TokenUpdateOne) AddStatus(u int8) *TokenUpdateOne {
	tuo.mutation.AddStatus(u)
	return tuo
}

// ClearStatus clears the value of the "status" field.
func (tuo *TokenUpdateOne) ClearStatus() *TokenUpdateOne {
	tuo.mutation.ClearStatus()
	return tuo
}

// SetUUID sets the "uuid" field.
func (tuo *TokenUpdateOne) SetUUID(u uuid.UUID) *TokenUpdateOne {
	tuo.mutation.SetUUID(u)
	return tuo
}

// SetNillableUUID sets the "uuid" field if the given value is not nil.
func (tuo *TokenUpdateOne) SetNillableUUID(u *uuid.UUID) *TokenUpdateOne {
	if u != nil {
		tuo.SetUUID(*u)
	}
	return tuo
}

// SetUsername sets the "username" field.
func (tuo *TokenUpdateOne) SetUsername(s string) *TokenUpdateOne {
	tuo.mutation.SetUsername(s)
	return tuo
}

// SetNillableUsername sets the "username" field if the given value is not nil.
func (tuo *TokenUpdateOne) SetNillableUsername(s *string) *TokenUpdateOne {
	if s != nil {
		tuo.SetUsername(*s)
	}
	return tuo
}

// SetToken sets the "token" field.
func (tuo *TokenUpdateOne) SetToken(s string) *TokenUpdateOne {
	tuo.mutation.SetToken(s)
	return tuo
}

// SetNillableToken sets the "token" field if the given value is not nil.
func (tuo *TokenUpdateOne) SetNillableToken(s *string) *TokenUpdateOne {
	if s != nil {
		tuo.SetToken(*s)
	}
	return tuo
}

// SetSource sets the "source" field.
func (tuo *TokenUpdateOne) SetSource(s string) *TokenUpdateOne {
	tuo.mutation.SetSource(s)
	return tuo
}

// SetNillableSource sets the "source" field if the given value is not nil.
func (tuo *TokenUpdateOne) SetNillableSource(s *string) *TokenUpdateOne {
	if s != nil {
		tuo.SetSource(*s)
	}
	return tuo
}

// SetExpiredAt sets the "expired_at" field.
func (tuo *TokenUpdateOne) SetExpiredAt(t time.Time) *TokenUpdateOne {
	tuo.mutation.SetExpiredAt(t)
	return tuo
}

// SetNillableExpiredAt sets the "expired_at" field if the given value is not nil.
func (tuo *TokenUpdateOne) SetNillableExpiredAt(t *time.Time) *TokenUpdateOne {
	if t != nil {
		tuo.SetExpiredAt(*t)
	}
	return tuo
}

// Mutation returns the TokenMutation object of the builder.
func (tuo *TokenUpdateOne) Mutation() *TokenMutation {
	return tuo.mutation
}

// Where appends a list predicates to the TokenUpdate builder.
func (tuo *TokenUpdateOne) Where(ps ...predicate.Token) *TokenUpdateOne {
	tuo.mutation.Where(ps...)
	return tuo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (tuo *TokenUpdateOne) Select(field string, fields ...string) *TokenUpdateOne {
	tuo.fields = append([]string{field}, fields...)
	return tuo
}

// Save executes the query and returns the updated Token entity.
func (tuo *TokenUpdateOne) Save(ctx context.Context) (*Token, error) {
	tuo.defaults()
	return withHooks(ctx, tuo.sqlSave, tuo.mutation, tuo.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (tuo *TokenUpdateOne) SaveX(ctx context.Context) *Token {
	node, err := tuo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (tuo *TokenUpdateOne) Exec(ctx context.Context) error {
	_, err := tuo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (tuo *TokenUpdateOne) ExecX(ctx context.Context) {
	if err := tuo.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (tuo *TokenUpdateOne) defaults() {
	if _, ok := tuo.mutation.UpdatedAt(); !ok {
		v := token.UpdateDefaultUpdatedAt()
		tuo.mutation.SetUpdatedAt(v)
	}
}

func (tuo *TokenUpdateOne) sqlSave(ctx context.Context) (_node *Token, err error) {
	_spec := sqlgraph.NewUpdateSpec(token.Table, token.Columns, sqlgraph.NewFieldSpec(token.FieldID, field.TypeUUID))
	id, ok := tuo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "Token.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := tuo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, token.FieldID)
		for _, f := range fields {
			if !token.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != token.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := tuo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := tuo.mutation.UpdatedAt(); ok {
		_spec.SetField(token.FieldUpdatedAt, field.TypeTime, value)
	}
	if value, ok := tuo.mutation.Status(); ok {
		_spec.SetField(token.FieldStatus, field.TypeUint8, value)
	}
	if value, ok := tuo.mutation.AddedStatus(); ok {
		_spec.AddField(token.FieldStatus, field.TypeUint8, value)
	}
	if tuo.mutation.StatusCleared() {
		_spec.ClearField(token.FieldStatus, field.TypeUint8)
	}
	if value, ok := tuo.mutation.UUID(); ok {
		_spec.SetField(token.FieldUUID, field.TypeUUID, value)
	}
	if value, ok := tuo.mutation.Username(); ok {
		_spec.SetField(token.FieldUsername, field.TypeString, value)
	}
	if value, ok := tuo.mutation.Token(); ok {
		_spec.SetField(token.FieldToken, field.TypeString, value)
	}
	if value, ok := tuo.mutation.Source(); ok {
		_spec.SetField(token.FieldSource, field.TypeString, value)
	}
	if value, ok := tuo.mutation.ExpiredAt(); ok {
		_spec.SetField(token.FieldExpiredAt, field.TypeTime, value)
	}
	_node = &Token{config: tuo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, tuo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{token.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	tuo.mutation.done = true
	return _node, nil
}
