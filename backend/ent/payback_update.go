// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"
	"time"

	"github.com/facebookincubator/ent/dialect/sql"
	"github.com/facebookincubator/ent/dialect/sql/sqlgraph"
	"github.com/facebookincubator/ent/schema/field"
	"github.com/sut63/team05/ent/bank"
	"github.com/sut63/team05/ent/member"
	"github.com/sut63/team05/ent/officer"
	"github.com/sut63/team05/ent/payback"
	"github.com/sut63/team05/ent/predicate"
	"github.com/sut63/team05/ent/product"
)

// PaybackUpdate is the builder for updating Payback entities.
type PaybackUpdate struct {
	config
	hooks      []Hook
	mutation   *PaybackMutation
	predicates []predicate.Payback
}

// Where adds a new predicate for the builder.
func (pu *PaybackUpdate) Where(ps ...predicate.Payback) *PaybackUpdate {
	pu.predicates = append(pu.predicates, ps...)
	return pu
}

// SetAccountnumber sets the Accountnumber field.
func (pu *PaybackUpdate) SetAccountnumber(s string) *PaybackUpdate {
	pu.mutation.SetAccountnumber(s)
	return pu
}

// SetTransfertime sets the Transfertime field.
func (pu *PaybackUpdate) SetTransfertime(t time.Time) *PaybackUpdate {
	pu.mutation.SetTransfertime(t)
	return pu
}

// SetNillableTransfertime sets the Transfertime field if the given value is not nil.
func (pu *PaybackUpdate) SetNillableTransfertime(t *time.Time) *PaybackUpdate {
	if t != nil {
		pu.SetTransfertime(*t)
	}
	return pu
}

// SetOfficerID sets the Officer edge to Officer by id.
func (pu *PaybackUpdate) SetOfficerID(id int) *PaybackUpdate {
	pu.mutation.SetOfficerID(id)
	return pu
}

// SetNillableOfficerID sets the Officer edge to Officer by id if the given value is not nil.
func (pu *PaybackUpdate) SetNillableOfficerID(id *int) *PaybackUpdate {
	if id != nil {
		pu = pu.SetOfficerID(*id)
	}
	return pu
}

// SetOfficer sets the Officer edge to Officer.
func (pu *PaybackUpdate) SetOfficer(o *Officer) *PaybackUpdate {
	return pu.SetOfficerID(o.ID)
}

// SetMemberID sets the Member edge to Member by id.
func (pu *PaybackUpdate) SetMemberID(id int) *PaybackUpdate {
	pu.mutation.SetMemberID(id)
	return pu
}

// SetNillableMemberID sets the Member edge to Member by id if the given value is not nil.
func (pu *PaybackUpdate) SetNillableMemberID(id *int) *PaybackUpdate {
	if id != nil {
		pu = pu.SetMemberID(*id)
	}
	return pu
}

// SetMember sets the Member edge to Member.
func (pu *PaybackUpdate) SetMember(m *Member) *PaybackUpdate {
	return pu.SetMemberID(m.ID)
}

// SetProductID sets the Product edge to Product by id.
func (pu *PaybackUpdate) SetProductID(id int) *PaybackUpdate {
	pu.mutation.SetProductID(id)
	return pu
}

// SetNillableProductID sets the Product edge to Product by id if the given value is not nil.
func (pu *PaybackUpdate) SetNillableProductID(id *int) *PaybackUpdate {
	if id != nil {
		pu = pu.SetProductID(*id)
	}
	return pu
}

// SetProduct sets the Product edge to Product.
func (pu *PaybackUpdate) SetProduct(p *Product) *PaybackUpdate {
	return pu.SetProductID(p.ID)
}

// SetBankID sets the Bank edge to Bank by id.
func (pu *PaybackUpdate) SetBankID(id int) *PaybackUpdate {
	pu.mutation.SetBankID(id)
	return pu
}

// SetNillableBankID sets the Bank edge to Bank by id if the given value is not nil.
func (pu *PaybackUpdate) SetNillableBankID(id *int) *PaybackUpdate {
	if id != nil {
		pu = pu.SetBankID(*id)
	}
	return pu
}

// SetBank sets the Bank edge to Bank.
func (pu *PaybackUpdate) SetBank(b *Bank) *PaybackUpdate {
	return pu.SetBankID(b.ID)
}

// Mutation returns the PaybackMutation object of the builder.
func (pu *PaybackUpdate) Mutation() *PaybackMutation {
	return pu.mutation
}

// ClearOfficer clears the Officer edge to Officer.
func (pu *PaybackUpdate) ClearOfficer() *PaybackUpdate {
	pu.mutation.ClearOfficer()
	return pu
}

// ClearMember clears the Member edge to Member.
func (pu *PaybackUpdate) ClearMember() *PaybackUpdate {
	pu.mutation.ClearMember()
	return pu
}

// ClearProduct clears the Product edge to Product.
func (pu *PaybackUpdate) ClearProduct() *PaybackUpdate {
	pu.mutation.ClearProduct()
	return pu
}

// ClearBank clears the Bank edge to Bank.
func (pu *PaybackUpdate) ClearBank() *PaybackUpdate {
	pu.mutation.ClearBank()
	return pu
}

// Save executes the query and returns the number of rows/vertices matched by this operation.
func (pu *PaybackUpdate) Save(ctx context.Context) (int, error) {
	if v, ok := pu.mutation.Accountnumber(); ok {
		if err := payback.AccountnumberValidator(v); err != nil {
			return 0, &ValidationError{Name: "Accountnumber", err: fmt.Errorf("ent: validator failed for field \"Accountnumber\": %w", err)}
		}
	}

	var (
		err      error
		affected int
	)
	if len(pu.hooks) == 0 {
		affected, err = pu.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*PaybackMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			pu.mutation = mutation
			affected, err = pu.sqlSave(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(pu.hooks) - 1; i >= 0; i-- {
			mut = pu.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, pu.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// SaveX is like Save, but panics if an error occurs.
func (pu *PaybackUpdate) SaveX(ctx context.Context) int {
	affected, err := pu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (pu *PaybackUpdate) Exec(ctx context.Context) error {
	_, err := pu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (pu *PaybackUpdate) ExecX(ctx context.Context) {
	if err := pu.Exec(ctx); err != nil {
		panic(err)
	}
}

func (pu *PaybackUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   payback.Table,
			Columns: payback.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: payback.FieldID,
			},
		},
	}
	if ps := pu.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := pu.mutation.Accountnumber(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: payback.FieldAccountnumber,
		})
	}
	if value, ok := pu.mutation.Transfertime(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: payback.FieldTransfertime,
		})
	}
	if pu.mutation.OfficerCleared() {
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
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := pu.mutation.OfficerIDs(); len(nodes) > 0 {
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
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if pu.mutation.MemberCleared() {
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
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := pu.mutation.MemberIDs(); len(nodes) > 0 {
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
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if pu.mutation.ProductCleared() {
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
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := pu.mutation.ProductIDs(); len(nodes) > 0 {
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
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if pu.mutation.BankCleared() {
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
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := pu.mutation.BankIDs(); len(nodes) > 0 {
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
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, pu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{payback.Label}
		} else if cerr, ok := isSQLConstraintError(err); ok {
			err = cerr
		}
		return 0, err
	}
	return n, nil
}

// PaybackUpdateOne is the builder for updating a single Payback entity.
type PaybackUpdateOne struct {
	config
	hooks    []Hook
	mutation *PaybackMutation
}

// SetAccountnumber sets the Accountnumber field.
func (puo *PaybackUpdateOne) SetAccountnumber(s string) *PaybackUpdateOne {
	puo.mutation.SetAccountnumber(s)
	return puo
}

// SetTransfertime sets the Transfertime field.
func (puo *PaybackUpdateOne) SetTransfertime(t time.Time) *PaybackUpdateOne {
	puo.mutation.SetTransfertime(t)
	return puo
}

// SetNillableTransfertime sets the Transfertime field if the given value is not nil.
func (puo *PaybackUpdateOne) SetNillableTransfertime(t *time.Time) *PaybackUpdateOne {
	if t != nil {
		puo.SetTransfertime(*t)
	}
	return puo
}

// SetOfficerID sets the Officer edge to Officer by id.
func (puo *PaybackUpdateOne) SetOfficerID(id int) *PaybackUpdateOne {
	puo.mutation.SetOfficerID(id)
	return puo
}

// SetNillableOfficerID sets the Officer edge to Officer by id if the given value is not nil.
func (puo *PaybackUpdateOne) SetNillableOfficerID(id *int) *PaybackUpdateOne {
	if id != nil {
		puo = puo.SetOfficerID(*id)
	}
	return puo
}

// SetOfficer sets the Officer edge to Officer.
func (puo *PaybackUpdateOne) SetOfficer(o *Officer) *PaybackUpdateOne {
	return puo.SetOfficerID(o.ID)
}

// SetMemberID sets the Member edge to Member by id.
func (puo *PaybackUpdateOne) SetMemberID(id int) *PaybackUpdateOne {
	puo.mutation.SetMemberID(id)
	return puo
}

// SetNillableMemberID sets the Member edge to Member by id if the given value is not nil.
func (puo *PaybackUpdateOne) SetNillableMemberID(id *int) *PaybackUpdateOne {
	if id != nil {
		puo = puo.SetMemberID(*id)
	}
	return puo
}

// SetMember sets the Member edge to Member.
func (puo *PaybackUpdateOne) SetMember(m *Member) *PaybackUpdateOne {
	return puo.SetMemberID(m.ID)
}

// SetProductID sets the Product edge to Product by id.
func (puo *PaybackUpdateOne) SetProductID(id int) *PaybackUpdateOne {
	puo.mutation.SetProductID(id)
	return puo
}

// SetNillableProductID sets the Product edge to Product by id if the given value is not nil.
func (puo *PaybackUpdateOne) SetNillableProductID(id *int) *PaybackUpdateOne {
	if id != nil {
		puo = puo.SetProductID(*id)
	}
	return puo
}

// SetProduct sets the Product edge to Product.
func (puo *PaybackUpdateOne) SetProduct(p *Product) *PaybackUpdateOne {
	return puo.SetProductID(p.ID)
}

// SetBankID sets the Bank edge to Bank by id.
func (puo *PaybackUpdateOne) SetBankID(id int) *PaybackUpdateOne {
	puo.mutation.SetBankID(id)
	return puo
}

// SetNillableBankID sets the Bank edge to Bank by id if the given value is not nil.
func (puo *PaybackUpdateOne) SetNillableBankID(id *int) *PaybackUpdateOne {
	if id != nil {
		puo = puo.SetBankID(*id)
	}
	return puo
}

// SetBank sets the Bank edge to Bank.
func (puo *PaybackUpdateOne) SetBank(b *Bank) *PaybackUpdateOne {
	return puo.SetBankID(b.ID)
}

// Mutation returns the PaybackMutation object of the builder.
func (puo *PaybackUpdateOne) Mutation() *PaybackMutation {
	return puo.mutation
}

// ClearOfficer clears the Officer edge to Officer.
func (puo *PaybackUpdateOne) ClearOfficer() *PaybackUpdateOne {
	puo.mutation.ClearOfficer()
	return puo
}

// ClearMember clears the Member edge to Member.
func (puo *PaybackUpdateOne) ClearMember() *PaybackUpdateOne {
	puo.mutation.ClearMember()
	return puo
}

// ClearProduct clears the Product edge to Product.
func (puo *PaybackUpdateOne) ClearProduct() *PaybackUpdateOne {
	puo.mutation.ClearProduct()
	return puo
}

// ClearBank clears the Bank edge to Bank.
func (puo *PaybackUpdateOne) ClearBank() *PaybackUpdateOne {
	puo.mutation.ClearBank()
	return puo
}

// Save executes the query and returns the updated entity.
func (puo *PaybackUpdateOne) Save(ctx context.Context) (*Payback, error) {
	if v, ok := puo.mutation.Accountnumber(); ok {
		if err := payback.AccountnumberValidator(v); err != nil {
			return nil, &ValidationError{Name: "Accountnumber", err: fmt.Errorf("ent: validator failed for field \"Accountnumber\": %w", err)}
		}
	}

	var (
		err  error
		node *Payback
	)
	if len(puo.hooks) == 0 {
		node, err = puo.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*PaybackMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			puo.mutation = mutation
			node, err = puo.sqlSave(ctx)
			mutation.done = true
			return node, err
		})
		for i := len(puo.hooks) - 1; i >= 0; i-- {
			mut = puo.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, puo.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX is like Save, but panics if an error occurs.
func (puo *PaybackUpdateOne) SaveX(ctx context.Context) *Payback {
	pa, err := puo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return pa
}

// Exec executes the query on the entity.
func (puo *PaybackUpdateOne) Exec(ctx context.Context) error {
	_, err := puo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (puo *PaybackUpdateOne) ExecX(ctx context.Context) {
	if err := puo.Exec(ctx); err != nil {
		panic(err)
	}
}

func (puo *PaybackUpdateOne) sqlSave(ctx context.Context) (pa *Payback, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   payback.Table,
			Columns: payback.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: payback.FieldID,
			},
		},
	}
	id, ok := puo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "ID", err: fmt.Errorf("missing Payback.ID for update")}
	}
	_spec.Node.ID.Value = id
	if value, ok := puo.mutation.Accountnumber(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: payback.FieldAccountnumber,
		})
	}
	if value, ok := puo.mutation.Transfertime(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: payback.FieldTransfertime,
		})
	}
	if puo.mutation.OfficerCleared() {
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
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := puo.mutation.OfficerIDs(); len(nodes) > 0 {
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
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if puo.mutation.MemberCleared() {
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
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := puo.mutation.MemberIDs(); len(nodes) > 0 {
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
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if puo.mutation.ProductCleared() {
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
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := puo.mutation.ProductIDs(); len(nodes) > 0 {
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
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if puo.mutation.BankCleared() {
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
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := puo.mutation.BankIDs(); len(nodes) > 0 {
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
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	pa = &Payback{config: puo.config}
	_spec.Assign = pa.assignValues
	_spec.ScanValues = pa.scanValues()
	if err = sqlgraph.UpdateNode(ctx, puo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{payback.Label}
		} else if cerr, ok := isSQLConstraintError(err); ok {
			err = cerr
		}
		return nil, err
	}
	return pa, nil
}
