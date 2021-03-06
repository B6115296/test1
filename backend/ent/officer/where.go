// Code generated by entc, DO NOT EDIT.

package officer

import (
	"github.com/facebookincubator/ent/dialect/sql"
	"github.com/facebookincubator/ent/dialect/sql/sqlgraph"
	"github.com/sut63/team05/ent/predicate"
)

// ID filters vertices based on their identifier.
func ID(id int) predicate.Officer {
	return predicate.Officer(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldID), id))
	})
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id int) predicate.Officer {
	return predicate.Officer(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldID), id))
	})
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id int) predicate.Officer {
	return predicate.Officer(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldID), id))
	})
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...int) predicate.Officer {
	return predicate.Officer(func(s *sql.Selector) {
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
func IDNotIn(ids ...int) predicate.Officer {
	return predicate.Officer(func(s *sql.Selector) {
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
func IDGT(id int) predicate.Officer {
	return predicate.Officer(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldID), id))
	})
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id int) predicate.Officer {
	return predicate.Officer(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldID), id))
	})
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id int) predicate.Officer {
	return predicate.Officer(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldID), id))
	})
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id int) predicate.Officer {
	return predicate.Officer(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldID), id))
	})
}

// OfficerEmail applies equality check predicate on the "officer_email" field. It's identical to OfficerEmailEQ.
func OfficerEmail(v string) predicate.Officer {
	return predicate.Officer(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldOfficerEmail), v))
	})
}

// OfficerName applies equality check predicate on the "officer_name" field. It's identical to OfficerNameEQ.
func OfficerName(v string) predicate.Officer {
	return predicate.Officer(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldOfficerName), v))
	})
}

// OfficerPassword applies equality check predicate on the "officer_password" field. It's identical to OfficerPasswordEQ.
func OfficerPassword(v string) predicate.Officer {
	return predicate.Officer(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldOfficerPassword), v))
	})
}

// OfficerEmailEQ applies the EQ predicate on the "officer_email" field.
func OfficerEmailEQ(v string) predicate.Officer {
	return predicate.Officer(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldOfficerEmail), v))
	})
}

// OfficerEmailNEQ applies the NEQ predicate on the "officer_email" field.
func OfficerEmailNEQ(v string) predicate.Officer {
	return predicate.Officer(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldOfficerEmail), v))
	})
}

// OfficerEmailIn applies the In predicate on the "officer_email" field.
func OfficerEmailIn(vs ...string) predicate.Officer {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Officer(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldOfficerEmail), v...))
	})
}

// OfficerEmailNotIn applies the NotIn predicate on the "officer_email" field.
func OfficerEmailNotIn(vs ...string) predicate.Officer {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Officer(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldOfficerEmail), v...))
	})
}

// OfficerEmailGT applies the GT predicate on the "officer_email" field.
func OfficerEmailGT(v string) predicate.Officer {
	return predicate.Officer(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldOfficerEmail), v))
	})
}

// OfficerEmailGTE applies the GTE predicate on the "officer_email" field.
func OfficerEmailGTE(v string) predicate.Officer {
	return predicate.Officer(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldOfficerEmail), v))
	})
}

// OfficerEmailLT applies the LT predicate on the "officer_email" field.
func OfficerEmailLT(v string) predicate.Officer {
	return predicate.Officer(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldOfficerEmail), v))
	})
}

// OfficerEmailLTE applies the LTE predicate on the "officer_email" field.
func OfficerEmailLTE(v string) predicate.Officer {
	return predicate.Officer(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldOfficerEmail), v))
	})
}

// OfficerEmailContains applies the Contains predicate on the "officer_email" field.
func OfficerEmailContains(v string) predicate.Officer {
	return predicate.Officer(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldOfficerEmail), v))
	})
}

// OfficerEmailHasPrefix applies the HasPrefix predicate on the "officer_email" field.
func OfficerEmailHasPrefix(v string) predicate.Officer {
	return predicate.Officer(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldOfficerEmail), v))
	})
}

// OfficerEmailHasSuffix applies the HasSuffix predicate on the "officer_email" field.
func OfficerEmailHasSuffix(v string) predicate.Officer {
	return predicate.Officer(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldOfficerEmail), v))
	})
}

// OfficerEmailEqualFold applies the EqualFold predicate on the "officer_email" field.
func OfficerEmailEqualFold(v string) predicate.Officer {
	return predicate.Officer(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldOfficerEmail), v))
	})
}

// OfficerEmailContainsFold applies the ContainsFold predicate on the "officer_email" field.
func OfficerEmailContainsFold(v string) predicate.Officer {
	return predicate.Officer(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldOfficerEmail), v))
	})
}

// OfficerNameEQ applies the EQ predicate on the "officer_name" field.
func OfficerNameEQ(v string) predicate.Officer {
	return predicate.Officer(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldOfficerName), v))
	})
}

// OfficerNameNEQ applies the NEQ predicate on the "officer_name" field.
func OfficerNameNEQ(v string) predicate.Officer {
	return predicate.Officer(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldOfficerName), v))
	})
}

// OfficerNameIn applies the In predicate on the "officer_name" field.
func OfficerNameIn(vs ...string) predicate.Officer {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Officer(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldOfficerName), v...))
	})
}

// OfficerNameNotIn applies the NotIn predicate on the "officer_name" field.
func OfficerNameNotIn(vs ...string) predicate.Officer {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Officer(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldOfficerName), v...))
	})
}

// OfficerNameGT applies the GT predicate on the "officer_name" field.
func OfficerNameGT(v string) predicate.Officer {
	return predicate.Officer(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldOfficerName), v))
	})
}

// OfficerNameGTE applies the GTE predicate on the "officer_name" field.
func OfficerNameGTE(v string) predicate.Officer {
	return predicate.Officer(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldOfficerName), v))
	})
}

// OfficerNameLT applies the LT predicate on the "officer_name" field.
func OfficerNameLT(v string) predicate.Officer {
	return predicate.Officer(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldOfficerName), v))
	})
}

// OfficerNameLTE applies the LTE predicate on the "officer_name" field.
func OfficerNameLTE(v string) predicate.Officer {
	return predicate.Officer(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldOfficerName), v))
	})
}

// OfficerNameContains applies the Contains predicate on the "officer_name" field.
func OfficerNameContains(v string) predicate.Officer {
	return predicate.Officer(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldOfficerName), v))
	})
}

// OfficerNameHasPrefix applies the HasPrefix predicate on the "officer_name" field.
func OfficerNameHasPrefix(v string) predicate.Officer {
	return predicate.Officer(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldOfficerName), v))
	})
}

// OfficerNameHasSuffix applies the HasSuffix predicate on the "officer_name" field.
func OfficerNameHasSuffix(v string) predicate.Officer {
	return predicate.Officer(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldOfficerName), v))
	})
}

// OfficerNameEqualFold applies the EqualFold predicate on the "officer_name" field.
func OfficerNameEqualFold(v string) predicate.Officer {
	return predicate.Officer(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldOfficerName), v))
	})
}

// OfficerNameContainsFold applies the ContainsFold predicate on the "officer_name" field.
func OfficerNameContainsFold(v string) predicate.Officer {
	return predicate.Officer(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldOfficerName), v))
	})
}

// OfficerPasswordEQ applies the EQ predicate on the "officer_password" field.
func OfficerPasswordEQ(v string) predicate.Officer {
	return predicate.Officer(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldOfficerPassword), v))
	})
}

// OfficerPasswordNEQ applies the NEQ predicate on the "officer_password" field.
func OfficerPasswordNEQ(v string) predicate.Officer {
	return predicate.Officer(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldOfficerPassword), v))
	})
}

// OfficerPasswordIn applies the In predicate on the "officer_password" field.
func OfficerPasswordIn(vs ...string) predicate.Officer {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Officer(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldOfficerPassword), v...))
	})
}

// OfficerPasswordNotIn applies the NotIn predicate on the "officer_password" field.
func OfficerPasswordNotIn(vs ...string) predicate.Officer {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Officer(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldOfficerPassword), v...))
	})
}

// OfficerPasswordGT applies the GT predicate on the "officer_password" field.
func OfficerPasswordGT(v string) predicate.Officer {
	return predicate.Officer(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldOfficerPassword), v))
	})
}

// OfficerPasswordGTE applies the GTE predicate on the "officer_password" field.
func OfficerPasswordGTE(v string) predicate.Officer {
	return predicate.Officer(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldOfficerPassword), v))
	})
}

// OfficerPasswordLT applies the LT predicate on the "officer_password" field.
func OfficerPasswordLT(v string) predicate.Officer {
	return predicate.Officer(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldOfficerPassword), v))
	})
}

// OfficerPasswordLTE applies the LTE predicate on the "officer_password" field.
func OfficerPasswordLTE(v string) predicate.Officer {
	return predicate.Officer(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldOfficerPassword), v))
	})
}

// OfficerPasswordContains applies the Contains predicate on the "officer_password" field.
func OfficerPasswordContains(v string) predicate.Officer {
	return predicate.Officer(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldOfficerPassword), v))
	})
}

// OfficerPasswordHasPrefix applies the HasPrefix predicate on the "officer_password" field.
func OfficerPasswordHasPrefix(v string) predicate.Officer {
	return predicate.Officer(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldOfficerPassword), v))
	})
}

// OfficerPasswordHasSuffix applies the HasSuffix predicate on the "officer_password" field.
func OfficerPasswordHasSuffix(v string) predicate.Officer {
	return predicate.Officer(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldOfficerPassword), v))
	})
}

// OfficerPasswordEqualFold applies the EqualFold predicate on the "officer_password" field.
func OfficerPasswordEqualFold(v string) predicate.Officer {
	return predicate.Officer(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldOfficerPassword), v))
	})
}

// OfficerPasswordContainsFold applies the ContainsFold predicate on the "officer_password" field.
func OfficerPasswordContainsFold(v string) predicate.Officer {
	return predicate.Officer(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldOfficerPassword), v))
	})
}

// HasOfficerProduct applies the HasEdge predicate on the "officer_product" edge.
func HasOfficerProduct() predicate.Officer {
	return predicate.Officer(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(OfficerProductTable, FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, OfficerProductTable, OfficerProductColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasOfficerProductWith applies the HasEdge predicate on the "officer_product" edge with a given conditions (other predicates).
func HasOfficerProductWith(preds ...predicate.Product) predicate.Officer {
	return predicate.Officer(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(OfficerProductInverseTable, FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, OfficerProductTable, OfficerProductColumn),
		)
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// HasOfficerInsurance applies the HasEdge predicate on the "officer_insurance" edge.
func HasOfficerInsurance() predicate.Officer {
	return predicate.Officer(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(OfficerInsuranceTable, FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, OfficerInsuranceTable, OfficerInsuranceColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasOfficerInsuranceWith applies the HasEdge predicate on the "officer_insurance" edge with a given conditions (other predicates).
func HasOfficerInsuranceWith(preds ...predicate.Insurance) predicate.Officer {
	return predicate.Officer(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(OfficerInsuranceInverseTable, FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, OfficerInsuranceTable, OfficerInsuranceColumn),
		)
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// HasOfficerInquiry applies the HasEdge predicate on the "officer_inquiry" edge.
func HasOfficerInquiry() predicate.Officer {
	return predicate.Officer(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(OfficerInquiryTable, FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, OfficerInquiryTable, OfficerInquiryColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasOfficerInquiryWith applies the HasEdge predicate on the "officer_inquiry" edge with a given conditions (other predicates).
func HasOfficerInquiryWith(preds ...predicate.Inquiry) predicate.Officer {
	return predicate.Officer(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(OfficerInquiryInverseTable, FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, OfficerInquiryTable, OfficerInquiryColumn),
		)
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// HasOfficerPayback applies the HasEdge predicate on the "officer_payback" edge.
func HasOfficerPayback() predicate.Officer {
	return predicate.Officer(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(OfficerPaybackTable, FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, OfficerPaybackTable, OfficerPaybackColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasOfficerPaybackWith applies the HasEdge predicate on the "officer_payback" edge with a given conditions (other predicates).
func HasOfficerPaybackWith(preds ...predicate.Payback) predicate.Officer {
	return predicate.Officer(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(OfficerPaybackInverseTable, FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, OfficerPaybackTable, OfficerPaybackColumn),
		)
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// HasOfficerRecordinsurance applies the HasEdge predicate on the "officer_recordinsurance" edge.
func HasOfficerRecordinsurance() predicate.Officer {
	return predicate.Officer(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(OfficerRecordinsuranceTable, FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, OfficerRecordinsuranceTable, OfficerRecordinsuranceColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasOfficerRecordinsuranceWith applies the HasEdge predicate on the "officer_recordinsurance" edge with a given conditions (other predicates).
func HasOfficerRecordinsuranceWith(preds ...predicate.Recordinsurance) predicate.Officer {
	return predicate.Officer(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(OfficerRecordinsuranceInverseTable, FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, OfficerRecordinsuranceTable, OfficerRecordinsuranceColumn),
		)
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// And groups list of predicates with the AND operator between them.
func And(predicates ...predicate.Officer) predicate.Officer {
	return predicate.Officer(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for _, p := range predicates {
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Or groups list of predicates with the OR operator between them.
func Or(predicates ...predicate.Officer) predicate.Officer {
	return predicate.Officer(func(s *sql.Selector) {
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
func Not(p predicate.Officer) predicate.Officer {
	return predicate.Officer(func(s *sql.Selector) {
		p(s.Not())
	})
}
