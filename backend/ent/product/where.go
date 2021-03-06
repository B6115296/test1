// Code generated by entc, DO NOT EDIT.

package product

import (
	"github.com/facebookincubator/ent/dialect/sql"
	"github.com/facebookincubator/ent/dialect/sql/sqlgraph"
	"github.com/sut63/team05/ent/predicate"
)

// ID filters vertices based on their identifier.
func ID(id int) predicate.Product {
	return predicate.Product(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldID), id))
	})
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id int) predicate.Product {
	return predicate.Product(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldID), id))
	})
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id int) predicate.Product {
	return predicate.Product(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldID), id))
	})
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...int) predicate.Product {
	return predicate.Product(func(s *sql.Selector) {
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
func IDNotIn(ids ...int) predicate.Product {
	return predicate.Product(func(s *sql.Selector) {
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
func IDGT(id int) predicate.Product {
	return predicate.Product(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldID), id))
	})
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id int) predicate.Product {
	return predicate.Product(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldID), id))
	})
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id int) predicate.Product {
	return predicate.Product(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldID), id))
	})
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id int) predicate.Product {
	return predicate.Product(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldID), id))
	})
}

// ProductName applies equality check predicate on the "product_name" field. It's identical to ProductNameEQ.
func ProductName(v string) predicate.Product {
	return predicate.Product(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldProductName), v))
	})
}

// ProductPrice applies equality check predicate on the "product_price" field. It's identical to ProductPriceEQ.
func ProductPrice(v int) predicate.Product {
	return predicate.Product(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldProductPrice), v))
	})
}

// ProductTime applies equality check predicate on the "product_time" field. It's identical to ProductTimeEQ.
func ProductTime(v int) predicate.Product {
	return predicate.Product(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldProductTime), v))
	})
}

// ProductPaymentOfYear applies equality check predicate on the "product_payment_of_year" field. It's identical to ProductPaymentOfYearEQ.
func ProductPaymentOfYear(v float64) predicate.Product {
	return predicate.Product(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldProductPaymentOfYear), v))
	})
}

// ProductNameEQ applies the EQ predicate on the "product_name" field.
func ProductNameEQ(v string) predicate.Product {
	return predicate.Product(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldProductName), v))
	})
}

// ProductNameNEQ applies the NEQ predicate on the "product_name" field.
func ProductNameNEQ(v string) predicate.Product {
	return predicate.Product(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldProductName), v))
	})
}

// ProductNameIn applies the In predicate on the "product_name" field.
func ProductNameIn(vs ...string) predicate.Product {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Product(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldProductName), v...))
	})
}

// ProductNameNotIn applies the NotIn predicate on the "product_name" field.
func ProductNameNotIn(vs ...string) predicate.Product {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Product(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldProductName), v...))
	})
}

// ProductNameGT applies the GT predicate on the "product_name" field.
func ProductNameGT(v string) predicate.Product {
	return predicate.Product(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldProductName), v))
	})
}

// ProductNameGTE applies the GTE predicate on the "product_name" field.
func ProductNameGTE(v string) predicate.Product {
	return predicate.Product(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldProductName), v))
	})
}

// ProductNameLT applies the LT predicate on the "product_name" field.
func ProductNameLT(v string) predicate.Product {
	return predicate.Product(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldProductName), v))
	})
}

// ProductNameLTE applies the LTE predicate on the "product_name" field.
func ProductNameLTE(v string) predicate.Product {
	return predicate.Product(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldProductName), v))
	})
}

// ProductNameContains applies the Contains predicate on the "product_name" field.
func ProductNameContains(v string) predicate.Product {
	return predicate.Product(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldProductName), v))
	})
}

// ProductNameHasPrefix applies the HasPrefix predicate on the "product_name" field.
func ProductNameHasPrefix(v string) predicate.Product {
	return predicate.Product(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldProductName), v))
	})
}

// ProductNameHasSuffix applies the HasSuffix predicate on the "product_name" field.
func ProductNameHasSuffix(v string) predicate.Product {
	return predicate.Product(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldProductName), v))
	})
}

// ProductNameEqualFold applies the EqualFold predicate on the "product_name" field.
func ProductNameEqualFold(v string) predicate.Product {
	return predicate.Product(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldProductName), v))
	})
}

// ProductNameContainsFold applies the ContainsFold predicate on the "product_name" field.
func ProductNameContainsFold(v string) predicate.Product {
	return predicate.Product(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldProductName), v))
	})
}

// ProductPriceEQ applies the EQ predicate on the "product_price" field.
func ProductPriceEQ(v int) predicate.Product {
	return predicate.Product(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldProductPrice), v))
	})
}

// ProductPriceNEQ applies the NEQ predicate on the "product_price" field.
func ProductPriceNEQ(v int) predicate.Product {
	return predicate.Product(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldProductPrice), v))
	})
}

// ProductPriceIn applies the In predicate on the "product_price" field.
func ProductPriceIn(vs ...int) predicate.Product {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Product(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldProductPrice), v...))
	})
}

// ProductPriceNotIn applies the NotIn predicate on the "product_price" field.
func ProductPriceNotIn(vs ...int) predicate.Product {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Product(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldProductPrice), v...))
	})
}

// ProductPriceGT applies the GT predicate on the "product_price" field.
func ProductPriceGT(v int) predicate.Product {
	return predicate.Product(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldProductPrice), v))
	})
}

// ProductPriceGTE applies the GTE predicate on the "product_price" field.
func ProductPriceGTE(v int) predicate.Product {
	return predicate.Product(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldProductPrice), v))
	})
}

// ProductPriceLT applies the LT predicate on the "product_price" field.
func ProductPriceLT(v int) predicate.Product {
	return predicate.Product(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldProductPrice), v))
	})
}

// ProductPriceLTE applies the LTE predicate on the "product_price" field.
func ProductPriceLTE(v int) predicate.Product {
	return predicate.Product(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldProductPrice), v))
	})
}

// ProductTimeEQ applies the EQ predicate on the "product_time" field.
func ProductTimeEQ(v int) predicate.Product {
	return predicate.Product(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldProductTime), v))
	})
}

// ProductTimeNEQ applies the NEQ predicate on the "product_time" field.
func ProductTimeNEQ(v int) predicate.Product {
	return predicate.Product(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldProductTime), v))
	})
}

// ProductTimeIn applies the In predicate on the "product_time" field.
func ProductTimeIn(vs ...int) predicate.Product {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Product(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldProductTime), v...))
	})
}

// ProductTimeNotIn applies the NotIn predicate on the "product_time" field.
func ProductTimeNotIn(vs ...int) predicate.Product {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Product(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldProductTime), v...))
	})
}

// ProductTimeGT applies the GT predicate on the "product_time" field.
func ProductTimeGT(v int) predicate.Product {
	return predicate.Product(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldProductTime), v))
	})
}

// ProductTimeGTE applies the GTE predicate on the "product_time" field.
func ProductTimeGTE(v int) predicate.Product {
	return predicate.Product(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldProductTime), v))
	})
}

// ProductTimeLT applies the LT predicate on the "product_time" field.
func ProductTimeLT(v int) predicate.Product {
	return predicate.Product(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldProductTime), v))
	})
}

// ProductTimeLTE applies the LTE predicate on the "product_time" field.
func ProductTimeLTE(v int) predicate.Product {
	return predicate.Product(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldProductTime), v))
	})
}

// ProductPaymentOfYearEQ applies the EQ predicate on the "product_payment_of_year" field.
func ProductPaymentOfYearEQ(v float64) predicate.Product {
	return predicate.Product(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldProductPaymentOfYear), v))
	})
}

// ProductPaymentOfYearNEQ applies the NEQ predicate on the "product_payment_of_year" field.
func ProductPaymentOfYearNEQ(v float64) predicate.Product {
	return predicate.Product(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldProductPaymentOfYear), v))
	})
}

// ProductPaymentOfYearIn applies the In predicate on the "product_payment_of_year" field.
func ProductPaymentOfYearIn(vs ...float64) predicate.Product {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Product(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldProductPaymentOfYear), v...))
	})
}

// ProductPaymentOfYearNotIn applies the NotIn predicate on the "product_payment_of_year" field.
func ProductPaymentOfYearNotIn(vs ...float64) predicate.Product {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Product(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldProductPaymentOfYear), v...))
	})
}

// ProductPaymentOfYearGT applies the GT predicate on the "product_payment_of_year" field.
func ProductPaymentOfYearGT(v float64) predicate.Product {
	return predicate.Product(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldProductPaymentOfYear), v))
	})
}

// ProductPaymentOfYearGTE applies the GTE predicate on the "product_payment_of_year" field.
func ProductPaymentOfYearGTE(v float64) predicate.Product {
	return predicate.Product(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldProductPaymentOfYear), v))
	})
}

// ProductPaymentOfYearLT applies the LT predicate on the "product_payment_of_year" field.
func ProductPaymentOfYearLT(v float64) predicate.Product {
	return predicate.Product(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldProductPaymentOfYear), v))
	})
}

// ProductPaymentOfYearLTE applies the LTE predicate on the "product_payment_of_year" field.
func ProductPaymentOfYearLTE(v float64) predicate.Product {
	return predicate.Product(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldProductPaymentOfYear), v))
	})
}

// HasProductGender applies the HasEdge predicate on the "product_gender" edge.
func HasProductGender() predicate.Product {
	return predicate.Product(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(ProductGenderTable, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, ProductGenderTable, ProductGenderColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasProductGenderWith applies the HasEdge predicate on the "product_gender" edge with a given conditions (other predicates).
func HasProductGenderWith(preds ...predicate.Gender) predicate.Product {
	return predicate.Product(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(ProductGenderInverseTable, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, ProductGenderTable, ProductGenderColumn),
		)
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// HasProductGroupage applies the HasEdge predicate on the "product_groupage" edge.
func HasProductGroupage() predicate.Product {
	return predicate.Product(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(ProductGroupageTable, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, ProductGroupageTable, ProductGroupageColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasProductGroupageWith applies the HasEdge predicate on the "product_groupage" edge with a given conditions (other predicates).
func HasProductGroupageWith(preds ...predicate.GroupOfAge) predicate.Product {
	return predicate.Product(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(ProductGroupageInverseTable, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, ProductGroupageTable, ProductGroupageColumn),
		)
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// HasProductOfficer applies the HasEdge predicate on the "product_officer" edge.
func HasProductOfficer() predicate.Product {
	return predicate.Product(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(ProductOfficerTable, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, ProductOfficerTable, ProductOfficerColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasProductOfficerWith applies the HasEdge predicate on the "product_officer" edge with a given conditions (other predicates).
func HasProductOfficerWith(preds ...predicate.Officer) predicate.Product {
	return predicate.Product(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(ProductOfficerInverseTable, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, ProductOfficerTable, ProductOfficerColumn),
		)
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// HasProductInsurance applies the HasEdge predicate on the "product_insurance" edge.
func HasProductInsurance() predicate.Product {
	return predicate.Product(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(ProductInsuranceTable, FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, ProductInsuranceTable, ProductInsuranceColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasProductInsuranceWith applies the HasEdge predicate on the "product_insurance" edge with a given conditions (other predicates).
func HasProductInsuranceWith(preds ...predicate.Insurance) predicate.Product {
	return predicate.Product(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(ProductInsuranceInverseTable, FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, ProductInsuranceTable, ProductInsuranceColumn),
		)
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// HasProductInquiry applies the HasEdge predicate on the "product_inquiry" edge.
func HasProductInquiry() predicate.Product {
	return predicate.Product(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(ProductInquiryTable, FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, ProductInquiryTable, ProductInquiryColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasProductInquiryWith applies the HasEdge predicate on the "product_inquiry" edge with a given conditions (other predicates).
func HasProductInquiryWith(preds ...predicate.Inquiry) predicate.Product {
	return predicate.Product(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(ProductInquiryInverseTable, FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, ProductInquiryTable, ProductInquiryColumn),
		)
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// HasProductPayback applies the HasEdge predicate on the "product_payback" edge.
func HasProductPayback() predicate.Product {
	return predicate.Product(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(ProductPaybackTable, FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, ProductPaybackTable, ProductPaybackColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasProductPaybackWith applies the HasEdge predicate on the "product_payback" edge with a given conditions (other predicates).
func HasProductPaybackWith(preds ...predicate.Payback) predicate.Product {
	return predicate.Product(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(ProductPaybackInverseTable, FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, ProductPaybackTable, ProductPaybackColumn),
		)
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// HasProductRecordinsurance applies the HasEdge predicate on the "product_recordinsurance" edge.
func HasProductRecordinsurance() predicate.Product {
	return predicate.Product(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(ProductRecordinsuranceTable, FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, ProductRecordinsuranceTable, ProductRecordinsuranceColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasProductRecordinsuranceWith applies the HasEdge predicate on the "product_recordinsurance" edge with a given conditions (other predicates).
func HasProductRecordinsuranceWith(preds ...predicate.Recordinsurance) predicate.Product {
	return predicate.Product(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(ProductRecordinsuranceInverseTable, FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, ProductRecordinsuranceTable, ProductRecordinsuranceColumn),
		)
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// And groups list of predicates with the AND operator between them.
func And(predicates ...predicate.Product) predicate.Product {
	return predicate.Product(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for _, p := range predicates {
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Or groups list of predicates with the OR operator between them.
func Or(predicates ...predicate.Product) predicate.Product {
	return predicate.Product(func(s *sql.Selector) {
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
func Not(p predicate.Product) predicate.Product {
	return predicate.Product(func(s *sql.Selector) {
		p(s.Not())
	})
}
