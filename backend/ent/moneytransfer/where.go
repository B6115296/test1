// Code generated by entc, DO NOT EDIT.

package moneytransfer

import (
	"github.com/facebookincubator/ent/dialect/sql"
	"github.com/facebookincubator/ent/dialect/sql/sqlgraph"
	"github.com/sut63/team05/ent/predicate"
)

// ID filters vertices based on their identifier.
func ID(id int) predicate.MoneyTransfer {
	return predicate.MoneyTransfer(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldID), id))
	})
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id int) predicate.MoneyTransfer {
	return predicate.MoneyTransfer(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldID), id))
	})
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id int) predicate.MoneyTransfer {
	return predicate.MoneyTransfer(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldID), id))
	})
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...int) predicate.MoneyTransfer {
	return predicate.MoneyTransfer(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(ids) == 0 {
			s.Where(sql.False())
			return
		}
		v := make([]interface{}, len(ids))
		for i := range v {
			v[i] = ids[i]
		}
		s.Where(sql.In(s.C(FieldID), v...))
	})
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...int) predicate.MoneyTransfer {
	return predicate.MoneyTransfer(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(ids) == 0 {
			s.Where(sql.False())
			return
		}
		v := make([]interface{}, len(ids))
		for i := range v {
			v[i] = ids[i]
		}
		s.Where(sql.NotIn(s.C(FieldID), v...))
	})
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id int) predicate.MoneyTransfer {
	return predicate.MoneyTransfer(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldID), id))
	})
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id int) predicate.MoneyTransfer {
	return predicate.MoneyTransfer(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldID), id))
	})
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id int) predicate.MoneyTransfer {
	return predicate.MoneyTransfer(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldID), id))
	})
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id int) predicate.MoneyTransfer {
	return predicate.MoneyTransfer(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldID), id))
	})
}

// MoneytransferType applies equality check predicate on the "moneytransfer_type" field. It's identical to MoneytransferTypeEQ.
func MoneytransferType(v string) predicate.MoneyTransfer {
	return predicate.MoneyTransfer(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldMoneytransferType), v))
	})
}

// MoneytransferTypeEQ applies the EQ predicate on the "moneytransfer_type" field.
func MoneytransferTypeEQ(v string) predicate.MoneyTransfer {
	return predicate.MoneyTransfer(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldMoneytransferType), v))
	})
}

// MoneytransferTypeNEQ applies the NEQ predicate on the "moneytransfer_type" field.
func MoneytransferTypeNEQ(v string) predicate.MoneyTransfer {
	return predicate.MoneyTransfer(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldMoneytransferType), v))
	})
}

// MoneytransferTypeIn applies the In predicate on the "moneytransfer_type" field.
func MoneytransferTypeIn(vs ...string) predicate.MoneyTransfer {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.MoneyTransfer(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldMoneytransferType), v...))
	})
}

// MoneytransferTypeNotIn applies the NotIn predicate on the "moneytransfer_type" field.
func MoneytransferTypeNotIn(vs ...string) predicate.MoneyTransfer {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.MoneyTransfer(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldMoneytransferType), v...))
	})
}

// MoneytransferTypeGT applies the GT predicate on the "moneytransfer_type" field.
func MoneytransferTypeGT(v string) predicate.MoneyTransfer {
	return predicate.MoneyTransfer(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldMoneytransferType), v))
	})
}

// MoneytransferTypeGTE applies the GTE predicate on the "moneytransfer_type" field.
func MoneytransferTypeGTE(v string) predicate.MoneyTransfer {
	return predicate.MoneyTransfer(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldMoneytransferType), v))
	})
}

// MoneytransferTypeLT applies the LT predicate on the "moneytransfer_type" field.
func MoneytransferTypeLT(v string) predicate.MoneyTransfer {
	return predicate.MoneyTransfer(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldMoneytransferType), v))
	})
}

// MoneytransferTypeLTE applies the LTE predicate on the "moneytransfer_type" field.
func MoneytransferTypeLTE(v string) predicate.MoneyTransfer {
	return predicate.MoneyTransfer(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldMoneytransferType), v))
	})
}

// MoneytransferTypeContains applies the Contains predicate on the "moneytransfer_type" field.
func MoneytransferTypeContains(v string) predicate.MoneyTransfer {
	return predicate.MoneyTransfer(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldMoneytransferType), v))
	})
}

// MoneytransferTypeHasPrefix applies the HasPrefix predicate on the "moneytransfer_type" field.
func MoneytransferTypeHasPrefix(v string) predicate.MoneyTransfer {
	return predicate.MoneyTransfer(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldMoneytransferType), v))
	})
}

// MoneytransferTypeHasSuffix applies the HasSuffix predicate on the "moneytransfer_type" field.
func MoneytransferTypeHasSuffix(v string) predicate.MoneyTransfer {
	return predicate.MoneyTransfer(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldMoneytransferType), v))
	})
}

// MoneytransferTypeEqualFold applies the EqualFold predicate on the "moneytransfer_type" field.
func MoneytransferTypeEqualFold(v string) predicate.MoneyTransfer {
	return predicate.MoneyTransfer(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldMoneytransferType), v))
	})
}

// MoneytransferTypeContainsFold applies the ContainsFold predicate on the "moneytransfer_type" field.
func MoneytransferTypeContainsFold(v string) predicate.MoneyTransfer {
	return predicate.MoneyTransfer(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldMoneytransferType), v))
	})
}

// HasMoneytransferPayment applies the HasEdge predicate on the "moneytransfer_payment" edge.
func HasMoneytransferPayment() predicate.MoneyTransfer {
	return predicate.MoneyTransfer(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(MoneytransferPaymentTable, FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, MoneytransferPaymentTable, MoneytransferPaymentColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasMoneytransferPaymentWith applies the HasEdge predicate on the "moneytransfer_payment" edge with a given conditions (other predicates).
func HasMoneytransferPaymentWith(preds ...predicate.Payment) predicate.MoneyTransfer {
	return predicate.MoneyTransfer(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(MoneytransferPaymentInverseTable, FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, MoneytransferPaymentTable, MoneytransferPaymentColumn),
		)
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// And groups list of predicates with the AND operator between them.
func And(predicates ...predicate.MoneyTransfer) predicate.MoneyTransfer {
	return predicate.MoneyTransfer(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for _, p := range predicates {
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Or groups list of predicates with the OR operator between them.
func Or(predicates ...predicate.MoneyTransfer) predicate.MoneyTransfer {
	return predicate.MoneyTransfer(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for i, p := range predicates {
			if i > 0 {
				s1.Or()
			}
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Not applies the not operator on the given predicate.
func Not(p predicate.MoneyTransfer) predicate.MoneyTransfer {
	return predicate.MoneyTransfer(func(s *sql.Selector) {
		p(s.Not())
	})
}