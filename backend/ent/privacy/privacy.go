// Code generated by entc, DO NOT EDIT.

package privacy

import (
	"context"
	"errors"
	"fmt"

	"github.com/sut63/team05/ent"
)

var (
	// Allow may be returned by rules to indicate that the policy
	// evaluation should terminate with an allow decision.
	Allow = errors.New("ent/privacy: allow rule")

	// Deny may be returned by rules to indicate that the policy
	// evaluation should terminate with an deny decision.
	Deny = errors.New("ent/privacy: deny rule")

	// Skip may be returned by rules to indicate that the policy
	// evaluation should continue to the next rule.
	Skip = errors.New("ent/privacy: skip rule")
)

// Allowf returns an formatted wrapped Allow decision.
func Allowf(format string, a ...interface{}) error {
	return fmt.Errorf(format+": %w", append(a, Allow)...)
}

// Denyf returns an formatted wrapped Deny decision.
func Denyf(format string, a ...interface{}) error {
	return fmt.Errorf(format+": %w", append(a, Deny)...)
}

// Skipf returns an formatted wrapped Skip decision.
func Skipf(format string, a ...interface{}) error {
	return fmt.Errorf(format+": %w", append(a, Skip)...)
}

type decisionCtxKey struct{}

// DecisionContext creates a decision context.
func DecisionContext(parent context.Context, decision error) context.Context {
	if decision == nil || errors.Is(decision, Skip) {
		return parent
	}
	return context.WithValue(parent, decisionCtxKey{}, decision)
}

func decisionFromContext(ctx context.Context) (error, bool) {
	decision, ok := ctx.Value(decisionCtxKey{}).(error)
	if ok && errors.Is(decision, Allow) {
		decision = nil
	}
	return decision, ok
}

type (
	// QueryPolicy combines multiple query rules into a single policy.
	QueryPolicy []QueryRule

	// QueryRule defines the interface deciding whether a
	// query is allowed and optionally modify it.
	QueryRule interface {
		EvalQuery(context.Context, ent.Query) error
	}
)

// EvalQuery evaluates a query against a query policy.
func (policy QueryPolicy) EvalQuery(ctx context.Context, q ent.Query) error {
	if decision, ok := decisionFromContext(ctx); ok {
		return decision
	}
	for _, rule := range policy {
		switch decision := rule.EvalQuery(ctx, q); {
		case decision == nil || errors.Is(decision, Skip):
		case errors.Is(decision, Allow):
			return nil
		default:
			return decision
		}
	}
	return nil
}

// QueryRuleFunc type is an adapter to allow the use of
// ordinary functions as query rules.
type QueryRuleFunc func(context.Context, ent.Query) error

// Eval returns f(ctx, q).
func (f QueryRuleFunc) EvalQuery(ctx context.Context, q ent.Query) error {
	return f(ctx, q)
}

type (
	// MutationPolicy combines multiple mutation rules into a single policy.
	MutationPolicy []MutationRule

	// MutationRule defines the interface deciding whether a
	// mutation is allowed and optionally modify it.
	MutationRule interface {
		EvalMutation(context.Context, ent.Mutation) error
	}
)

// EvalMutation evaluates a mutation against a mutation policy.
func (policy MutationPolicy) EvalMutation(ctx context.Context, m ent.Mutation) error {
	if decision, ok := decisionFromContext(ctx); ok {
		return decision
	}
	for _, rule := range policy {
		switch decision := rule.EvalMutation(ctx, m); {
		case decision == nil || errors.Is(decision, Skip):
		case errors.Is(decision, Allow):
			return nil
		default:
			return decision
		}
	}
	return nil
}

// MutationRuleFunc type is an adapter to allow the use of
// ordinary functions as mutation rules.
type MutationRuleFunc func(context.Context, ent.Mutation) error

// EvalMutation returns f(ctx, m).
func (f MutationRuleFunc) EvalMutation(ctx context.Context, m ent.Mutation) error {
	return f(ctx, m)
}

// Policy groups query and mutation policies.
type Policy struct {
	Query    QueryPolicy
	Mutation MutationPolicy
}

// EvalQuery forwards evaluation to query policy.
func (policy Policy) EvalQuery(ctx context.Context, q ent.Query) error {
	return policy.Query.EvalQuery(ctx, q)
}

// EvalMutation forwards evaluation to mutation policy.
func (policy Policy) EvalMutation(ctx context.Context, m ent.Mutation) error {
	return policy.Mutation.EvalMutation(ctx, m)
}

// QueryMutationRule is the interface that groups query and mutation rules.
type QueryMutationRule interface {
	QueryRule
	MutationRule
}

// AlwaysAllowRule returns a rule that returns an allow decision.
func AlwaysAllowRule() QueryMutationRule {
	return fixedDecision{Allow}
}

// AlwaysDenyRule returns a rule that returns a deny decision.
func AlwaysDenyRule() QueryMutationRule {
	return fixedDecision{Deny}
}

type fixedDecision struct {
	decision error
}

func (f fixedDecision) EvalQuery(context.Context, ent.Query) error {
	return f.decision
}

func (f fixedDecision) EvalMutation(context.Context, ent.Mutation) error {
	return f.decision
}

type contextDecision struct {
	eval func(context.Context) error
}

// ContextQueryMutationRule creates a query/mutation rule from a context eval func.
func ContextQueryMutationRule(eval func(context.Context) error) QueryMutationRule {
	return contextDecision{eval}
}

func (c contextDecision) EvalQuery(ctx context.Context, _ ent.Query) error {
	return c.eval(ctx)
}

func (c contextDecision) EvalMutation(ctx context.Context, _ ent.Mutation) error {
	return c.eval(ctx)
}

// OnMutationOperation evaluates the given rule only on a given mutation operation.
func OnMutationOperation(rule MutationRule, op ent.Op) MutationRule {
	return MutationRuleFunc(func(ctx context.Context, m ent.Mutation) error {
		if m.Op().Is(op) {
			return rule.EvalMutation(ctx, m)
		}
		return Skip
	})
}

// DenyMutationOperationRule returns a rule denying specified mutation operation.
func DenyMutationOperationRule(op ent.Op) MutationRule {
	rule := MutationRuleFunc(func(_ context.Context, m ent.Mutation) error {
		return Denyf("ent/privacy: operation %s is not allowed", m.Op())
	})
	return OnMutationOperation(rule, op)
}

// The AmountpaidQueryRuleFunc type is an adapter to allow the use of ordinary
// functions as a query rule.
type AmountpaidQueryRuleFunc func(context.Context, *ent.AmountpaidQuery) error

// EvalQuery return f(ctx, q).
func (f AmountpaidQueryRuleFunc) EvalQuery(ctx context.Context, q ent.Query) error {
	if q, ok := q.(*ent.AmountpaidQuery); ok {
		return f(ctx, q)
	}
	return Denyf("ent/privacy: unexpected query type %T, expect *ent.AmountpaidQuery", q)
}

// The AmountpaidMutationRuleFunc type is an adapter to allow the use of ordinary
// functions as a mutation rule.
type AmountpaidMutationRuleFunc func(context.Context, *ent.AmountpaidMutation) error

// EvalMutation calls f(ctx, m).
func (f AmountpaidMutationRuleFunc) EvalMutation(ctx context.Context, m ent.Mutation) error {
	if m, ok := m.(*ent.AmountpaidMutation); ok {
		return f(ctx, m)
	}
	return Denyf("ent/privacy: unexpected mutation type %T, expect *ent.AmountpaidMutation", m)
}

// The BankQueryRuleFunc type is an adapter to allow the use of ordinary
// functions as a query rule.
type BankQueryRuleFunc func(context.Context, *ent.BankQuery) error

// EvalQuery return f(ctx, q).
func (f BankQueryRuleFunc) EvalQuery(ctx context.Context, q ent.Query) error {
	if q, ok := q.(*ent.BankQuery); ok {
		return f(ctx, q)
	}
	return Denyf("ent/privacy: unexpected query type %T, expect *ent.BankQuery", q)
}

// The BankMutationRuleFunc type is an adapter to allow the use of ordinary
// functions as a mutation rule.
type BankMutationRuleFunc func(context.Context, *ent.BankMutation) error

// EvalMutation calls f(ctx, m).
func (f BankMutationRuleFunc) EvalMutation(ctx context.Context, m ent.Mutation) error {
	if m, ok := m.(*ent.BankMutation); ok {
		return f(ctx, m)
	}
	return Denyf("ent/privacy: unexpected mutation type %T, expect *ent.BankMutation", m)
}

// The CategoryQueryRuleFunc type is an adapter to allow the use of ordinary
// functions as a query rule.
type CategoryQueryRuleFunc func(context.Context, *ent.CategoryQuery) error

// EvalQuery return f(ctx, q).
func (f CategoryQueryRuleFunc) EvalQuery(ctx context.Context, q ent.Query) error {
	if q, ok := q.(*ent.CategoryQuery); ok {
		return f(ctx, q)
	}
	return Denyf("ent/privacy: unexpected query type %T, expect *ent.CategoryQuery", q)
}

// The CategoryMutationRuleFunc type is an adapter to allow the use of ordinary
// functions as a mutation rule.
type CategoryMutationRuleFunc func(context.Context, *ent.CategoryMutation) error

// EvalMutation calls f(ctx, m).
func (f CategoryMutationRuleFunc) EvalMutation(ctx context.Context, m ent.Mutation) error {
	if m, ok := m.(*ent.CategoryMutation); ok {
		return f(ctx, m)
	}
	return Denyf("ent/privacy: unexpected mutation type %T, expect *ent.CategoryMutation", m)
}

// The GenderQueryRuleFunc type is an adapter to allow the use of ordinary
// functions as a query rule.
type GenderQueryRuleFunc func(context.Context, *ent.GenderQuery) error

// EvalQuery return f(ctx, q).
func (f GenderQueryRuleFunc) EvalQuery(ctx context.Context, q ent.Query) error {
	if q, ok := q.(*ent.GenderQuery); ok {
		return f(ctx, q)
	}
	return Denyf("ent/privacy: unexpected query type %T, expect *ent.GenderQuery", q)
}

// The GenderMutationRuleFunc type is an adapter to allow the use of ordinary
// functions as a mutation rule.
type GenderMutationRuleFunc func(context.Context, *ent.GenderMutation) error

// EvalMutation calls f(ctx, m).
func (f GenderMutationRuleFunc) EvalMutation(ctx context.Context, m ent.Mutation) error {
	if m, ok := m.(*ent.GenderMutation); ok {
		return f(ctx, m)
	}
	return Denyf("ent/privacy: unexpected mutation type %T, expect *ent.GenderMutation", m)
}

// The GroupOfAgeQueryRuleFunc type is an adapter to allow the use of ordinary
// functions as a query rule.
type GroupOfAgeQueryRuleFunc func(context.Context, *ent.GroupOfAgeQuery) error

// EvalQuery return f(ctx, q).
func (f GroupOfAgeQueryRuleFunc) EvalQuery(ctx context.Context, q ent.Query) error {
	if q, ok := q.(*ent.GroupOfAgeQuery); ok {
		return f(ctx, q)
	}
	return Denyf("ent/privacy: unexpected query type %T, expect *ent.GroupOfAgeQuery", q)
}

// The GroupOfAgeMutationRuleFunc type is an adapter to allow the use of ordinary
// functions as a mutation rule.
type GroupOfAgeMutationRuleFunc func(context.Context, *ent.GroupOfAgeMutation) error

// EvalMutation calls f(ctx, m).
func (f GroupOfAgeMutationRuleFunc) EvalMutation(ctx context.Context, m ent.Mutation) error {
	if m, ok := m.(*ent.GroupOfAgeMutation); ok {
		return f(ctx, m)
	}
	return Denyf("ent/privacy: unexpected mutation type %T, expect *ent.GroupOfAgeMutation", m)
}

// The HospitalQueryRuleFunc type is an adapter to allow the use of ordinary
// functions as a query rule.
type HospitalQueryRuleFunc func(context.Context, *ent.HospitalQuery) error

// EvalQuery return f(ctx, q).
func (f HospitalQueryRuleFunc) EvalQuery(ctx context.Context, q ent.Query) error {
	if q, ok := q.(*ent.HospitalQuery); ok {
		return f(ctx, q)
	}
	return Denyf("ent/privacy: unexpected query type %T, expect *ent.HospitalQuery", q)
}

// The HospitalMutationRuleFunc type is an adapter to allow the use of ordinary
// functions as a mutation rule.
type HospitalMutationRuleFunc func(context.Context, *ent.HospitalMutation) error

// EvalMutation calls f(ctx, m).
func (f HospitalMutationRuleFunc) EvalMutation(ctx context.Context, m ent.Mutation) error {
	if m, ok := m.(*ent.HospitalMutation); ok {
		return f(ctx, m)
	}
	return Denyf("ent/privacy: unexpected mutation type %T, expect *ent.HospitalMutation", m)
}

// The InquiryQueryRuleFunc type is an adapter to allow the use of ordinary
// functions as a query rule.
type InquiryQueryRuleFunc func(context.Context, *ent.InquiryQuery) error

// EvalQuery return f(ctx, q).
func (f InquiryQueryRuleFunc) EvalQuery(ctx context.Context, q ent.Query) error {
	if q, ok := q.(*ent.InquiryQuery); ok {
		return f(ctx, q)
	}
	return Denyf("ent/privacy: unexpected query type %T, expect *ent.InquiryQuery", q)
}

// The InquiryMutationRuleFunc type is an adapter to allow the use of ordinary
// functions as a mutation rule.
type InquiryMutationRuleFunc func(context.Context, *ent.InquiryMutation) error

// EvalMutation calls f(ctx, m).
func (f InquiryMutationRuleFunc) EvalMutation(ctx context.Context, m ent.Mutation) error {
	if m, ok := m.(*ent.InquiryMutation); ok {
		return f(ctx, m)
	}
	return Denyf("ent/privacy: unexpected mutation type %T, expect *ent.InquiryMutation", m)
}

// The InsuranceQueryRuleFunc type is an adapter to allow the use of ordinary
// functions as a query rule.
type InsuranceQueryRuleFunc func(context.Context, *ent.InsuranceQuery) error

// EvalQuery return f(ctx, q).
func (f InsuranceQueryRuleFunc) EvalQuery(ctx context.Context, q ent.Query) error {
	if q, ok := q.(*ent.InsuranceQuery); ok {
		return f(ctx, q)
	}
	return Denyf("ent/privacy: unexpected query type %T, expect *ent.InsuranceQuery", q)
}

// The InsuranceMutationRuleFunc type is an adapter to allow the use of ordinary
// functions as a mutation rule.
type InsuranceMutationRuleFunc func(context.Context, *ent.InsuranceMutation) error

// EvalMutation calls f(ctx, m).
func (f InsuranceMutationRuleFunc) EvalMutation(ctx context.Context, m ent.Mutation) error {
	if m, ok := m.(*ent.InsuranceMutation); ok {
		return f(ctx, m)
	}
	return Denyf("ent/privacy: unexpected mutation type %T, expect *ent.InsuranceMutation", m)
}

// The MemberQueryRuleFunc type is an adapter to allow the use of ordinary
// functions as a query rule.
type MemberQueryRuleFunc func(context.Context, *ent.MemberQuery) error

// EvalQuery return f(ctx, q).
func (f MemberQueryRuleFunc) EvalQuery(ctx context.Context, q ent.Query) error {
	if q, ok := q.(*ent.MemberQuery); ok {
		return f(ctx, q)
	}
	return Denyf("ent/privacy: unexpected query type %T, expect *ent.MemberQuery", q)
}

// The MemberMutationRuleFunc type is an adapter to allow the use of ordinary
// functions as a mutation rule.
type MemberMutationRuleFunc func(context.Context, *ent.MemberMutation) error

// EvalMutation calls f(ctx, m).
func (f MemberMutationRuleFunc) EvalMutation(ctx context.Context, m ent.Mutation) error {
	if m, ok := m.(*ent.MemberMutation); ok {
		return f(ctx, m)
	}
	return Denyf("ent/privacy: unexpected mutation type %T, expect *ent.MemberMutation", m)
}

// The MoneyTransferQueryRuleFunc type is an adapter to allow the use of ordinary
// functions as a query rule.
type MoneyTransferQueryRuleFunc func(context.Context, *ent.MoneyTransferQuery) error

// EvalQuery return f(ctx, q).
func (f MoneyTransferQueryRuleFunc) EvalQuery(ctx context.Context, q ent.Query) error {
	if q, ok := q.(*ent.MoneyTransferQuery); ok {
		return f(ctx, q)
	}
	return Denyf("ent/privacy: unexpected query type %T, expect *ent.MoneyTransferQuery", q)
}

// The MoneyTransferMutationRuleFunc type is an adapter to allow the use of ordinary
// functions as a mutation rule.
type MoneyTransferMutationRuleFunc func(context.Context, *ent.MoneyTransferMutation) error

// EvalMutation calls f(ctx, m).
func (f MoneyTransferMutationRuleFunc) EvalMutation(ctx context.Context, m ent.Mutation) error {
	if m, ok := m.(*ent.MoneyTransferMutation); ok {
		return f(ctx, m)
	}
	return Denyf("ent/privacy: unexpected mutation type %T, expect *ent.MoneyTransferMutation", m)
}

// The OfficerQueryRuleFunc type is an adapter to allow the use of ordinary
// functions as a query rule.
type OfficerQueryRuleFunc func(context.Context, *ent.OfficerQuery) error

// EvalQuery return f(ctx, q).
func (f OfficerQueryRuleFunc) EvalQuery(ctx context.Context, q ent.Query) error {
	if q, ok := q.(*ent.OfficerQuery); ok {
		return f(ctx, q)
	}
	return Denyf("ent/privacy: unexpected query type %T, expect *ent.OfficerQuery", q)
}

// The OfficerMutationRuleFunc type is an adapter to allow the use of ordinary
// functions as a mutation rule.
type OfficerMutationRuleFunc func(context.Context, *ent.OfficerMutation) error

// EvalMutation calls f(ctx, m).
func (f OfficerMutationRuleFunc) EvalMutation(ctx context.Context, m ent.Mutation) error {
	if m, ok := m.(*ent.OfficerMutation); ok {
		return f(ctx, m)
	}
	return Denyf("ent/privacy: unexpected mutation type %T, expect *ent.OfficerMutation", m)
}

// The PaybackQueryRuleFunc type is an adapter to allow the use of ordinary
// functions as a query rule.
type PaybackQueryRuleFunc func(context.Context, *ent.PaybackQuery) error

// EvalQuery return f(ctx, q).
func (f PaybackQueryRuleFunc) EvalQuery(ctx context.Context, q ent.Query) error {
	if q, ok := q.(*ent.PaybackQuery); ok {
		return f(ctx, q)
	}
	return Denyf("ent/privacy: unexpected query type %T, expect *ent.PaybackQuery", q)
}

// The PaybackMutationRuleFunc type is an adapter to allow the use of ordinary
// functions as a mutation rule.
type PaybackMutationRuleFunc func(context.Context, *ent.PaybackMutation) error

// EvalMutation calls f(ctx, m).
func (f PaybackMutationRuleFunc) EvalMutation(ctx context.Context, m ent.Mutation) error {
	if m, ok := m.(*ent.PaybackMutation); ok {
		return f(ctx, m)
	}
	return Denyf("ent/privacy: unexpected mutation type %T, expect *ent.PaybackMutation", m)
}

// The PaymentQueryRuleFunc type is an adapter to allow the use of ordinary
// functions as a query rule.
type PaymentQueryRuleFunc func(context.Context, *ent.PaymentQuery) error

// EvalQuery return f(ctx, q).
func (f PaymentQueryRuleFunc) EvalQuery(ctx context.Context, q ent.Query) error {
	if q, ok := q.(*ent.PaymentQuery); ok {
		return f(ctx, q)
	}
	return Denyf("ent/privacy: unexpected query type %T, expect *ent.PaymentQuery", q)
}

// The PaymentMutationRuleFunc type is an adapter to allow the use of ordinary
// functions as a mutation rule.
type PaymentMutationRuleFunc func(context.Context, *ent.PaymentMutation) error

// EvalMutation calls f(ctx, m).
func (f PaymentMutationRuleFunc) EvalMutation(ctx context.Context, m ent.Mutation) error {
	if m, ok := m.(*ent.PaymentMutation); ok {
		return f(ctx, m)
	}
	return Denyf("ent/privacy: unexpected mutation type %T, expect *ent.PaymentMutation", m)
}

// The ProductQueryRuleFunc type is an adapter to allow the use of ordinary
// functions as a query rule.
type ProductQueryRuleFunc func(context.Context, *ent.ProductQuery) error

// EvalQuery return f(ctx, q).
func (f ProductQueryRuleFunc) EvalQuery(ctx context.Context, q ent.Query) error {
	if q, ok := q.(*ent.ProductQuery); ok {
		return f(ctx, q)
	}
	return Denyf("ent/privacy: unexpected query type %T, expect *ent.ProductQuery", q)
}

// The ProductMutationRuleFunc type is an adapter to allow the use of ordinary
// functions as a mutation rule.
type ProductMutationRuleFunc func(context.Context, *ent.ProductMutation) error

// EvalMutation calls f(ctx, m).
func (f ProductMutationRuleFunc) EvalMutation(ctx context.Context, m ent.Mutation) error {
	if m, ok := m.(*ent.ProductMutation); ok {
		return f(ctx, m)
	}
	return Denyf("ent/privacy: unexpected mutation type %T, expect *ent.ProductMutation", m)
}

// The RecordinsuranceQueryRuleFunc type is an adapter to allow the use of ordinary
// functions as a query rule.
type RecordinsuranceQueryRuleFunc func(context.Context, *ent.RecordinsuranceQuery) error

// EvalQuery return f(ctx, q).
func (f RecordinsuranceQueryRuleFunc) EvalQuery(ctx context.Context, q ent.Query) error {
	if q, ok := q.(*ent.RecordinsuranceQuery); ok {
		return f(ctx, q)
	}
	return Denyf("ent/privacy: unexpected query type %T, expect *ent.RecordinsuranceQuery", q)
}

// The RecordinsuranceMutationRuleFunc type is an adapter to allow the use of ordinary
// functions as a mutation rule.
type RecordinsuranceMutationRuleFunc func(context.Context, *ent.RecordinsuranceMutation) error

// EvalMutation calls f(ctx, m).
func (f RecordinsuranceMutationRuleFunc) EvalMutation(ctx context.Context, m ent.Mutation) error {
	if m, ok := m.(*ent.RecordinsuranceMutation); ok {
		return f(ctx, m)
	}
	return Denyf("ent/privacy: unexpected mutation type %T, expect *ent.RecordinsuranceMutation", m)
}
