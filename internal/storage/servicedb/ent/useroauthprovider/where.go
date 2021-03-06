// Code generated by entc, DO NOT EDIT.

package useroauthprovider

import (
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"github.com/sundaytycoon/buttons-api/internal/storage/servicedb/ent/predicate"
)

// ID filters vertices based on their ID field.
func ID(id string) predicate.UserOAuthProvider {
	return predicate.UserOAuthProvider(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldID), id))
	})
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id string) predicate.UserOAuthProvider {
	return predicate.UserOAuthProvider(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldID), id))
	})
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id string) predicate.UserOAuthProvider {
	return predicate.UserOAuthProvider(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldID), id))
	})
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...string) predicate.UserOAuthProvider {
	return predicate.UserOAuthProvider(func(s *sql.Selector) {
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
func IDNotIn(ids ...string) predicate.UserOAuthProvider {
	return predicate.UserOAuthProvider(func(s *sql.Selector) {
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
func IDGT(id string) predicate.UserOAuthProvider {
	return predicate.UserOAuthProvider(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldID), id))
	})
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id string) predicate.UserOAuthProvider {
	return predicate.UserOAuthProvider(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldID), id))
	})
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id string) predicate.UserOAuthProvider {
	return predicate.UserOAuthProvider(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldID), id))
	})
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id string) predicate.UserOAuthProvider {
	return predicate.UserOAuthProvider(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldID), id))
	})
}

// UserID applies equality check predicate on the "user_id" field. It's identical to UserIDEQ.
func UserID(v string) predicate.UserOAuthProvider {
	return predicate.UserOAuthProvider(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldUserID), v))
	})
}

// CreatedAt applies equality check predicate on the "created_at" field. It's identical to CreatedAtEQ.
func CreatedAt(v time.Time) predicate.UserOAuthProvider {
	return predicate.UserOAuthProvider(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldCreatedAt), v))
	})
}

// CreatedBy applies equality check predicate on the "created_by" field. It's identical to CreatedByEQ.
func CreatedBy(v string) predicate.UserOAuthProvider {
	return predicate.UserOAuthProvider(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldCreatedBy), v))
	})
}

// UpdatedAt applies equality check predicate on the "updated_at" field. It's identical to UpdatedAtEQ.
func UpdatedAt(v time.Time) predicate.UserOAuthProvider {
	return predicate.UserOAuthProvider(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldUpdatedAt), v))
	})
}

// UpdatedBy applies equality check predicate on the "updated_by" field. It's identical to UpdatedByEQ.
func UpdatedBy(v string) predicate.UserOAuthProvider {
	return predicate.UserOAuthProvider(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldUpdatedBy), v))
	})
}

// Expiry applies equality check predicate on the "expiry" field. It's identical to ExpiryEQ.
func Expiry(v time.Time) predicate.UserOAuthProvider {
	return predicate.UserOAuthProvider(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldExpiry), v))
	})
}

// AccessToken applies equality check predicate on the "access_token" field. It's identical to AccessTokenEQ.
func AccessToken(v string) predicate.UserOAuthProvider {
	return predicate.UserOAuthProvider(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldAccessToken), v))
	})
}

// RefreshToken applies equality check predicate on the "refresh_token" field. It's identical to RefreshTokenEQ.
func RefreshToken(v string) predicate.UserOAuthProvider {
	return predicate.UserOAuthProvider(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldRefreshToken), v))
	})
}

// UserIDEQ applies the EQ predicate on the "user_id" field.
func UserIDEQ(v string) predicate.UserOAuthProvider {
	return predicate.UserOAuthProvider(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldUserID), v))
	})
}

// UserIDNEQ applies the NEQ predicate on the "user_id" field.
func UserIDNEQ(v string) predicate.UserOAuthProvider {
	return predicate.UserOAuthProvider(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldUserID), v))
	})
}

// UserIDIn applies the In predicate on the "user_id" field.
func UserIDIn(vs ...string) predicate.UserOAuthProvider {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.UserOAuthProvider(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldUserID), v...))
	})
}

// UserIDNotIn applies the NotIn predicate on the "user_id" field.
func UserIDNotIn(vs ...string) predicate.UserOAuthProvider {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.UserOAuthProvider(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldUserID), v...))
	})
}

// UserIDGT applies the GT predicate on the "user_id" field.
func UserIDGT(v string) predicate.UserOAuthProvider {
	return predicate.UserOAuthProvider(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldUserID), v))
	})
}

// UserIDGTE applies the GTE predicate on the "user_id" field.
func UserIDGTE(v string) predicate.UserOAuthProvider {
	return predicate.UserOAuthProvider(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldUserID), v))
	})
}

// UserIDLT applies the LT predicate on the "user_id" field.
func UserIDLT(v string) predicate.UserOAuthProvider {
	return predicate.UserOAuthProvider(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldUserID), v))
	})
}

// UserIDLTE applies the LTE predicate on the "user_id" field.
func UserIDLTE(v string) predicate.UserOAuthProvider {
	return predicate.UserOAuthProvider(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldUserID), v))
	})
}

// UserIDContains applies the Contains predicate on the "user_id" field.
func UserIDContains(v string) predicate.UserOAuthProvider {
	return predicate.UserOAuthProvider(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldUserID), v))
	})
}

// UserIDHasPrefix applies the HasPrefix predicate on the "user_id" field.
func UserIDHasPrefix(v string) predicate.UserOAuthProvider {
	return predicate.UserOAuthProvider(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldUserID), v))
	})
}

// UserIDHasSuffix applies the HasSuffix predicate on the "user_id" field.
func UserIDHasSuffix(v string) predicate.UserOAuthProvider {
	return predicate.UserOAuthProvider(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldUserID), v))
	})
}

// UserIDEqualFold applies the EqualFold predicate on the "user_id" field.
func UserIDEqualFold(v string) predicate.UserOAuthProvider {
	return predicate.UserOAuthProvider(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldUserID), v))
	})
}

// UserIDContainsFold applies the ContainsFold predicate on the "user_id" field.
func UserIDContainsFold(v string) predicate.UserOAuthProvider {
	return predicate.UserOAuthProvider(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldUserID), v))
	})
}

// CreatedAtEQ applies the EQ predicate on the "created_at" field.
func CreatedAtEQ(v time.Time) predicate.UserOAuthProvider {
	return predicate.UserOAuthProvider(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldCreatedAt), v))
	})
}

// CreatedAtNEQ applies the NEQ predicate on the "created_at" field.
func CreatedAtNEQ(v time.Time) predicate.UserOAuthProvider {
	return predicate.UserOAuthProvider(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldCreatedAt), v))
	})
}

// CreatedAtIn applies the In predicate on the "created_at" field.
func CreatedAtIn(vs ...time.Time) predicate.UserOAuthProvider {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.UserOAuthProvider(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldCreatedAt), v...))
	})
}

// CreatedAtNotIn applies the NotIn predicate on the "created_at" field.
func CreatedAtNotIn(vs ...time.Time) predicate.UserOAuthProvider {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.UserOAuthProvider(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldCreatedAt), v...))
	})
}

// CreatedAtGT applies the GT predicate on the "created_at" field.
func CreatedAtGT(v time.Time) predicate.UserOAuthProvider {
	return predicate.UserOAuthProvider(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldCreatedAt), v))
	})
}

// CreatedAtGTE applies the GTE predicate on the "created_at" field.
func CreatedAtGTE(v time.Time) predicate.UserOAuthProvider {
	return predicate.UserOAuthProvider(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldCreatedAt), v))
	})
}

// CreatedAtLT applies the LT predicate on the "created_at" field.
func CreatedAtLT(v time.Time) predicate.UserOAuthProvider {
	return predicate.UserOAuthProvider(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldCreatedAt), v))
	})
}

// CreatedAtLTE applies the LTE predicate on the "created_at" field.
func CreatedAtLTE(v time.Time) predicate.UserOAuthProvider {
	return predicate.UserOAuthProvider(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldCreatedAt), v))
	})
}

// CreatedByEQ applies the EQ predicate on the "created_by" field.
func CreatedByEQ(v string) predicate.UserOAuthProvider {
	return predicate.UserOAuthProvider(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldCreatedBy), v))
	})
}

// CreatedByNEQ applies the NEQ predicate on the "created_by" field.
func CreatedByNEQ(v string) predicate.UserOAuthProvider {
	return predicate.UserOAuthProvider(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldCreatedBy), v))
	})
}

// CreatedByIn applies the In predicate on the "created_by" field.
func CreatedByIn(vs ...string) predicate.UserOAuthProvider {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.UserOAuthProvider(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldCreatedBy), v...))
	})
}

// CreatedByNotIn applies the NotIn predicate on the "created_by" field.
func CreatedByNotIn(vs ...string) predicate.UserOAuthProvider {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.UserOAuthProvider(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldCreatedBy), v...))
	})
}

// CreatedByGT applies the GT predicate on the "created_by" field.
func CreatedByGT(v string) predicate.UserOAuthProvider {
	return predicate.UserOAuthProvider(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldCreatedBy), v))
	})
}

// CreatedByGTE applies the GTE predicate on the "created_by" field.
func CreatedByGTE(v string) predicate.UserOAuthProvider {
	return predicate.UserOAuthProvider(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldCreatedBy), v))
	})
}

// CreatedByLT applies the LT predicate on the "created_by" field.
func CreatedByLT(v string) predicate.UserOAuthProvider {
	return predicate.UserOAuthProvider(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldCreatedBy), v))
	})
}

// CreatedByLTE applies the LTE predicate on the "created_by" field.
func CreatedByLTE(v string) predicate.UserOAuthProvider {
	return predicate.UserOAuthProvider(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldCreatedBy), v))
	})
}

// CreatedByContains applies the Contains predicate on the "created_by" field.
func CreatedByContains(v string) predicate.UserOAuthProvider {
	return predicate.UserOAuthProvider(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldCreatedBy), v))
	})
}

// CreatedByHasPrefix applies the HasPrefix predicate on the "created_by" field.
func CreatedByHasPrefix(v string) predicate.UserOAuthProvider {
	return predicate.UserOAuthProvider(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldCreatedBy), v))
	})
}

// CreatedByHasSuffix applies the HasSuffix predicate on the "created_by" field.
func CreatedByHasSuffix(v string) predicate.UserOAuthProvider {
	return predicate.UserOAuthProvider(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldCreatedBy), v))
	})
}

// CreatedByEqualFold applies the EqualFold predicate on the "created_by" field.
func CreatedByEqualFold(v string) predicate.UserOAuthProvider {
	return predicate.UserOAuthProvider(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldCreatedBy), v))
	})
}

// CreatedByContainsFold applies the ContainsFold predicate on the "created_by" field.
func CreatedByContainsFold(v string) predicate.UserOAuthProvider {
	return predicate.UserOAuthProvider(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldCreatedBy), v))
	})
}

// UpdatedAtEQ applies the EQ predicate on the "updated_at" field.
func UpdatedAtEQ(v time.Time) predicate.UserOAuthProvider {
	return predicate.UserOAuthProvider(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldUpdatedAt), v))
	})
}

// UpdatedAtNEQ applies the NEQ predicate on the "updated_at" field.
func UpdatedAtNEQ(v time.Time) predicate.UserOAuthProvider {
	return predicate.UserOAuthProvider(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldUpdatedAt), v))
	})
}

// UpdatedAtIn applies the In predicate on the "updated_at" field.
func UpdatedAtIn(vs ...time.Time) predicate.UserOAuthProvider {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.UserOAuthProvider(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldUpdatedAt), v...))
	})
}

// UpdatedAtNotIn applies the NotIn predicate on the "updated_at" field.
func UpdatedAtNotIn(vs ...time.Time) predicate.UserOAuthProvider {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.UserOAuthProvider(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldUpdatedAt), v...))
	})
}

// UpdatedAtGT applies the GT predicate on the "updated_at" field.
func UpdatedAtGT(v time.Time) predicate.UserOAuthProvider {
	return predicate.UserOAuthProvider(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldUpdatedAt), v))
	})
}

// UpdatedAtGTE applies the GTE predicate on the "updated_at" field.
func UpdatedAtGTE(v time.Time) predicate.UserOAuthProvider {
	return predicate.UserOAuthProvider(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldUpdatedAt), v))
	})
}

// UpdatedAtLT applies the LT predicate on the "updated_at" field.
func UpdatedAtLT(v time.Time) predicate.UserOAuthProvider {
	return predicate.UserOAuthProvider(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldUpdatedAt), v))
	})
}

// UpdatedAtLTE applies the LTE predicate on the "updated_at" field.
func UpdatedAtLTE(v time.Time) predicate.UserOAuthProvider {
	return predicate.UserOAuthProvider(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldUpdatedAt), v))
	})
}

// UpdatedByEQ applies the EQ predicate on the "updated_by" field.
func UpdatedByEQ(v string) predicate.UserOAuthProvider {
	return predicate.UserOAuthProvider(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldUpdatedBy), v))
	})
}

// UpdatedByNEQ applies the NEQ predicate on the "updated_by" field.
func UpdatedByNEQ(v string) predicate.UserOAuthProvider {
	return predicate.UserOAuthProvider(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldUpdatedBy), v))
	})
}

// UpdatedByIn applies the In predicate on the "updated_by" field.
func UpdatedByIn(vs ...string) predicate.UserOAuthProvider {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.UserOAuthProvider(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldUpdatedBy), v...))
	})
}

// UpdatedByNotIn applies the NotIn predicate on the "updated_by" field.
func UpdatedByNotIn(vs ...string) predicate.UserOAuthProvider {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.UserOAuthProvider(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldUpdatedBy), v...))
	})
}

// UpdatedByGT applies the GT predicate on the "updated_by" field.
func UpdatedByGT(v string) predicate.UserOAuthProvider {
	return predicate.UserOAuthProvider(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldUpdatedBy), v))
	})
}

// UpdatedByGTE applies the GTE predicate on the "updated_by" field.
func UpdatedByGTE(v string) predicate.UserOAuthProvider {
	return predicate.UserOAuthProvider(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldUpdatedBy), v))
	})
}

// UpdatedByLT applies the LT predicate on the "updated_by" field.
func UpdatedByLT(v string) predicate.UserOAuthProvider {
	return predicate.UserOAuthProvider(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldUpdatedBy), v))
	})
}

// UpdatedByLTE applies the LTE predicate on the "updated_by" field.
func UpdatedByLTE(v string) predicate.UserOAuthProvider {
	return predicate.UserOAuthProvider(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldUpdatedBy), v))
	})
}

// UpdatedByContains applies the Contains predicate on the "updated_by" field.
func UpdatedByContains(v string) predicate.UserOAuthProvider {
	return predicate.UserOAuthProvider(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldUpdatedBy), v))
	})
}

// UpdatedByHasPrefix applies the HasPrefix predicate on the "updated_by" field.
func UpdatedByHasPrefix(v string) predicate.UserOAuthProvider {
	return predicate.UserOAuthProvider(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldUpdatedBy), v))
	})
}

// UpdatedByHasSuffix applies the HasSuffix predicate on the "updated_by" field.
func UpdatedByHasSuffix(v string) predicate.UserOAuthProvider {
	return predicate.UserOAuthProvider(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldUpdatedBy), v))
	})
}

// UpdatedByEqualFold applies the EqualFold predicate on the "updated_by" field.
func UpdatedByEqualFold(v string) predicate.UserOAuthProvider {
	return predicate.UserOAuthProvider(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldUpdatedBy), v))
	})
}

// UpdatedByContainsFold applies the ContainsFold predicate on the "updated_by" field.
func UpdatedByContainsFold(v string) predicate.UserOAuthProvider {
	return predicate.UserOAuthProvider(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldUpdatedBy), v))
	})
}

// StatusEQ applies the EQ predicate on the "status" field.
func StatusEQ(v Status) predicate.UserOAuthProvider {
	return predicate.UserOAuthProvider(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldStatus), v))
	})
}

// StatusNEQ applies the NEQ predicate on the "status" field.
func StatusNEQ(v Status) predicate.UserOAuthProvider {
	return predicate.UserOAuthProvider(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldStatus), v))
	})
}

// StatusIn applies the In predicate on the "status" field.
func StatusIn(vs ...Status) predicate.UserOAuthProvider {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.UserOAuthProvider(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldStatus), v...))
	})
}

// StatusNotIn applies the NotIn predicate on the "status" field.
func StatusNotIn(vs ...Status) predicate.UserOAuthProvider {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.UserOAuthProvider(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldStatus), v...))
	})
}

// ProviderEQ applies the EQ predicate on the "provider" field.
func ProviderEQ(v Provider) predicate.UserOAuthProvider {
	return predicate.UserOAuthProvider(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldProvider), v))
	})
}

// ProviderNEQ applies the NEQ predicate on the "provider" field.
func ProviderNEQ(v Provider) predicate.UserOAuthProvider {
	return predicate.UserOAuthProvider(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldProvider), v))
	})
}

// ProviderIn applies the In predicate on the "provider" field.
func ProviderIn(vs ...Provider) predicate.UserOAuthProvider {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.UserOAuthProvider(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldProvider), v...))
	})
}

// ProviderNotIn applies the NotIn predicate on the "provider" field.
func ProviderNotIn(vs ...Provider) predicate.UserOAuthProvider {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.UserOAuthProvider(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldProvider), v...))
	})
}

// ExpiryEQ applies the EQ predicate on the "expiry" field.
func ExpiryEQ(v time.Time) predicate.UserOAuthProvider {
	return predicate.UserOAuthProvider(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldExpiry), v))
	})
}

// ExpiryNEQ applies the NEQ predicate on the "expiry" field.
func ExpiryNEQ(v time.Time) predicate.UserOAuthProvider {
	return predicate.UserOAuthProvider(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldExpiry), v))
	})
}

// ExpiryIn applies the In predicate on the "expiry" field.
func ExpiryIn(vs ...time.Time) predicate.UserOAuthProvider {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.UserOAuthProvider(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldExpiry), v...))
	})
}

// ExpiryNotIn applies the NotIn predicate on the "expiry" field.
func ExpiryNotIn(vs ...time.Time) predicate.UserOAuthProvider {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.UserOAuthProvider(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldExpiry), v...))
	})
}

// ExpiryGT applies the GT predicate on the "expiry" field.
func ExpiryGT(v time.Time) predicate.UserOAuthProvider {
	return predicate.UserOAuthProvider(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldExpiry), v))
	})
}

// ExpiryGTE applies the GTE predicate on the "expiry" field.
func ExpiryGTE(v time.Time) predicate.UserOAuthProvider {
	return predicate.UserOAuthProvider(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldExpiry), v))
	})
}

// ExpiryLT applies the LT predicate on the "expiry" field.
func ExpiryLT(v time.Time) predicate.UserOAuthProvider {
	return predicate.UserOAuthProvider(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldExpiry), v))
	})
}

// ExpiryLTE applies the LTE predicate on the "expiry" field.
func ExpiryLTE(v time.Time) predicate.UserOAuthProvider {
	return predicate.UserOAuthProvider(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldExpiry), v))
	})
}

// AccessTokenEQ applies the EQ predicate on the "access_token" field.
func AccessTokenEQ(v string) predicate.UserOAuthProvider {
	return predicate.UserOAuthProvider(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldAccessToken), v))
	})
}

// AccessTokenNEQ applies the NEQ predicate on the "access_token" field.
func AccessTokenNEQ(v string) predicate.UserOAuthProvider {
	return predicate.UserOAuthProvider(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldAccessToken), v))
	})
}

// AccessTokenIn applies the In predicate on the "access_token" field.
func AccessTokenIn(vs ...string) predicate.UserOAuthProvider {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.UserOAuthProvider(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldAccessToken), v...))
	})
}

// AccessTokenNotIn applies the NotIn predicate on the "access_token" field.
func AccessTokenNotIn(vs ...string) predicate.UserOAuthProvider {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.UserOAuthProvider(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldAccessToken), v...))
	})
}

// AccessTokenGT applies the GT predicate on the "access_token" field.
func AccessTokenGT(v string) predicate.UserOAuthProvider {
	return predicate.UserOAuthProvider(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldAccessToken), v))
	})
}

// AccessTokenGTE applies the GTE predicate on the "access_token" field.
func AccessTokenGTE(v string) predicate.UserOAuthProvider {
	return predicate.UserOAuthProvider(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldAccessToken), v))
	})
}

// AccessTokenLT applies the LT predicate on the "access_token" field.
func AccessTokenLT(v string) predicate.UserOAuthProvider {
	return predicate.UserOAuthProvider(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldAccessToken), v))
	})
}

// AccessTokenLTE applies the LTE predicate on the "access_token" field.
func AccessTokenLTE(v string) predicate.UserOAuthProvider {
	return predicate.UserOAuthProvider(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldAccessToken), v))
	})
}

// AccessTokenContains applies the Contains predicate on the "access_token" field.
func AccessTokenContains(v string) predicate.UserOAuthProvider {
	return predicate.UserOAuthProvider(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldAccessToken), v))
	})
}

// AccessTokenHasPrefix applies the HasPrefix predicate on the "access_token" field.
func AccessTokenHasPrefix(v string) predicate.UserOAuthProvider {
	return predicate.UserOAuthProvider(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldAccessToken), v))
	})
}

// AccessTokenHasSuffix applies the HasSuffix predicate on the "access_token" field.
func AccessTokenHasSuffix(v string) predicate.UserOAuthProvider {
	return predicate.UserOAuthProvider(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldAccessToken), v))
	})
}

// AccessTokenEqualFold applies the EqualFold predicate on the "access_token" field.
func AccessTokenEqualFold(v string) predicate.UserOAuthProvider {
	return predicate.UserOAuthProvider(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldAccessToken), v))
	})
}

// AccessTokenContainsFold applies the ContainsFold predicate on the "access_token" field.
func AccessTokenContainsFold(v string) predicate.UserOAuthProvider {
	return predicate.UserOAuthProvider(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldAccessToken), v))
	})
}

// RefreshTokenEQ applies the EQ predicate on the "refresh_token" field.
func RefreshTokenEQ(v string) predicate.UserOAuthProvider {
	return predicate.UserOAuthProvider(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldRefreshToken), v))
	})
}

// RefreshTokenNEQ applies the NEQ predicate on the "refresh_token" field.
func RefreshTokenNEQ(v string) predicate.UserOAuthProvider {
	return predicate.UserOAuthProvider(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldRefreshToken), v))
	})
}

// RefreshTokenIn applies the In predicate on the "refresh_token" field.
func RefreshTokenIn(vs ...string) predicate.UserOAuthProvider {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.UserOAuthProvider(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldRefreshToken), v...))
	})
}

// RefreshTokenNotIn applies the NotIn predicate on the "refresh_token" field.
func RefreshTokenNotIn(vs ...string) predicate.UserOAuthProvider {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.UserOAuthProvider(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldRefreshToken), v...))
	})
}

// RefreshTokenGT applies the GT predicate on the "refresh_token" field.
func RefreshTokenGT(v string) predicate.UserOAuthProvider {
	return predicate.UserOAuthProvider(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldRefreshToken), v))
	})
}

// RefreshTokenGTE applies the GTE predicate on the "refresh_token" field.
func RefreshTokenGTE(v string) predicate.UserOAuthProvider {
	return predicate.UserOAuthProvider(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldRefreshToken), v))
	})
}

// RefreshTokenLT applies the LT predicate on the "refresh_token" field.
func RefreshTokenLT(v string) predicate.UserOAuthProvider {
	return predicate.UserOAuthProvider(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldRefreshToken), v))
	})
}

// RefreshTokenLTE applies the LTE predicate on the "refresh_token" field.
func RefreshTokenLTE(v string) predicate.UserOAuthProvider {
	return predicate.UserOAuthProvider(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldRefreshToken), v))
	})
}

// RefreshTokenContains applies the Contains predicate on the "refresh_token" field.
func RefreshTokenContains(v string) predicate.UserOAuthProvider {
	return predicate.UserOAuthProvider(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldRefreshToken), v))
	})
}

// RefreshTokenHasPrefix applies the HasPrefix predicate on the "refresh_token" field.
func RefreshTokenHasPrefix(v string) predicate.UserOAuthProvider {
	return predicate.UserOAuthProvider(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldRefreshToken), v))
	})
}

// RefreshTokenHasSuffix applies the HasSuffix predicate on the "refresh_token" field.
func RefreshTokenHasSuffix(v string) predicate.UserOAuthProvider {
	return predicate.UserOAuthProvider(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldRefreshToken), v))
	})
}

// RefreshTokenEqualFold applies the EqualFold predicate on the "refresh_token" field.
func RefreshTokenEqualFold(v string) predicate.UserOAuthProvider {
	return predicate.UserOAuthProvider(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldRefreshToken), v))
	})
}

// RefreshTokenContainsFold applies the ContainsFold predicate on the "refresh_token" field.
func RefreshTokenContainsFold(v string) predicate.UserOAuthProvider {
	return predicate.UserOAuthProvider(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldRefreshToken), v))
	})
}

// HasUser applies the HasEdge predicate on the "user" edge.
func HasUser() predicate.UserOAuthProvider {
	return predicate.UserOAuthProvider(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(UserTable, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, UserTable, UserColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasUserWith applies the HasEdge predicate on the "user" edge with a given conditions (other predicates).
func HasUserWith(preds ...predicate.User) predicate.UserOAuthProvider {
	return predicate.UserOAuthProvider(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(UserInverseTable, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, UserTable, UserColumn),
		)
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.UserOAuthProvider) predicate.UserOAuthProvider {
	return predicate.UserOAuthProvider(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for _, p := range predicates {
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.UserOAuthProvider) predicate.UserOAuthProvider {
	return predicate.UserOAuthProvider(func(s *sql.Selector) {
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
func Not(p predicate.UserOAuthProvider) predicate.UserOAuthProvider {
	return predicate.UserOAuthProvider(func(s *sql.Selector) {
		p(s.Not())
	})
}
