// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"
	"time"

	"github.com/facebookincubator/ent/dialect/sql"
	"github.com/facebookincubator/ent/dialect/sql/sqlgraph"
	"github.com/facebookincubator/ent/schema/field"
	"github.com/sut63/team05/ent/amountpaid"
	"github.com/sut63/team05/ent/hospital"
	"github.com/sut63/team05/ent/member"
	"github.com/sut63/team05/ent/officer"
	"github.com/sut63/team05/ent/predicate"
	"github.com/sut63/team05/ent/product"
	"github.com/sut63/team05/ent/recordinsurance"
)

// RecordinsuranceUpdate is the builder for updating Recordinsurance entities.
type RecordinsuranceUpdate struct {
	config
	hooks      []Hook
	mutation   *RecordinsuranceMutation
	predicates []predicate.Recordinsurance
}

// Where adds a new predicate for the builder.
func (ru *RecordinsuranceUpdate) Where(ps ...predicate.Recordinsurance) *RecordinsuranceUpdate {
	ru.predicates = append(ru.predicates, ps...)
	return ru
}

// SetRecordinsuranceTime sets the recordinsurance_time field.
func (ru *RecordinsuranceUpdate) SetRecordinsuranceTime(t time.Time) *RecordinsuranceUpdate {
	ru.mutation.SetRecordinsuranceTime(t)
	return ru
}

// SetNillableRecordinsuranceTime sets the recordinsurance_time field if the given value is not nil.
func (ru *RecordinsuranceUpdate) SetNillableRecordinsuranceTime(t *time.Time) *RecordinsuranceUpdate {
	if t != nil {
		ru.SetRecordinsuranceTime(*t)
	}
	return ru
}

// SetMemberID sets the Member edge to Member by id.
func (ru *RecordinsuranceUpdate) SetMemberID(id int) *RecordinsuranceUpdate {
	ru.mutation.SetMemberID(id)
	return ru
}

// SetNillableMemberID sets the Member edge to Member by id if the given value is not nil.
func (ru *RecordinsuranceUpdate) SetNillableMemberID(id *int) *RecordinsuranceUpdate {
	if id != nil {
		ru = ru.SetMemberID(*id)
	}
	return ru
}

// SetMember sets the Member edge to Member.
func (ru *RecordinsuranceUpdate) SetMember(m *Member) *RecordinsuranceUpdate {
	return ru.SetMemberID(m.ID)
}

// SetHospitalID sets the Hospital edge to Hospital by id.
func (ru *RecordinsuranceUpdate) SetHospitalID(id int) *RecordinsuranceUpdate {
	ru.mutation.SetHospitalID(id)
	return ru
}

// SetNillableHospitalID sets the Hospital edge to Hospital by id if the given value is not nil.
func (ru *RecordinsuranceUpdate) SetNillableHospitalID(id *int) *RecordinsuranceUpdate {
	if id != nil {
		ru = ru.SetHospitalID(*id)
	}
	return ru
}

// SetHospital sets the Hospital edge to Hospital.
func (ru *RecordinsuranceUpdate) SetHospital(h *Hospital) *RecordinsuranceUpdate {
	return ru.SetHospitalID(h.ID)
}

// SetOfficerID sets the Officer edge to Officer by id.
func (ru *RecordinsuranceUpdate) SetOfficerID(id int) *RecordinsuranceUpdate {
	ru.mutation.SetOfficerID(id)
	return ru
}

// SetNillableOfficerID sets the Officer edge to Officer by id if the given value is not nil.
func (ru *RecordinsuranceUpdate) SetNillableOfficerID(id *int) *RecordinsuranceUpdate {
	if id != nil {
		ru = ru.SetOfficerID(*id)
	}
	return ru
}

// SetOfficer sets the Officer edge to Officer.
func (ru *RecordinsuranceUpdate) SetOfficer(o *Officer) *RecordinsuranceUpdate {
	return ru.SetOfficerID(o.ID)
}

// SetProductID sets the Product edge to Product by id.
func (ru *RecordinsuranceUpdate) SetProductID(id int) *RecordinsuranceUpdate {
	ru.mutation.SetProductID(id)
	return ru
}

// SetNillableProductID sets the Product edge to Product by id if the given value is not nil.
func (ru *RecordinsuranceUpdate) SetNillableProductID(id *int) *RecordinsuranceUpdate {
	if id != nil {
		ru = ru.SetProductID(*id)
	}
	return ru
}

// SetProduct sets the Product edge to Product.
func (ru *RecordinsuranceUpdate) SetProduct(p *Product) *RecordinsuranceUpdate {
	return ru.SetProductID(p.ID)
}

// SetAmountpaidID sets the Amountpaid edge to Amountpaid by id.
func (ru *RecordinsuranceUpdate) SetAmountpaidID(id int) *RecordinsuranceUpdate {
	ru.mutation.SetAmountpaidID(id)
	return ru
}

// SetNillableAmountpaidID sets the Amountpaid edge to Amountpaid by id if the given value is not nil.
func (ru *RecordinsuranceUpdate) SetNillableAmountpaidID(id *int) *RecordinsuranceUpdate {
	if id != nil {
		ru = ru.SetAmountpaidID(*id)
	}
	return ru
}

// SetAmountpaid sets the Amountpaid edge to Amountpaid.
func (ru *RecordinsuranceUpdate) SetAmountpaid(a *Amountpaid) *RecordinsuranceUpdate {
	return ru.SetAmountpaidID(a.ID)
}

// Mutation returns the RecordinsuranceMutation object of the builder.
func (ru *RecordinsuranceUpdate) Mutation() *RecordinsuranceMutation {
	return ru.mutation
}

// ClearMember clears the Member edge to Member.
func (ru *RecordinsuranceUpdate) ClearMember() *RecordinsuranceUpdate {
	ru.mutation.ClearMember()
	return ru
}

// ClearHospital clears the Hospital edge to Hospital.
func (ru *RecordinsuranceUpdate) ClearHospital() *RecordinsuranceUpdate {
	ru.mutation.ClearHospital()
	return ru
}

// ClearOfficer clears the Officer edge to Officer.
func (ru *RecordinsuranceUpdate) ClearOfficer() *RecordinsuranceUpdate {
	ru.mutation.ClearOfficer()
	return ru
}

// ClearProduct clears the Product edge to Product.
func (ru *RecordinsuranceUpdate) ClearProduct() *RecordinsuranceUpdate {
	ru.mutation.ClearProduct()
	return ru
}

// ClearAmountpaid clears the Amountpaid edge to Amountpaid.
func (ru *RecordinsuranceUpdate) ClearAmountpaid() *RecordinsuranceUpdate {
	ru.mutation.ClearAmountpaid()
	return ru
}

// Save executes the query and returns the number of rows/vertices matched by this operation.
func (ru *RecordinsuranceUpdate) Save(ctx context.Context) (int, error) {

	var (
		err      error
		affected int
	)
	if len(ru.hooks) == 0 {
		affected, err = ru.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*RecordinsuranceMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			ru.mutation = mutation
			affected, err = ru.sqlSave(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(ru.hooks) - 1; i >= 0; i-- {
			mut = ru.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, ru.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// SaveX is like Save, but panics if an error occurs.
func (ru *RecordinsuranceUpdate) SaveX(ctx context.Context) int {
	affected, err := ru.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (ru *RecordinsuranceUpdate) Exec(ctx context.Context) error {
	_, err := ru.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (ru *RecordinsuranceUpdate) ExecX(ctx context.Context) {
	if err := ru.Exec(ctx); err != nil {
		panic(err)
	}
}

func (ru *RecordinsuranceUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   recordinsurance.Table,
			Columns: recordinsurance.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: recordinsurance.FieldID,
			},
		},
	}
	if ps := ru.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := ru.mutation.RecordinsuranceTime(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: recordinsurance.FieldRecordinsuranceTime,
		})
	}
	if ru.mutation.MemberCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   recordinsurance.MemberTable,
			Columns: []string{recordinsurance.MemberColumn},
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
	if nodes := ru.mutation.MemberIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   recordinsurance.MemberTable,
			Columns: []string{recordinsurance.MemberColumn},
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
	if ru.mutation.HospitalCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   recordinsurance.HospitalTable,
			Columns: []string{recordinsurance.HospitalColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: hospital.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := ru.mutation.HospitalIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   recordinsurance.HospitalTable,
			Columns: []string{recordinsurance.HospitalColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: hospital.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if ru.mutation.OfficerCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   recordinsurance.OfficerTable,
			Columns: []string{recordinsurance.OfficerColumn},
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
	if nodes := ru.mutation.OfficerIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   recordinsurance.OfficerTable,
			Columns: []string{recordinsurance.OfficerColumn},
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
	if ru.mutation.ProductCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   recordinsurance.ProductTable,
			Columns: []string{recordinsurance.ProductColumn},
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
	if nodes := ru.mutation.ProductIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   recordinsurance.ProductTable,
			Columns: []string{recordinsurance.ProductColumn},
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
	if ru.mutation.AmountpaidCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   recordinsurance.AmountpaidTable,
			Columns: []string{recordinsurance.AmountpaidColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: amountpaid.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := ru.mutation.AmountpaidIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   recordinsurance.AmountpaidTable,
			Columns: []string{recordinsurance.AmountpaidColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: amountpaid.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, ru.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{recordinsurance.Label}
		} else if cerr, ok := isSQLConstraintError(err); ok {
			err = cerr
		}
		return 0, err
	}
	return n, nil
}

// RecordinsuranceUpdateOne is the builder for updating a single Recordinsurance entity.
type RecordinsuranceUpdateOne struct {
	config
	hooks    []Hook
	mutation *RecordinsuranceMutation
}

// SetRecordinsuranceTime sets the recordinsurance_time field.
func (ruo *RecordinsuranceUpdateOne) SetRecordinsuranceTime(t time.Time) *RecordinsuranceUpdateOne {
	ruo.mutation.SetRecordinsuranceTime(t)
	return ruo
}

// SetNillableRecordinsuranceTime sets the recordinsurance_time field if the given value is not nil.
func (ruo *RecordinsuranceUpdateOne) SetNillableRecordinsuranceTime(t *time.Time) *RecordinsuranceUpdateOne {
	if t != nil {
		ruo.SetRecordinsuranceTime(*t)
	}
	return ruo
}

// SetMemberID sets the Member edge to Member by id.
func (ruo *RecordinsuranceUpdateOne) SetMemberID(id int) *RecordinsuranceUpdateOne {
	ruo.mutation.SetMemberID(id)
	return ruo
}

// SetNillableMemberID sets the Member edge to Member by id if the given value is not nil.
func (ruo *RecordinsuranceUpdateOne) SetNillableMemberID(id *int) *RecordinsuranceUpdateOne {
	if id != nil {
		ruo = ruo.SetMemberID(*id)
	}
	return ruo
}

// SetMember sets the Member edge to Member.
func (ruo *RecordinsuranceUpdateOne) SetMember(m *Member) *RecordinsuranceUpdateOne {
	return ruo.SetMemberID(m.ID)
}

// SetHospitalID sets the Hospital edge to Hospital by id.
func (ruo *RecordinsuranceUpdateOne) SetHospitalID(id int) *RecordinsuranceUpdateOne {
	ruo.mutation.SetHospitalID(id)
	return ruo
}

// SetNillableHospitalID sets the Hospital edge to Hospital by id if the given value is not nil.
func (ruo *RecordinsuranceUpdateOne) SetNillableHospitalID(id *int) *RecordinsuranceUpdateOne {
	if id != nil {
		ruo = ruo.SetHospitalID(*id)
	}
	return ruo
}

// SetHospital sets the Hospital edge to Hospital.
func (ruo *RecordinsuranceUpdateOne) SetHospital(h *Hospital) *RecordinsuranceUpdateOne {
	return ruo.SetHospitalID(h.ID)
}

// SetOfficerID sets the Officer edge to Officer by id.
func (ruo *RecordinsuranceUpdateOne) SetOfficerID(id int) *RecordinsuranceUpdateOne {
	ruo.mutation.SetOfficerID(id)
	return ruo
}

// SetNillableOfficerID sets the Officer edge to Officer by id if the given value is not nil.
func (ruo *RecordinsuranceUpdateOne) SetNillableOfficerID(id *int) *RecordinsuranceUpdateOne {
	if id != nil {
		ruo = ruo.SetOfficerID(*id)
	}
	return ruo
}

// SetOfficer sets the Officer edge to Officer.
func (ruo *RecordinsuranceUpdateOne) SetOfficer(o *Officer) *RecordinsuranceUpdateOne {
	return ruo.SetOfficerID(o.ID)
}

// SetProductID sets the Product edge to Product by id.
func (ruo *RecordinsuranceUpdateOne) SetProductID(id int) *RecordinsuranceUpdateOne {
	ruo.mutation.SetProductID(id)
	return ruo
}

// SetNillableProductID sets the Product edge to Product by id if the given value is not nil.
func (ruo *RecordinsuranceUpdateOne) SetNillableProductID(id *int) *RecordinsuranceUpdateOne {
	if id != nil {
		ruo = ruo.SetProductID(*id)
	}
	return ruo
}

// SetProduct sets the Product edge to Product.
func (ruo *RecordinsuranceUpdateOne) SetProduct(p *Product) *RecordinsuranceUpdateOne {
	return ruo.SetProductID(p.ID)
}

// SetAmountpaidID sets the Amountpaid edge to Amountpaid by id.
func (ruo *RecordinsuranceUpdateOne) SetAmountpaidID(id int) *RecordinsuranceUpdateOne {
	ruo.mutation.SetAmountpaidID(id)
	return ruo
}

// SetNillableAmountpaidID sets the Amountpaid edge to Amountpaid by id if the given value is not nil.
func (ruo *RecordinsuranceUpdateOne) SetNillableAmountpaidID(id *int) *RecordinsuranceUpdateOne {
	if id != nil {
		ruo = ruo.SetAmountpaidID(*id)
	}
	return ruo
}

// SetAmountpaid sets the Amountpaid edge to Amountpaid.
func (ruo *RecordinsuranceUpdateOne) SetAmountpaid(a *Amountpaid) *RecordinsuranceUpdateOne {
	return ruo.SetAmountpaidID(a.ID)
}

// Mutation returns the RecordinsuranceMutation object of the builder.
func (ruo *RecordinsuranceUpdateOne) Mutation() *RecordinsuranceMutation {
	return ruo.mutation
}

// ClearMember clears the Member edge to Member.
func (ruo *RecordinsuranceUpdateOne) ClearMember() *RecordinsuranceUpdateOne {
	ruo.mutation.ClearMember()
	return ruo
}

// ClearHospital clears the Hospital edge to Hospital.
func (ruo *RecordinsuranceUpdateOne) ClearHospital() *RecordinsuranceUpdateOne {
	ruo.mutation.ClearHospital()
	return ruo
}

// ClearOfficer clears the Officer edge to Officer.
func (ruo *RecordinsuranceUpdateOne) ClearOfficer() *RecordinsuranceUpdateOne {
	ruo.mutation.ClearOfficer()
	return ruo
}

// ClearProduct clears the Product edge to Product.
func (ruo *RecordinsuranceUpdateOne) ClearProduct() *RecordinsuranceUpdateOne {
	ruo.mutation.ClearProduct()
	return ruo
}

// ClearAmountpaid clears the Amountpaid edge to Amountpaid.
func (ruo *RecordinsuranceUpdateOne) ClearAmountpaid() *RecordinsuranceUpdateOne {
	ruo.mutation.ClearAmountpaid()
	return ruo
}

// Save executes the query and returns the updated entity.
func (ruo *RecordinsuranceUpdateOne) Save(ctx context.Context) (*Recordinsurance, error) {

	var (
		err  error
		node *Recordinsurance
	)
	if len(ruo.hooks) == 0 {
		node, err = ruo.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*RecordinsuranceMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			ruo.mutation = mutation
			node, err = ruo.sqlSave(ctx)
			mutation.done = true
			return node, err
		})
		for i := len(ruo.hooks) - 1; i >= 0; i-- {
			mut = ruo.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, ruo.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX is like Save, but panics if an error occurs.
func (ruo *RecordinsuranceUpdateOne) SaveX(ctx context.Context) *Recordinsurance {
	r, err := ruo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return r
}

// Exec executes the query on the entity.
func (ruo *RecordinsuranceUpdateOne) Exec(ctx context.Context) error {
	_, err := ruo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (ruo *RecordinsuranceUpdateOne) ExecX(ctx context.Context) {
	if err := ruo.Exec(ctx); err != nil {
		panic(err)
	}
}

func (ruo *RecordinsuranceUpdateOne) sqlSave(ctx context.Context) (r *Recordinsurance, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   recordinsurance.Table,
			Columns: recordinsurance.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: recordinsurance.FieldID,
			},
		},
	}
	id, ok := ruo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "ID", err: fmt.Errorf("missing Recordinsurance.ID for update")}
	}
	_spec.Node.ID.Value = id
	if value, ok := ruo.mutation.RecordinsuranceTime(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: recordinsurance.FieldRecordinsuranceTime,
		})
	}
	if ruo.mutation.MemberCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   recordinsurance.MemberTable,
			Columns: []string{recordinsurance.MemberColumn},
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
	if nodes := ruo.mutation.MemberIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   recordinsurance.MemberTable,
			Columns: []string{recordinsurance.MemberColumn},
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
	if ruo.mutation.HospitalCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   recordinsurance.HospitalTable,
			Columns: []string{recordinsurance.HospitalColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: hospital.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := ruo.mutation.HospitalIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   recordinsurance.HospitalTable,
			Columns: []string{recordinsurance.HospitalColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: hospital.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if ruo.mutation.OfficerCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   recordinsurance.OfficerTable,
			Columns: []string{recordinsurance.OfficerColumn},
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
	if nodes := ruo.mutation.OfficerIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   recordinsurance.OfficerTable,
			Columns: []string{recordinsurance.OfficerColumn},
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
	if ruo.mutation.ProductCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   recordinsurance.ProductTable,
			Columns: []string{recordinsurance.ProductColumn},
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
	if nodes := ruo.mutation.ProductIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   recordinsurance.ProductTable,
			Columns: []string{recordinsurance.ProductColumn},
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
	if ruo.mutation.AmountpaidCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   recordinsurance.AmountpaidTable,
			Columns: []string{recordinsurance.AmountpaidColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: amountpaid.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := ruo.mutation.AmountpaidIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   recordinsurance.AmountpaidTable,
			Columns: []string{recordinsurance.AmountpaidColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: amountpaid.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	r = &Recordinsurance{config: ruo.config}
	_spec.Assign = r.assignValues
	_spec.ScanValues = r.scanValues()
	if err = sqlgraph.UpdateNode(ctx, ruo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{recordinsurance.Label}
		} else if cerr, ok := isSQLConstraintError(err); ok {
			err = cerr
		}
		return nil, err
	}
	return r, nil
}
