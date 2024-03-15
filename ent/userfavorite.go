// Code generated by ent, DO NOT EDIT.

package ent

import (
	"fmt"
	"strings"
	"time"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
	"github.com/google/uuid"
	"github.com/k0kishima/golang-realworld-example-app/ent/article"
	"github.com/k0kishima/golang-realworld-example-app/ent/user"
	"github.com/k0kishima/golang-realworld-example-app/ent/userfavorite"
)

// UserFavorite is the model entity for the UserFavorite schema.
type UserFavorite struct {
	config `json:"-"`
	// ID of the ent.
	ID uuid.UUID `json:"id,omitempty"`
	// UserID holds the value of the "user_id" field.
	UserID uuid.UUID `json:"user_id,omitempty"`
	// ArticleID holds the value of the "article_id" field.
	ArticleID uuid.UUID `json:"article_id,omitempty"`
	// CreatedAt holds the value of the "created_at" field.
	CreatedAt time.Time `json:"created_at,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the UserFavoriteQuery when eager-loading is set.
	Edges        UserFavoriteEdges `json:"edges"`
	selectValues sql.SelectValues
}

// UserFavoriteEdges holds the relations/edges for other nodes in the graph.
type UserFavoriteEdges struct {
	// User holds the value of the user edge.
	User *User `json:"user,omitempty"`
	// Article holds the value of the article edge.
	Article *Article `json:"article,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [2]bool
}

// UserOrErr returns the User value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e UserFavoriteEdges) UserOrErr() (*User, error) {
	if e.User != nil {
		return e.User, nil
	} else if e.loadedTypes[0] {
		return nil, &NotFoundError{label: user.Label}
	}
	return nil, &NotLoadedError{edge: "user"}
}

// ArticleOrErr returns the Article value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e UserFavoriteEdges) ArticleOrErr() (*Article, error) {
	if e.Article != nil {
		return e.Article, nil
	} else if e.loadedTypes[1] {
		return nil, &NotFoundError{label: article.Label}
	}
	return nil, &NotLoadedError{edge: "article"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*UserFavorite) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case userfavorite.FieldCreatedAt:
			values[i] = new(sql.NullTime)
		case userfavorite.FieldID, userfavorite.FieldUserID, userfavorite.FieldArticleID:
			values[i] = new(uuid.UUID)
		default:
			values[i] = new(sql.UnknownType)
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the UserFavorite fields.
func (uf *UserFavorite) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case userfavorite.FieldID:
			if value, ok := values[i].(*uuid.UUID); !ok {
				return fmt.Errorf("unexpected type %T for field id", values[i])
			} else if value != nil {
				uf.ID = *value
			}
		case userfavorite.FieldUserID:
			if value, ok := values[i].(*uuid.UUID); !ok {
				return fmt.Errorf("unexpected type %T for field user_id", values[i])
			} else if value != nil {
				uf.UserID = *value
			}
		case userfavorite.FieldArticleID:
			if value, ok := values[i].(*uuid.UUID); !ok {
				return fmt.Errorf("unexpected type %T for field article_id", values[i])
			} else if value != nil {
				uf.ArticleID = *value
			}
		case userfavorite.FieldCreatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field created_at", values[i])
			} else if value.Valid {
				uf.CreatedAt = value.Time
			}
		default:
			uf.selectValues.Set(columns[i], values[i])
		}
	}
	return nil
}

// Value returns the ent.Value that was dynamically selected and assigned to the UserFavorite.
// This includes values selected through modifiers, order, etc.
func (uf *UserFavorite) Value(name string) (ent.Value, error) {
	return uf.selectValues.Get(name)
}

// QueryUser queries the "user" edge of the UserFavorite entity.
func (uf *UserFavorite) QueryUser() *UserQuery {
	return NewUserFavoriteClient(uf.config).QueryUser(uf)
}

// QueryArticle queries the "article" edge of the UserFavorite entity.
func (uf *UserFavorite) QueryArticle() *ArticleQuery {
	return NewUserFavoriteClient(uf.config).QueryArticle(uf)
}

// Update returns a builder for updating this UserFavorite.
// Note that you need to call UserFavorite.Unwrap() before calling this method if this UserFavorite
// was returned from a transaction, and the transaction was committed or rolled back.
func (uf *UserFavorite) Update() *UserFavoriteUpdateOne {
	return NewUserFavoriteClient(uf.config).UpdateOne(uf)
}

// Unwrap unwraps the UserFavorite entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (uf *UserFavorite) Unwrap() *UserFavorite {
	_tx, ok := uf.config.driver.(*txDriver)
	if !ok {
		panic("ent: UserFavorite is not a transactional entity")
	}
	uf.config.driver = _tx.drv
	return uf
}

// String implements the fmt.Stringer.
func (uf *UserFavorite) String() string {
	var builder strings.Builder
	builder.WriteString("UserFavorite(")
	builder.WriteString(fmt.Sprintf("id=%v, ", uf.ID))
	builder.WriteString("user_id=")
	builder.WriteString(fmt.Sprintf("%v", uf.UserID))
	builder.WriteString(", ")
	builder.WriteString("article_id=")
	builder.WriteString(fmt.Sprintf("%v", uf.ArticleID))
	builder.WriteString(", ")
	builder.WriteString("created_at=")
	builder.WriteString(uf.CreatedAt.Format(time.ANSIC))
	builder.WriteByte(')')
	return builder.String()
}

// UserFavorites is a parsable slice of UserFavorite.
type UserFavorites []*UserFavorite
