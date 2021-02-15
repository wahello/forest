// Code generated by entc, DO NOT EDIT.

package configuration

import (
	"time"

	"entgo.io/ent/dialect/sql"
	"github.com/vicanso/forest/ent/predicate"
	"github.com/vicanso/forest/ent/schema"
)

// ID filters vertices based on their ID field.
func ID(id int) predicate.Configuration {
	return predicate.Configuration(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldID), id))
	})
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id int) predicate.Configuration {
	return predicate.Configuration(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldID), id))
	})
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id int) predicate.Configuration {
	return predicate.Configuration(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldID), id))
	})
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...int) predicate.Configuration {
	return predicate.Configuration(func(s *sql.Selector) {
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
func IDNotIn(ids ...int) predicate.Configuration {
	return predicate.Configuration(func(s *sql.Selector) {
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
func IDGT(id int) predicate.Configuration {
	return predicate.Configuration(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldID), id))
	})
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id int) predicate.Configuration {
	return predicate.Configuration(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldID), id))
	})
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id int) predicate.Configuration {
	return predicate.Configuration(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldID), id))
	})
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id int) predicate.Configuration {
	return predicate.Configuration(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldID), id))
	})
}

// CreatedAt applies equality check predicate on the "created_at" field. It's identical to CreatedAtEQ.
func CreatedAt(v time.Time) predicate.Configuration {
	return predicate.Configuration(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldCreatedAt), v))
	})
}

// UpdatedAt applies equality check predicate on the "updated_at" field. It's identical to UpdatedAtEQ.
func UpdatedAt(v time.Time) predicate.Configuration {
	return predicate.Configuration(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldUpdatedAt), v))
	})
}

// Status applies equality check predicate on the "status" field. It's identical to StatusEQ.
func Status(v schema.Status) predicate.Configuration {
	vc := int8(v)
	return predicate.Configuration(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldStatus), vc))
	})
}

// Name applies equality check predicate on the "name" field. It's identical to NameEQ.
func Name(v string) predicate.Configuration {
	return predicate.Configuration(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldName), v))
	})
}

// Owner applies equality check predicate on the "owner" field. It's identical to OwnerEQ.
func Owner(v string) predicate.Configuration {
	return predicate.Configuration(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldOwner), v))
	})
}

// Data applies equality check predicate on the "data" field. It's identical to DataEQ.
func Data(v string) predicate.Configuration {
	return predicate.Configuration(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldData), v))
	})
}

// StartedAt applies equality check predicate on the "started_at" field. It's identical to StartedAtEQ.
func StartedAt(v time.Time) predicate.Configuration {
	return predicate.Configuration(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldStartedAt), v))
	})
}

// EndedAt applies equality check predicate on the "ended_at" field. It's identical to EndedAtEQ.
func EndedAt(v time.Time) predicate.Configuration {
	return predicate.Configuration(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldEndedAt), v))
	})
}

// CreatedAtEQ applies the EQ predicate on the "created_at" field.
func CreatedAtEQ(v time.Time) predicate.Configuration {
	return predicate.Configuration(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldCreatedAt), v))
	})
}

// CreatedAtNEQ applies the NEQ predicate on the "created_at" field.
func CreatedAtNEQ(v time.Time) predicate.Configuration {
	return predicate.Configuration(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldCreatedAt), v))
	})
}

// CreatedAtIn applies the In predicate on the "created_at" field.
func CreatedAtIn(vs ...time.Time) predicate.Configuration {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Configuration(func(s *sql.Selector) {
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
func CreatedAtNotIn(vs ...time.Time) predicate.Configuration {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Configuration(func(s *sql.Selector) {
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
func CreatedAtGT(v time.Time) predicate.Configuration {
	return predicate.Configuration(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldCreatedAt), v))
	})
}

// CreatedAtGTE applies the GTE predicate on the "created_at" field.
func CreatedAtGTE(v time.Time) predicate.Configuration {
	return predicate.Configuration(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldCreatedAt), v))
	})
}

// CreatedAtLT applies the LT predicate on the "created_at" field.
func CreatedAtLT(v time.Time) predicate.Configuration {
	return predicate.Configuration(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldCreatedAt), v))
	})
}

// CreatedAtLTE applies the LTE predicate on the "created_at" field.
func CreatedAtLTE(v time.Time) predicate.Configuration {
	return predicate.Configuration(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldCreatedAt), v))
	})
}

// UpdatedAtEQ applies the EQ predicate on the "updated_at" field.
func UpdatedAtEQ(v time.Time) predicate.Configuration {
	return predicate.Configuration(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldUpdatedAt), v))
	})
}

// UpdatedAtNEQ applies the NEQ predicate on the "updated_at" field.
func UpdatedAtNEQ(v time.Time) predicate.Configuration {
	return predicate.Configuration(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldUpdatedAt), v))
	})
}

// UpdatedAtIn applies the In predicate on the "updated_at" field.
func UpdatedAtIn(vs ...time.Time) predicate.Configuration {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Configuration(func(s *sql.Selector) {
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
func UpdatedAtNotIn(vs ...time.Time) predicate.Configuration {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Configuration(func(s *sql.Selector) {
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
func UpdatedAtGT(v time.Time) predicate.Configuration {
	return predicate.Configuration(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldUpdatedAt), v))
	})
}

// UpdatedAtGTE applies the GTE predicate on the "updated_at" field.
func UpdatedAtGTE(v time.Time) predicate.Configuration {
	return predicate.Configuration(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldUpdatedAt), v))
	})
}

// UpdatedAtLT applies the LT predicate on the "updated_at" field.
func UpdatedAtLT(v time.Time) predicate.Configuration {
	return predicate.Configuration(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldUpdatedAt), v))
	})
}

// UpdatedAtLTE applies the LTE predicate on the "updated_at" field.
func UpdatedAtLTE(v time.Time) predicate.Configuration {
	return predicate.Configuration(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldUpdatedAt), v))
	})
}

// StatusEQ applies the EQ predicate on the "status" field.
func StatusEQ(v schema.Status) predicate.Configuration {
	vc := int8(v)
	return predicate.Configuration(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldStatus), vc))
	})
}

// StatusNEQ applies the NEQ predicate on the "status" field.
func StatusNEQ(v schema.Status) predicate.Configuration {
	vc := int8(v)
	return predicate.Configuration(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldStatus), vc))
	})
}

// StatusIn applies the In predicate on the "status" field.
func StatusIn(vs ...schema.Status) predicate.Configuration {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = int8(vs[i])
	}
	return predicate.Configuration(func(s *sql.Selector) {
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
func StatusNotIn(vs ...schema.Status) predicate.Configuration {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = int8(vs[i])
	}
	return predicate.Configuration(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldStatus), v...))
	})
}

// StatusGT applies the GT predicate on the "status" field.
func StatusGT(v schema.Status) predicate.Configuration {
	vc := int8(v)
	return predicate.Configuration(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldStatus), vc))
	})
}

// StatusGTE applies the GTE predicate on the "status" field.
func StatusGTE(v schema.Status) predicate.Configuration {
	vc := int8(v)
	return predicate.Configuration(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldStatus), vc))
	})
}

// StatusLT applies the LT predicate on the "status" field.
func StatusLT(v schema.Status) predicate.Configuration {
	vc := int8(v)
	return predicate.Configuration(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldStatus), vc))
	})
}

// StatusLTE applies the LTE predicate on the "status" field.
func StatusLTE(v schema.Status) predicate.Configuration {
	vc := int8(v)
	return predicate.Configuration(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldStatus), vc))
	})
}

// NameEQ applies the EQ predicate on the "name" field.
func NameEQ(v string) predicate.Configuration {
	return predicate.Configuration(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldName), v))
	})
}

// NameNEQ applies the NEQ predicate on the "name" field.
func NameNEQ(v string) predicate.Configuration {
	return predicate.Configuration(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldName), v))
	})
}

// NameIn applies the In predicate on the "name" field.
func NameIn(vs ...string) predicate.Configuration {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Configuration(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldName), v...))
	})
}

// NameNotIn applies the NotIn predicate on the "name" field.
func NameNotIn(vs ...string) predicate.Configuration {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Configuration(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldName), v...))
	})
}

// NameGT applies the GT predicate on the "name" field.
func NameGT(v string) predicate.Configuration {
	return predicate.Configuration(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldName), v))
	})
}

// NameGTE applies the GTE predicate on the "name" field.
func NameGTE(v string) predicate.Configuration {
	return predicate.Configuration(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldName), v))
	})
}

// NameLT applies the LT predicate on the "name" field.
func NameLT(v string) predicate.Configuration {
	return predicate.Configuration(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldName), v))
	})
}

// NameLTE applies the LTE predicate on the "name" field.
func NameLTE(v string) predicate.Configuration {
	return predicate.Configuration(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldName), v))
	})
}

// NameContains applies the Contains predicate on the "name" field.
func NameContains(v string) predicate.Configuration {
	return predicate.Configuration(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldName), v))
	})
}

// NameHasPrefix applies the HasPrefix predicate on the "name" field.
func NameHasPrefix(v string) predicate.Configuration {
	return predicate.Configuration(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldName), v))
	})
}

// NameHasSuffix applies the HasSuffix predicate on the "name" field.
func NameHasSuffix(v string) predicate.Configuration {
	return predicate.Configuration(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldName), v))
	})
}

// NameEqualFold applies the EqualFold predicate on the "name" field.
func NameEqualFold(v string) predicate.Configuration {
	return predicate.Configuration(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldName), v))
	})
}

// NameContainsFold applies the ContainsFold predicate on the "name" field.
func NameContainsFold(v string) predicate.Configuration {
	return predicate.Configuration(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldName), v))
	})
}

// CategoryEQ applies the EQ predicate on the "category" field.
func CategoryEQ(v Category) predicate.Configuration {
	return predicate.Configuration(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldCategory), v))
	})
}

// CategoryNEQ applies the NEQ predicate on the "category" field.
func CategoryNEQ(v Category) predicate.Configuration {
	return predicate.Configuration(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldCategory), v))
	})
}

// CategoryIn applies the In predicate on the "category" field.
func CategoryIn(vs ...Category) predicate.Configuration {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Configuration(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldCategory), v...))
	})
}

// CategoryNotIn applies the NotIn predicate on the "category" field.
func CategoryNotIn(vs ...Category) predicate.Configuration {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Configuration(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldCategory), v...))
	})
}

// OwnerEQ applies the EQ predicate on the "owner" field.
func OwnerEQ(v string) predicate.Configuration {
	return predicate.Configuration(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldOwner), v))
	})
}

// OwnerNEQ applies the NEQ predicate on the "owner" field.
func OwnerNEQ(v string) predicate.Configuration {
	return predicate.Configuration(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldOwner), v))
	})
}

// OwnerIn applies the In predicate on the "owner" field.
func OwnerIn(vs ...string) predicate.Configuration {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Configuration(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldOwner), v...))
	})
}

// OwnerNotIn applies the NotIn predicate on the "owner" field.
func OwnerNotIn(vs ...string) predicate.Configuration {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Configuration(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldOwner), v...))
	})
}

// OwnerGT applies the GT predicate on the "owner" field.
func OwnerGT(v string) predicate.Configuration {
	return predicate.Configuration(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldOwner), v))
	})
}

// OwnerGTE applies the GTE predicate on the "owner" field.
func OwnerGTE(v string) predicate.Configuration {
	return predicate.Configuration(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldOwner), v))
	})
}

// OwnerLT applies the LT predicate on the "owner" field.
func OwnerLT(v string) predicate.Configuration {
	return predicate.Configuration(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldOwner), v))
	})
}

// OwnerLTE applies the LTE predicate on the "owner" field.
func OwnerLTE(v string) predicate.Configuration {
	return predicate.Configuration(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldOwner), v))
	})
}

// OwnerContains applies the Contains predicate on the "owner" field.
func OwnerContains(v string) predicate.Configuration {
	return predicate.Configuration(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldOwner), v))
	})
}

// OwnerHasPrefix applies the HasPrefix predicate on the "owner" field.
func OwnerHasPrefix(v string) predicate.Configuration {
	return predicate.Configuration(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldOwner), v))
	})
}

// OwnerHasSuffix applies the HasSuffix predicate on the "owner" field.
func OwnerHasSuffix(v string) predicate.Configuration {
	return predicate.Configuration(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldOwner), v))
	})
}

// OwnerEqualFold applies the EqualFold predicate on the "owner" field.
func OwnerEqualFold(v string) predicate.Configuration {
	return predicate.Configuration(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldOwner), v))
	})
}

// OwnerContainsFold applies the ContainsFold predicate on the "owner" field.
func OwnerContainsFold(v string) predicate.Configuration {
	return predicate.Configuration(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldOwner), v))
	})
}

// DataEQ applies the EQ predicate on the "data" field.
func DataEQ(v string) predicate.Configuration {
	return predicate.Configuration(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldData), v))
	})
}

// DataNEQ applies the NEQ predicate on the "data" field.
func DataNEQ(v string) predicate.Configuration {
	return predicate.Configuration(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldData), v))
	})
}

// DataIn applies the In predicate on the "data" field.
func DataIn(vs ...string) predicate.Configuration {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Configuration(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldData), v...))
	})
}

// DataNotIn applies the NotIn predicate on the "data" field.
func DataNotIn(vs ...string) predicate.Configuration {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Configuration(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldData), v...))
	})
}

// DataGT applies the GT predicate on the "data" field.
func DataGT(v string) predicate.Configuration {
	return predicate.Configuration(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldData), v))
	})
}

// DataGTE applies the GTE predicate on the "data" field.
func DataGTE(v string) predicate.Configuration {
	return predicate.Configuration(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldData), v))
	})
}

// DataLT applies the LT predicate on the "data" field.
func DataLT(v string) predicate.Configuration {
	return predicate.Configuration(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldData), v))
	})
}

// DataLTE applies the LTE predicate on the "data" field.
func DataLTE(v string) predicate.Configuration {
	return predicate.Configuration(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldData), v))
	})
}

// DataContains applies the Contains predicate on the "data" field.
func DataContains(v string) predicate.Configuration {
	return predicate.Configuration(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldData), v))
	})
}

// DataHasPrefix applies the HasPrefix predicate on the "data" field.
func DataHasPrefix(v string) predicate.Configuration {
	return predicate.Configuration(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldData), v))
	})
}

// DataHasSuffix applies the HasSuffix predicate on the "data" field.
func DataHasSuffix(v string) predicate.Configuration {
	return predicate.Configuration(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldData), v))
	})
}

// DataEqualFold applies the EqualFold predicate on the "data" field.
func DataEqualFold(v string) predicate.Configuration {
	return predicate.Configuration(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldData), v))
	})
}

// DataContainsFold applies the ContainsFold predicate on the "data" field.
func DataContainsFold(v string) predicate.Configuration {
	return predicate.Configuration(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldData), v))
	})
}

// StartedAtEQ applies the EQ predicate on the "started_at" field.
func StartedAtEQ(v time.Time) predicate.Configuration {
	return predicate.Configuration(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldStartedAt), v))
	})
}

// StartedAtNEQ applies the NEQ predicate on the "started_at" field.
func StartedAtNEQ(v time.Time) predicate.Configuration {
	return predicate.Configuration(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldStartedAt), v))
	})
}

// StartedAtIn applies the In predicate on the "started_at" field.
func StartedAtIn(vs ...time.Time) predicate.Configuration {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Configuration(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldStartedAt), v...))
	})
}

// StartedAtNotIn applies the NotIn predicate on the "started_at" field.
func StartedAtNotIn(vs ...time.Time) predicate.Configuration {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Configuration(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldStartedAt), v...))
	})
}

// StartedAtGT applies the GT predicate on the "started_at" field.
func StartedAtGT(v time.Time) predicate.Configuration {
	return predicate.Configuration(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldStartedAt), v))
	})
}

// StartedAtGTE applies the GTE predicate on the "started_at" field.
func StartedAtGTE(v time.Time) predicate.Configuration {
	return predicate.Configuration(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldStartedAt), v))
	})
}

// StartedAtLT applies the LT predicate on the "started_at" field.
func StartedAtLT(v time.Time) predicate.Configuration {
	return predicate.Configuration(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldStartedAt), v))
	})
}

// StartedAtLTE applies the LTE predicate on the "started_at" field.
func StartedAtLTE(v time.Time) predicate.Configuration {
	return predicate.Configuration(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldStartedAt), v))
	})
}

// EndedAtEQ applies the EQ predicate on the "ended_at" field.
func EndedAtEQ(v time.Time) predicate.Configuration {
	return predicate.Configuration(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldEndedAt), v))
	})
}

// EndedAtNEQ applies the NEQ predicate on the "ended_at" field.
func EndedAtNEQ(v time.Time) predicate.Configuration {
	return predicate.Configuration(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldEndedAt), v))
	})
}

// EndedAtIn applies the In predicate on the "ended_at" field.
func EndedAtIn(vs ...time.Time) predicate.Configuration {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Configuration(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldEndedAt), v...))
	})
}

// EndedAtNotIn applies the NotIn predicate on the "ended_at" field.
func EndedAtNotIn(vs ...time.Time) predicate.Configuration {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Configuration(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldEndedAt), v...))
	})
}

// EndedAtGT applies the GT predicate on the "ended_at" field.
func EndedAtGT(v time.Time) predicate.Configuration {
	return predicate.Configuration(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldEndedAt), v))
	})
}

// EndedAtGTE applies the GTE predicate on the "ended_at" field.
func EndedAtGTE(v time.Time) predicate.Configuration {
	return predicate.Configuration(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldEndedAt), v))
	})
}

// EndedAtLT applies the LT predicate on the "ended_at" field.
func EndedAtLT(v time.Time) predicate.Configuration {
	return predicate.Configuration(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldEndedAt), v))
	})
}

// EndedAtLTE applies the LTE predicate on the "ended_at" field.
func EndedAtLTE(v time.Time) predicate.Configuration {
	return predicate.Configuration(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldEndedAt), v))
	})
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.Configuration) predicate.Configuration {
	return predicate.Configuration(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for _, p := range predicates {
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.Configuration) predicate.Configuration {
	return predicate.Configuration(func(s *sql.Selector) {
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
func Not(p predicate.Configuration) predicate.Configuration {
	return predicate.Configuration(func(s *sql.Selector) {
		p(s.Not())
	})
}
