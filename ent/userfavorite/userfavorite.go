// Code generated by ent, DO NOT EDIT.

package userfavorite

import (
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"github.com/google/uuid"
)

const (
	// Label holds the string label denoting the userfavorite type in the database.
	Label = "user_favorite"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldUserID holds the string denoting the user_id field in the database.
	FieldUserID = "user_id"
	// FieldArticleID holds the string denoting the article_id field in the database.
	FieldArticleID = "article_id"
	// FieldCreatedAt holds the string denoting the created_at field in the database.
	FieldCreatedAt = "created_at"
	// EdgeUser holds the string denoting the user edge name in mutations.
	EdgeUser = "user"
	// EdgeArticle holds the string denoting the article edge name in mutations.
	EdgeArticle = "article"
	// Table holds the table name of the userfavorite in the database.
	Table = "user_favorites"
	// UserTable is the table that holds the user relation/edge.
	UserTable = "user_favorites"
	// UserInverseTable is the table name for the User entity.
	// It exists in this package in order to avoid circular dependency with the "user" package.
	UserInverseTable = "users"
	// UserColumn is the table column denoting the user relation/edge.
	UserColumn = "user_id"
	// ArticleTable is the table that holds the article relation/edge.
	ArticleTable = "user_favorites"
	// ArticleInverseTable is the table name for the Article entity.
	// It exists in this package in order to avoid circular dependency with the "article" package.
	ArticleInverseTable = "articles"
	// ArticleColumn is the table column denoting the article relation/edge.
	ArticleColumn = "article_id"
)

// Columns holds all SQL columns for userfavorite fields.
var Columns = []string{
	FieldID,
	FieldUserID,
	FieldArticleID,
	FieldCreatedAt,
}

// ValidColumn reports if the column name is valid (part of the table columns).
func ValidColumn(column string) bool {
	for i := range Columns {
		if column == Columns[i] {
			return true
		}
	}
	return false
}

var (
	// DefaultCreatedAt holds the default value on creation for the "created_at" field.
	DefaultCreatedAt func() time.Time
	// DefaultID holds the default value on creation for the "id" field.
	DefaultID func() uuid.UUID
)

// OrderOption defines the ordering options for the UserFavorite queries.
type OrderOption func(*sql.Selector)

// ByID orders the results by the id field.
func ByID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldID, opts...).ToFunc()
}

// ByUserID orders the results by the user_id field.
func ByUserID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldUserID, opts...).ToFunc()
}

// ByArticleID orders the results by the article_id field.
func ByArticleID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldArticleID, opts...).ToFunc()
}

// ByCreatedAt orders the results by the created_at field.
func ByCreatedAt(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldCreatedAt, opts...).ToFunc()
}

// ByUserField orders the results by user field.
func ByUserField(field string, opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newUserStep(), sql.OrderByField(field, opts...))
	}
}

// ByArticleField orders the results by article field.
func ByArticleField(field string, opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newArticleStep(), sql.OrderByField(field, opts...))
	}
}
func newUserStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(UserInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.M2O, true, UserTable, UserColumn),
	)
}
func newArticleStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(ArticleInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.M2O, false, ArticleTable, ArticleColumn),
	)
}
