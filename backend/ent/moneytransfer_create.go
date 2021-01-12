// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"

	"github.com/facebookincubator/ent/dialect/sql/sqlgraph"
	"github.com/facebookincubator/ent/schema/field"
	"github.com/sut63/team05/ent/moneytransfer"
	"github.com/sut63/team05/ent/payment"
)

// MoneyTransferCreate is the builder for creating a MoneyTransfer entity.
type MoneyTransferCreate struct {
	config
	mutation *MoneyTransferMutation
	hooks    []Hook
}

// SetMoneytransferType sets the moneytransfer_type field.
func (mtc *MoneyTransferCreate) SetMoneytransferType(s string) *MoneyTransferCreate {
	mtc.mutation.SetMoneytransferType(s)
	return mtc
}

// AddMoneytransferPaymentIDs adds the moneytransfer_payment edge to Payment by ids.
func (mtc *MoneyTransferCreate) AddMoneytransferPaymentIDs(ids ...int) *MoneyTransferCreate {
	mtc.mutation.AddMoneytransferPaymentIDs(ids...)
	return mtc
}

// AddMoneytransferPayment adds the moneytransfer_payment edges to Payment.
func (mtc *MoneyTransferCreate) AddMoneytransferPayment(p ...*Payment) *MoneyTransferCreate {
	ids := make([]int, len(p))
	for i := range p {
		ids[i] = p[i].ID
	}
	return mtc.AddMoneytransferPaymentIDs(ids...)
}

// Mutation returns the MoneyTransferMutation object of the builder.
func (mtc *MoneyTransferCreate) Mutation() *MoneyTransferMutation {
	return mtc.mutation
}

// Save creates the MoneyTransfer in the database.
func (mtc *MoneyTransferCreate) Save(ctx context.Context) (*MoneyTransfer, error) {
	if _, ok := mtc.mutation.MoneytransferType(); !ok {
		return nil, &ValidationError{Name: "moneytransfer_type", err: errors.New("ent: missing required field \"moneytransfer_type\"")}
	}
	if v, ok := mtc.mutation.MoneytransferType(); ok {
		if err := moneytransfer.MoneytransferTypeValidator(v); err != nil {
			return nil, &ValidationError{Name: "moneytransfer_type", err: fmt.Errorf("ent: validator failed for field \"moneytransfer_type\": %w", err)}
		}
	}
	var (
		err  error
		node *MoneyTransfer
	)
	if len(mtc.hooks) == 0 {
		node, err = mtc.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*MoneyTransferMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			mtc.mutation = mutation
			node, err = mtc.sqlSave(ctx)
			mutation.done = true
			return node, err
		})
		for i := len(mtc.hooks) - 1; i >= 0; i-- {
			mut = mtc.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, mtc.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX calls Save and panics if Save returns an error.
func (mtc *MoneyTransferCreate) SaveX(ctx context.Context) *MoneyTransfer {
	v, err := mtc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

func (mtc *MoneyTransferCreate) sqlSave(ctx context.Context) (*MoneyTransfer, error) {
	mt, _spec := mtc.createSpec()
	if err := sqlgraph.CreateNode(ctx, mtc.driver, _spec); err != nil {
		if cerr, ok := isSQLConstraintError(err); ok {
			err = cerr
		}
		return nil, err
	}
	id := _spec.ID.Value.(int64)
	mt.ID = int(id)
	return mt, nil
}

func (mtc *MoneyTransferCreate) createSpec() (*MoneyTransfer, *sqlgraph.CreateSpec) {
	var (
		mt    = &MoneyTransfer{config: mtc.config}
		_spec = &sqlgraph.CreateSpec{
			Table: moneytransfer.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: moneytransfer.FieldID,
			},
		}
	)
	if value, ok := mtc.mutation.MoneytransferType(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: moneytransfer.FieldMoneytransferType,
		})
		mt.MoneytransferType = value
	}
	if nodes := mtc.mutation.MoneytransferPaymentIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   moneytransfer.MoneytransferPaymentTable,
			Columns: []string{moneytransfer.MoneytransferPaymentColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: payment.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	return mt, _spec
}