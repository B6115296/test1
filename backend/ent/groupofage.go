// Code generated by entc, DO NOT EDIT.

package ent

import (
	"fmt"
	"strings"

	"github.com/facebookincubator/ent/dialect/sql"
	"github.com/sut63/team05/ent/groupofage"
)

// GroupOfAge is the model entity for the GroupOfAge schema.
type GroupOfAge struct {
	config `json:"-"`
	// ID of the ent.
	ID int `json:"id,omitempty"`
	// GroupOfAgeName holds the value of the "group_of_age_name" field.
	GroupOfAgeName string `json:"group_of_age_name,omitempty"`
	// GroupOfAgeAge holds the value of the "group_of_age_age" field.
	GroupOfAgeAge string `json:"group_of_age_age,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the GroupOfAgeQuery when eager-loading is set.
	Edges GroupOfAgeEdges `json:"edges"`
}

// GroupOfAgeEdges holds the relations/edges for other nodes in the graph.
type GroupOfAgeEdges struct {
	// GroupageProduct holds the value of the groupage_product edge.
	GroupageProduct []*Product
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [1]bool
}

// GroupageProductOrErr returns the GroupageProduct value or an error if the edge
// was not loaded in eager-loading.
func (e GroupOfAgeEdges) GroupageProductOrErr() ([]*Product, error) {
	if e.loadedTypes[0] {
		return e.GroupageProduct, nil
	}
	return nil, &NotLoadedError{edge: "groupage_product"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*GroupOfAge) scanValues() []interface{} {
	return []interface{}{
		&sql.NullInt64{},  // id
		&sql.NullString{}, // group_of_age_name
		&sql.NullString{}, // group_of_age_age
	}
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the GroupOfAge fields.
func (goa *GroupOfAge) assignValues(values ...interface{}) error {
	if m, n := len(values), len(groupofage.Columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	value, ok := values[0].(*sql.NullInt64)
	if !ok {
		return fmt.Errorf("unexpected type %T for field id", value)
	}
	goa.ID = int(value.Int64)
	values = values[1:]
	if value, ok := values[0].(*sql.NullString); !ok {
		return fmt.Errorf("unexpected type %T for field group_of_age_name", values[0])
	} else if value.Valid {
		goa.GroupOfAgeName = value.String
	}
	if value, ok := values[1].(*sql.NullString); !ok {
		return fmt.Errorf("unexpected type %T for field group_of_age_age", values[1])
	} else if value.Valid {
		goa.GroupOfAgeAge = value.String
	}
	return nil
}

// QueryGroupageProduct queries the groupage_product edge of the GroupOfAge.
func (goa *GroupOfAge) QueryGroupageProduct() *ProductQuery {
	return (&GroupOfAgeClient{config: goa.config}).QueryGroupageProduct(goa)
}

// Update returns a builder for updating this GroupOfAge.
// Note that, you need to call GroupOfAge.Unwrap() before calling this method, if this GroupOfAge
// was returned from a transaction, and the transaction was committed or rolled back.
func (goa *GroupOfAge) Update() *GroupOfAgeUpdateOne {
	return (&GroupOfAgeClient{config: goa.config}).UpdateOne(goa)
}

// Unwrap unwraps the entity that was returned from a transaction after it was closed,
// so that all next queries will be executed through the driver which created the transaction.
func (goa *GroupOfAge) Unwrap() *GroupOfAge {
	tx, ok := goa.config.driver.(*txDriver)
	if !ok {
		panic("ent: GroupOfAge is not a transactional entity")
	}
	goa.config.driver = tx.drv
	return goa
}

// String implements the fmt.Stringer.
func (goa *GroupOfAge) String() string {
	var builder strings.Builder
	builder.WriteString("GroupOfAge(")
	builder.WriteString(fmt.Sprintf("id=%v", goa.ID))
	builder.WriteString(", group_of_age_name=")
	builder.WriteString(goa.GroupOfAgeName)
	builder.WriteString(", group_of_age_age=")
	builder.WriteString(goa.GroupOfAgeAge)
	builder.WriteByte(')')
	return builder.String()
}

// GroupOfAges is a parsable slice of GroupOfAge.
type GroupOfAges []*GroupOfAge

func (goa GroupOfAges) config(cfg config) {
	for _i := range goa {
		goa[_i].config = cfg
	}
}