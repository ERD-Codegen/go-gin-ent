// Code generated by ent, DO NOT EDIT.

package articletag

import (
	"time"

	"entgo.io/ent/dialect/sql"
	"github.com/google/uuid"
	"github.com/k0kishima/golang-realworld-example-app/ent/predicate"
)

// ID filters vertices based on their ID field.
func ID(id uuid.UUID) predicate.ArticleTag {
	return predicate.ArticleTag(sql.FieldEQ(FieldID, id))
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id uuid.UUID) predicate.ArticleTag {
	return predicate.ArticleTag(sql.FieldEQ(FieldID, id))
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id uuid.UUID) predicate.ArticleTag {
	return predicate.ArticleTag(sql.FieldNEQ(FieldID, id))
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...uuid.UUID) predicate.ArticleTag {
	return predicate.ArticleTag(sql.FieldIn(FieldID, ids...))
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...uuid.UUID) predicate.ArticleTag {
	return predicate.ArticleTag(sql.FieldNotIn(FieldID, ids...))
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id uuid.UUID) predicate.ArticleTag {
	return predicate.ArticleTag(sql.FieldGT(FieldID, id))
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id uuid.UUID) predicate.ArticleTag {
	return predicate.ArticleTag(sql.FieldGTE(FieldID, id))
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id uuid.UUID) predicate.ArticleTag {
	return predicate.ArticleTag(sql.FieldLT(FieldID, id))
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id uuid.UUID) predicate.ArticleTag {
	return predicate.ArticleTag(sql.FieldLTE(FieldID, id))
}

// ArticleID applies equality check predicate on the "article_id" field. It's identical to ArticleIDEQ.
func ArticleID(v uuid.UUID) predicate.ArticleTag {
	return predicate.ArticleTag(sql.FieldEQ(FieldArticleID, v))
}

// TagID applies equality check predicate on the "tag_id" field. It's identical to TagIDEQ.
func TagID(v uuid.UUID) predicate.ArticleTag {
	return predicate.ArticleTag(sql.FieldEQ(FieldTagID, v))
}

// CreatedAt applies equality check predicate on the "created_at" field. It's identical to CreatedAtEQ.
func CreatedAt(v time.Time) predicate.ArticleTag {
	return predicate.ArticleTag(sql.FieldEQ(FieldCreatedAt, v))
}

// ArticleIDEQ applies the EQ predicate on the "article_id" field.
func ArticleIDEQ(v uuid.UUID) predicate.ArticleTag {
	return predicate.ArticleTag(sql.FieldEQ(FieldArticleID, v))
}

// ArticleIDNEQ applies the NEQ predicate on the "article_id" field.
func ArticleIDNEQ(v uuid.UUID) predicate.ArticleTag {
	return predicate.ArticleTag(sql.FieldNEQ(FieldArticleID, v))
}

// ArticleIDIn applies the In predicate on the "article_id" field.
func ArticleIDIn(vs ...uuid.UUID) predicate.ArticleTag {
	return predicate.ArticleTag(sql.FieldIn(FieldArticleID, vs...))
}

// ArticleIDNotIn applies the NotIn predicate on the "article_id" field.
func ArticleIDNotIn(vs ...uuid.UUID) predicate.ArticleTag {
	return predicate.ArticleTag(sql.FieldNotIn(FieldArticleID, vs...))
}

// ArticleIDGT applies the GT predicate on the "article_id" field.
func ArticleIDGT(v uuid.UUID) predicate.ArticleTag {
	return predicate.ArticleTag(sql.FieldGT(FieldArticleID, v))
}

// ArticleIDGTE applies the GTE predicate on the "article_id" field.
func ArticleIDGTE(v uuid.UUID) predicate.ArticleTag {
	return predicate.ArticleTag(sql.FieldGTE(FieldArticleID, v))
}

// ArticleIDLT applies the LT predicate on the "article_id" field.
func ArticleIDLT(v uuid.UUID) predicate.ArticleTag {
	return predicate.ArticleTag(sql.FieldLT(FieldArticleID, v))
}

// ArticleIDLTE applies the LTE predicate on the "article_id" field.
func ArticleIDLTE(v uuid.UUID) predicate.ArticleTag {
	return predicate.ArticleTag(sql.FieldLTE(FieldArticleID, v))
}

// TagIDEQ applies the EQ predicate on the "tag_id" field.
func TagIDEQ(v uuid.UUID) predicate.ArticleTag {
	return predicate.ArticleTag(sql.FieldEQ(FieldTagID, v))
}

// TagIDNEQ applies the NEQ predicate on the "tag_id" field.
func TagIDNEQ(v uuid.UUID) predicate.ArticleTag {
	return predicate.ArticleTag(sql.FieldNEQ(FieldTagID, v))
}

// TagIDIn applies the In predicate on the "tag_id" field.
func TagIDIn(vs ...uuid.UUID) predicate.ArticleTag {
	return predicate.ArticleTag(sql.FieldIn(FieldTagID, vs...))
}

// TagIDNotIn applies the NotIn predicate on the "tag_id" field.
func TagIDNotIn(vs ...uuid.UUID) predicate.ArticleTag {
	return predicate.ArticleTag(sql.FieldNotIn(FieldTagID, vs...))
}

// TagIDGT applies the GT predicate on the "tag_id" field.
func TagIDGT(v uuid.UUID) predicate.ArticleTag {
	return predicate.ArticleTag(sql.FieldGT(FieldTagID, v))
}

// TagIDGTE applies the GTE predicate on the "tag_id" field.
func TagIDGTE(v uuid.UUID) predicate.ArticleTag {
	return predicate.ArticleTag(sql.FieldGTE(FieldTagID, v))
}

// TagIDLT applies the LT predicate on the "tag_id" field.
func TagIDLT(v uuid.UUID) predicate.ArticleTag {
	return predicate.ArticleTag(sql.FieldLT(FieldTagID, v))
}

// TagIDLTE applies the LTE predicate on the "tag_id" field.
func TagIDLTE(v uuid.UUID) predicate.ArticleTag {
	return predicate.ArticleTag(sql.FieldLTE(FieldTagID, v))
}

// CreatedAtEQ applies the EQ predicate on the "created_at" field.
func CreatedAtEQ(v time.Time) predicate.ArticleTag {
	return predicate.ArticleTag(sql.FieldEQ(FieldCreatedAt, v))
}

// CreatedAtNEQ applies the NEQ predicate on the "created_at" field.
func CreatedAtNEQ(v time.Time) predicate.ArticleTag {
	return predicate.ArticleTag(sql.FieldNEQ(FieldCreatedAt, v))
}

// CreatedAtIn applies the In predicate on the "created_at" field.
func CreatedAtIn(vs ...time.Time) predicate.ArticleTag {
	return predicate.ArticleTag(sql.FieldIn(FieldCreatedAt, vs...))
}

// CreatedAtNotIn applies the NotIn predicate on the "created_at" field.
func CreatedAtNotIn(vs ...time.Time) predicate.ArticleTag {
	return predicate.ArticleTag(sql.FieldNotIn(FieldCreatedAt, vs...))
}

// CreatedAtGT applies the GT predicate on the "created_at" field.
func CreatedAtGT(v time.Time) predicate.ArticleTag {
	return predicate.ArticleTag(sql.FieldGT(FieldCreatedAt, v))
}

// CreatedAtGTE applies the GTE predicate on the "created_at" field.
func CreatedAtGTE(v time.Time) predicate.ArticleTag {
	return predicate.ArticleTag(sql.FieldGTE(FieldCreatedAt, v))
}

// CreatedAtLT applies the LT predicate on the "created_at" field.
func CreatedAtLT(v time.Time) predicate.ArticleTag {
	return predicate.ArticleTag(sql.FieldLT(FieldCreatedAt, v))
}

// CreatedAtLTE applies the LTE predicate on the "created_at" field.
func CreatedAtLTE(v time.Time) predicate.ArticleTag {
	return predicate.ArticleTag(sql.FieldLTE(FieldCreatedAt, v))
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.ArticleTag) predicate.ArticleTag {
	return predicate.ArticleTag(sql.AndPredicates(predicates...))
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.ArticleTag) predicate.ArticleTag {
	return predicate.ArticleTag(sql.OrPredicates(predicates...))
}

// Not applies the not operator on the given predicate.
func Not(p predicate.ArticleTag) predicate.ArticleTag {
	return predicate.ArticleTag(sql.NotPredicates(p))
}
