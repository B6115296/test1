// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/facebookincubator/ent/dialect/sql/sqlgraph"
	"github.com/facebookincubator/ent/schema/field"
	"github.com/sut63/team05/ent/bank"
	"github.com/sut63/team05/ent/member"
	"github.com/sut63/team05/ent/officer"
	"github.com/sut63/team05/ent/payback"
	"github.com/sut63/team05/ent/product"
)

// PaybackCreate is the builder for creating a Payback entity.
type PaybackCreate struct {
	config
	mutation *PaybackMutation
	hooks    []Hook
}

// SetAccountnumber sets the Accountnumber field.
func (pc *PaybackCreate) SetAccountnumber(s string) *PaybackCreate {
	pc.mutation.SetAccountnumber(s)
	return pc
}

// SetTransfertime sets the Transfertime field.
func (pc *PaybackCreate) SetTransfertime(t time.Time) *PaybackCreate {
	pc.mutation.SetTransfertime(t)
	return pc
}

// SetNillableTransfertime sets the Transfertime field if the given value is not nil.
func (pc *PaybackCreate) SetNillableTransfertime(t *time.Time) *PaybackCreate {
	if t != nil {
		pc.SetTransfertime(*t)
	}
	return pc
}

// SetOfficerID sets the Officer edge to Officer by id.
func (pc *PaybackCreate) SetOfficerID(id int) *PaybackCreate {
	pc.mutation.SetOfficerID(id)
	return pc
}

// SetNillableOfficerID sets the Officer edge to Officer by id if the given value is not nil.
func (pc *PaybackCreate) SetNillableOfficerID(id *int) *PaybackCreate {
	if id != nil {
		pc = pc.SetOfficerID(*id)
	}
	return pc
}

// SetOfficer sets the Officer edge to Officer.
func (pc *PaybackCreate) SetOfficer(o *Officer) *PaybackCreate {
	return pc.SetOfficerID(o.ID)
}

// SetMemberID sets the Member edge to Member by id.
func (pc *PaybackCreate) SetMemberID(id int) *PaybackCreate {
	pc.mutation.SetMemberID(id)
	return pc
}

// SetNillableMemberID sets the Member edge to Member by id if the given value is not nil.
func (pc *PaybackCreate) SetNillableMemberID(id *int) *PaybackCreate {
	if id != nil {
		pc = pc.SetMemberID(*id)
	}
	return pc
}

// SetMember sets the Member edge to Member.
func (pc *PaybackCreate) SetMember(m *Member) *PaybackCreate {
	return pc.SetMemberID(m.ID)
}

// SetProductID sets the Product edge to Product by id.
func (pc *PaybackCreate) SetProductID(id int) *PaybackCreate {
	pc.mutation.SetProductID(id)
	return pc
}

// SetNillableProductID sets the Product edge to Product by id if the given value is not nil.
func (pc *PaybackCreate) SetNillableProductID(id *int) *PaybackCreate {
	if id != nil {
		pc = pc.SetProductID(*id)
	}
	return pc
}

// SetProduct sets the Product edge to Product.
func (pc *PaybackCreate) SetProduct(p *Product) *PaybackCreate {
	return pc.SetProductID(p.ID)
}

// SetBankID sets the Bank edge to Bank by id.
func (pc *PaybackCreate) SetBankID(id int) *PaybackCreate {
	pc.mutation.SetBankID(id)
	return pc
}

// SetNillableBankID sets the Bank edge to Bank by id if the given value is not nil.
func (pc *PaybackCreate) SetNillableBankID(id *int) *PaybackCreate {
	if id != nil {
		pc = pc.SetBankID(*id)
	}
	return pc
}

// SetBank sets the Bank edge to Bank.
func (pc *PaybackCreate) SetBank(b *Bank) *PaybackCreate {
	return pc.SetBankID(b.ID)
}

// Mutation returns the PaybackMutation object of the builder.
func (pc *PaybackCreate) Mutation() *PaybackMutation {
	return pc.mutation
}

// Save creates the Payback in the database.
func (pc *PaybackCreate) Save(ctx context.Context) (*Payback, error) {
	if _, ok := pc.mutation.Accountnumber(); !ok {
		return nil, &ValidationError{Name: "Accountnumber", err: errors.New("ent: missing required field \"Accountnumber\"")}
	}
	if v, ok := pc.mutation.Accountnumber(); ok {
		if err := payback.AccountnumberValidator(v); err != nil {
			return nil, &ValidationError{Name: "Accountnumber", err: fmt.Errorf("ent: validator failed for field \"Accountnumber\": %w", err)}
		}
	}
	if _, ok := pc.mutation.Transfertime(); !ok {
		v := payback.DefaultTransfertime()
		pc.mutation.SetTransfertime(v)
	}
	var (
		err  error
		node *Payback
	)
	if len(pc.hooks) == 0 {
		node, err = pc.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*PaybackMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			pc.mutation = mutation
			node, err = pc.sqlSave(ctx)
			mutation.done = true
			return node, err
		})
		for i := len(pc.hooks) - 1; i >= 0; i-- {
			mut = pc.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, pc.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX calls Save and panics if Save returns an error.
func (pc *PaybackCreate) SaveX(ctx context.Context) *Payback {
	v, err := pc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

func (pc *PaybackCreate) sqlSave(ctx context.Context) (*Payback, error) {
	pa, _spec := pc.createSpec()
	if err := sqlgraph.CreateNode(ctx, pc.driver, _spec); err != nil {
		if cerr, ok := isSQLConstraintError(err); ok {
			err = cerr
		}
		return nil, err
	}
	id := _spec.ID.Value.(int64)
	pa.ID = int(id)
	return pa, nil
}

func (pc *PaybackCreate) createSpec() (*Payback, *sqlgraph.CreateSpec) {
	var (
		pa    = &Payback{config: pc.config}
		_spec = &sqlgraph.CreateSpec{
			Table: payback.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: payback.FieldID,
			},
		}
	)
	if value, ok := pc.mutation.Accountnumber(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: payback.FieldAccountnumber,
		})
		pa.Accountnumber = value
	}
	if value, ok := pc.mutation.Transfertime(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: payback.FieldTransfertime,
		})
		pa.Transfertime = value
	}
	if nodes := pc.mutation.OfficerIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   payback.OfficerTable,
			Columns: []string{payback.OfficerColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: officer.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := pc.mutation.MemberIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   payback.MemberTable,
			Columns: []string{payback.MemberColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: member.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := pc.mutation.ProductIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   payback.ProductTable,
			Columns: []string{payback.ProductColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: product.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := pc.mutation.BankIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   payback.BankTable,
			Columns: []string{payback.BankColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: bank.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	return pa, _spec
}
