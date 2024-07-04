// Code generated by ent, DO NOT EDIT.

package privacy

import (
	"context"
	"fmt"
	"waterfall-backend/internal/modules/stores/db/ent"

	"entgo.io/ent/privacy"
)

var (
	// Allow may be returned by rules to indicate that the policy
	// evaluation should terminate with allow decision.
	Allow = privacy.Allow

	// Deny may be returned by rules to indicate that the policy
	// evaluation should terminate with deny decision.
	Deny = privacy.Deny

	// Skip may be returned by rules to indicate that the policy
	// evaluation should continue to the next rule.
	Skip = privacy.Skip
)

// Allowf returns an formatted wrapped Allow decision.
func Allowf(format string, a ...any) error {
	return fmt.Errorf(format+": %w", append(a, Allow)...)
}

// Denyf returns an formatted wrapped Deny decision.
func Denyf(format string, a ...any) error {
	return fmt.Errorf(format+": %w", append(a, Deny)...)
}

// Skipf returns an formatted wrapped Skip decision.
func Skipf(format string, a ...any) error {
	return fmt.Errorf(format+": %w", append(a, Skip)...)
}

// DecisionContext creates a new context from the given parent context with
// a policy decision attach to it.
func DecisionContext(parent context.Context, decision error) context.Context {
	return privacy.DecisionContext(parent, decision)
}

// DecisionFromContext retrieves the policy decision from the context.
func DecisionFromContext(ctx context.Context) (error, bool) {
	return privacy.DecisionFromContext(ctx)
}

type (
	// Policy groups query and mutation policies.
	Policy = privacy.Policy

	// QueryRule defines the interface deciding whether a
	// query is allowed and optionally modify it.
	QueryRule = privacy.QueryRule
	// QueryPolicy combines multiple query rules into a single policy.
	QueryPolicy = privacy.QueryPolicy

	// MutationRule defines the interface which decides whether a
	// mutation is allowed and optionally modifies it.
	MutationRule = privacy.MutationRule
	// MutationPolicy combines multiple mutation rules into a single policy.
	MutationPolicy = privacy.MutationPolicy
)

// QueryRuleFunc type is an adapter to allow the use of
// ordinary functions as query rules.
type QueryRuleFunc func(context.Context, ent.Query) error

// Eval returns f(ctx, q).
func (f QueryRuleFunc) EvalQuery(ctx context.Context, q ent.Query) error {
	return f(ctx, q)
}

// MutationRuleFunc type is an adapter which allows the use of
// ordinary functions as mutation rules.
type MutationRuleFunc func(context.Context, ent.Mutation) error

// EvalMutation returns f(ctx, m).
func (f MutationRuleFunc) EvalMutation(ctx context.Context, m ent.Mutation) error {
	return f(ctx, m)
}

// QueryMutationRule is an interface which groups query and mutation rules.
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

// The ChapterQueryRuleFunc type is an adapter to allow the use of ordinary
// functions as a query rule.
type ChapterQueryRuleFunc func(context.Context, *ent.ChapterQuery) error

// EvalQuery return f(ctx, q).
func (f ChapterQueryRuleFunc) EvalQuery(ctx context.Context, q ent.Query) error {
	if q, ok := q.(*ent.ChapterQuery); ok {
		return f(ctx, q)
	}
	return Denyf("ent/privacy: unexpected query type %T, expect *ent.ChapterQuery", q)
}

// The ChapterMutationRuleFunc type is an adapter to allow the use of ordinary
// functions as a mutation rule.
type ChapterMutationRuleFunc func(context.Context, *ent.ChapterMutation) error

// EvalMutation calls f(ctx, m).
func (f ChapterMutationRuleFunc) EvalMutation(ctx context.Context, m ent.Mutation) error {
	if m, ok := m.(*ent.ChapterMutation); ok {
		return f(ctx, m)
	}
	return Denyf("ent/privacy: unexpected mutation type %T, expect *ent.ChapterMutation", m)
}

// The ChapterTextQueryRuleFunc type is an adapter to allow the use of ordinary
// functions as a query rule.
type ChapterTextQueryRuleFunc func(context.Context, *ent.ChapterTextQuery) error

// EvalQuery return f(ctx, q).
func (f ChapterTextQueryRuleFunc) EvalQuery(ctx context.Context, q ent.Query) error {
	if q, ok := q.(*ent.ChapterTextQuery); ok {
		return f(ctx, q)
	}
	return Denyf("ent/privacy: unexpected query type %T, expect *ent.ChapterTextQuery", q)
}

// The ChapterTextMutationRuleFunc type is an adapter to allow the use of ordinary
// functions as a mutation rule.
type ChapterTextMutationRuleFunc func(context.Context, *ent.ChapterTextMutation) error

// EvalMutation calls f(ctx, m).
func (f ChapterTextMutationRuleFunc) EvalMutation(ctx context.Context, m ent.Mutation) error {
	if m, ok := m.(*ent.ChapterTextMutation); ok {
		return f(ctx, m)
	}
	return Denyf("ent/privacy: unexpected mutation type %T, expect *ent.ChapterTextMutation", m)
}

// The CommentQueryRuleFunc type is an adapter to allow the use of ordinary
// functions as a query rule.
type CommentQueryRuleFunc func(context.Context, *ent.CommentQuery) error

// EvalQuery return f(ctx, q).
func (f CommentQueryRuleFunc) EvalQuery(ctx context.Context, q ent.Query) error {
	if q, ok := q.(*ent.CommentQuery); ok {
		return f(ctx, q)
	}
	return Denyf("ent/privacy: unexpected query type %T, expect *ent.CommentQuery", q)
}

// The CommentMutationRuleFunc type is an adapter to allow the use of ordinary
// functions as a mutation rule.
type CommentMutationRuleFunc func(context.Context, *ent.CommentMutation) error

// EvalMutation calls f(ctx, m).
func (f CommentMutationRuleFunc) EvalMutation(ctx context.Context, m ent.Mutation) error {
	if m, ok := m.(*ent.CommentMutation); ok {
		return f(ctx, m)
	}
	return Denyf("ent/privacy: unexpected mutation type %T, expect *ent.CommentMutation", m)
}

// The FileQueryRuleFunc type is an adapter to allow the use of ordinary
// functions as a query rule.
type FileQueryRuleFunc func(context.Context, *ent.FileQuery) error

// EvalQuery return f(ctx, q).
func (f FileQueryRuleFunc) EvalQuery(ctx context.Context, q ent.Query) error {
	if q, ok := q.(*ent.FileQuery); ok {
		return f(ctx, q)
	}
	return Denyf("ent/privacy: unexpected query type %T, expect *ent.FileQuery", q)
}

// The FileMutationRuleFunc type is an adapter to allow the use of ordinary
// functions as a mutation rule.
type FileMutationRuleFunc func(context.Context, *ent.FileMutation) error

// EvalMutation calls f(ctx, m).
func (f FileMutationRuleFunc) EvalMutation(ctx context.Context, m ent.Mutation) error {
	if m, ok := m.(*ent.FileMutation); ok {
		return f(ctx, m)
	}
	return Denyf("ent/privacy: unexpected mutation type %T, expect *ent.FileMutation", m)
}

// The MigrationsQueryRuleFunc type is an adapter to allow the use of ordinary
// functions as a query rule.
type MigrationsQueryRuleFunc func(context.Context, *ent.MigrationsQuery) error

// EvalQuery return f(ctx, q).
func (f MigrationsQueryRuleFunc) EvalQuery(ctx context.Context, q ent.Query) error {
	if q, ok := q.(*ent.MigrationsQuery); ok {
		return f(ctx, q)
	}
	return Denyf("ent/privacy: unexpected query type %T, expect *ent.MigrationsQuery", q)
}

// The MigrationsMutationRuleFunc type is an adapter to allow the use of ordinary
// functions as a mutation rule.
type MigrationsMutationRuleFunc func(context.Context, *ent.MigrationsMutation) error

// EvalMutation calls f(ctx, m).
func (f MigrationsMutationRuleFunc) EvalMutation(ctx context.Context, m ent.Mutation) error {
	if m, ok := m.(*ent.MigrationsMutation); ok {
		return f(ctx, m)
	}
	return Denyf("ent/privacy: unexpected mutation type %T, expect *ent.MigrationsMutation", m)
}

// The PartQueryRuleFunc type is an adapter to allow the use of ordinary
// functions as a query rule.
type PartQueryRuleFunc func(context.Context, *ent.PartQuery) error

// EvalQuery return f(ctx, q).
func (f PartQueryRuleFunc) EvalQuery(ctx context.Context, q ent.Query) error {
	if q, ok := q.(*ent.PartQuery); ok {
		return f(ctx, q)
	}
	return Denyf("ent/privacy: unexpected query type %T, expect *ent.PartQuery", q)
}

// The PartMutationRuleFunc type is an adapter to allow the use of ordinary
// functions as a mutation rule.
type PartMutationRuleFunc func(context.Context, *ent.PartMutation) error

// EvalMutation calls f(ctx, m).
func (f PartMutationRuleFunc) EvalMutation(ctx context.Context, m ent.Mutation) error {
	if m, ok := m.(*ent.PartMutation); ok {
		return f(ctx, m)
	}
	return Denyf("ent/privacy: unexpected mutation type %T, expect *ent.PartMutation", m)
}

// The PermissionQueryRuleFunc type is an adapter to allow the use of ordinary
// functions as a query rule.
type PermissionQueryRuleFunc func(context.Context, *ent.PermissionQuery) error

// EvalQuery return f(ctx, q).
func (f PermissionQueryRuleFunc) EvalQuery(ctx context.Context, q ent.Query) error {
	if q, ok := q.(*ent.PermissionQuery); ok {
		return f(ctx, q)
	}
	return Denyf("ent/privacy: unexpected query type %T, expect *ent.PermissionQuery", q)
}

// The PermissionMutationRuleFunc type is an adapter to allow the use of ordinary
// functions as a mutation rule.
type PermissionMutationRuleFunc func(context.Context, *ent.PermissionMutation) error

// EvalMutation calls f(ctx, m).
func (f PermissionMutationRuleFunc) EvalMutation(ctx context.Context, m ent.Mutation) error {
	if m, ok := m.(*ent.PermissionMutation); ok {
		return f(ctx, m)
	}
	return Denyf("ent/privacy: unexpected mutation type %T, expect *ent.PermissionMutation", m)
}

// The ReleaseQueryRuleFunc type is an adapter to allow the use of ordinary
// functions as a query rule.
type ReleaseQueryRuleFunc func(context.Context, *ent.ReleaseQuery) error

// EvalQuery return f(ctx, q).
func (f ReleaseQueryRuleFunc) EvalQuery(ctx context.Context, q ent.Query) error {
	if q, ok := q.(*ent.ReleaseQuery); ok {
		return f(ctx, q)
	}
	return Denyf("ent/privacy: unexpected query type %T, expect *ent.ReleaseQuery", q)
}

// The ReleaseMutationRuleFunc type is an adapter to allow the use of ordinary
// functions as a mutation rule.
type ReleaseMutationRuleFunc func(context.Context, *ent.ReleaseMutation) error

// EvalMutation calls f(ctx, m).
func (f ReleaseMutationRuleFunc) EvalMutation(ctx context.Context, m ent.Mutation) error {
	if m, ok := m.(*ent.ReleaseMutation); ok {
		return f(ctx, m)
	}
	return Denyf("ent/privacy: unexpected mutation type %T, expect *ent.ReleaseMutation", m)
}

// The RoleQueryRuleFunc type is an adapter to allow the use of ordinary
// functions as a query rule.
type RoleQueryRuleFunc func(context.Context, *ent.RoleQuery) error

// EvalQuery return f(ctx, q).
func (f RoleQueryRuleFunc) EvalQuery(ctx context.Context, q ent.Query) error {
	if q, ok := q.(*ent.RoleQuery); ok {
		return f(ctx, q)
	}
	return Denyf("ent/privacy: unexpected query type %T, expect *ent.RoleQuery", q)
}

// The RoleMutationRuleFunc type is an adapter to allow the use of ordinary
// functions as a mutation rule.
type RoleMutationRuleFunc func(context.Context, *ent.RoleMutation) error

// EvalMutation calls f(ctx, m).
func (f RoleMutationRuleFunc) EvalMutation(ctx context.Context, m ent.Mutation) error {
	if m, ok := m.(*ent.RoleMutation); ok {
		return f(ctx, m)
	}
	return Denyf("ent/privacy: unexpected mutation type %T, expect *ent.RoleMutation", m)
}

// The SeedMigrationsQueryRuleFunc type is an adapter to allow the use of ordinary
// functions as a query rule.
type SeedMigrationsQueryRuleFunc func(context.Context, *ent.SeedMigrationsQuery) error

// EvalQuery return f(ctx, q).
func (f SeedMigrationsQueryRuleFunc) EvalQuery(ctx context.Context, q ent.Query) error {
	if q, ok := q.(*ent.SeedMigrationsQuery); ok {
		return f(ctx, q)
	}
	return Denyf("ent/privacy: unexpected query type %T, expect *ent.SeedMigrationsQuery", q)
}

// The SeedMigrationsMutationRuleFunc type is an adapter to allow the use of ordinary
// functions as a mutation rule.
type SeedMigrationsMutationRuleFunc func(context.Context, *ent.SeedMigrationsMutation) error

// EvalMutation calls f(ctx, m).
func (f SeedMigrationsMutationRuleFunc) EvalMutation(ctx context.Context, m ent.Mutation) error {
	if m, ok := m.(*ent.SeedMigrationsMutation); ok {
		return f(ctx, m)
	}
	return Denyf("ent/privacy: unexpected mutation type %T, expect *ent.SeedMigrationsMutation", m)
}

// The UserQueryRuleFunc type is an adapter to allow the use of ordinary
// functions as a query rule.
type UserQueryRuleFunc func(context.Context, *ent.UserQuery) error

// EvalQuery return f(ctx, q).
func (f UserQueryRuleFunc) EvalQuery(ctx context.Context, q ent.Query) error {
	if q, ok := q.(*ent.UserQuery); ok {
		return f(ctx, q)
	}
	return Denyf("ent/privacy: unexpected query type %T, expect *ent.UserQuery", q)
}

// The UserMutationRuleFunc type is an adapter to allow the use of ordinary
// functions as a mutation rule.
type UserMutationRuleFunc func(context.Context, *ent.UserMutation) error

// EvalMutation calls f(ctx, m).
func (f UserMutationRuleFunc) EvalMutation(ctx context.Context, m ent.Mutation) error {
	if m, ok := m.(*ent.UserMutation); ok {
		return f(ctx, m)
	}
	return Denyf("ent/privacy: unexpected mutation type %T, expect *ent.UserMutation", m)
}
