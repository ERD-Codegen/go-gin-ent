// Code generated by ent, DO NOT EDIT.

package migrate

import (
	"entgo.io/ent/dialect/sql/schema"
	"entgo.io/ent/schema/field"
)

var (
	// ArticlesColumns holds the columns for the "articles" table.
	ArticlesColumns = []*schema.Column{
		{Name: "id", Type: field.TypeUUID},
		{Name: "author_id", Type: field.TypeUUID},
		{Name: "slug", Type: field.TypeString, Unique: true, Size: 255},
		{Name: "title", Type: field.TypeString, Size: 255},
		{Name: "description", Type: field.TypeString, Size: 255},
		{Name: "body", Type: field.TypeString, Size: 4096},
		{Name: "created_at", Type: field.TypeTime},
		{Name: "updated_at", Type: field.TypeTime},
	}
	// ArticlesTable holds the schema information for the "articles" table.
	ArticlesTable = &schema.Table{
		Name:       "articles",
		Columns:    ArticlesColumns,
		PrimaryKey: []*schema.Column{ArticlesColumns[0]},
		Indexes: []*schema.Index{
			{
				Name:    "article_slug",
				Unique:  true,
				Columns: []*schema.Column{ArticlesColumns[2]},
			},
		},
	}
	// ArticleTagsColumns holds the columns for the "article_tags" table.
	ArticleTagsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeUUID},
		{Name: "article_id", Type: field.TypeUUID},
		{Name: "tag_id", Type: field.TypeUUID},
		{Name: "created_at", Type: field.TypeTime},
	}
	// ArticleTagsTable holds the schema information for the "article_tags" table.
	ArticleTagsTable = &schema.Table{
		Name:       "article_tags",
		Columns:    ArticleTagsColumns,
		PrimaryKey: []*schema.Column{ArticleTagsColumns[0]},
		Indexes: []*schema.Index{
			{
				Name:    "articletag_article_id_tag_id",
				Unique:  true,
				Columns: []*schema.Column{ArticleTagsColumns[1], ArticleTagsColumns[2]},
			},
		},
	}
	// CommentsColumns holds the columns for the "comments" table.
	CommentsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeUUID},
		{Name: "author_id", Type: field.TypeUUID},
		{Name: "article_id", Type: field.TypeUUID},
		{Name: "body", Type: field.TypeString, Size: 4096},
		{Name: "created_at", Type: field.TypeTime},
		{Name: "updated_at", Type: field.TypeTime},
	}
	// CommentsTable holds the schema information for the "comments" table.
	CommentsTable = &schema.Table{
		Name:       "comments",
		Columns:    CommentsColumns,
		PrimaryKey: []*schema.Column{CommentsColumns[0]},
	}
	// TagsColumns holds the columns for the "tags" table.
	TagsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeUUID},
		{Name: "description", Type: field.TypeString, Unique: true},
		{Name: "created_at", Type: field.TypeTime},
	}
	// TagsTable holds the schema information for the "tags" table.
	TagsTable = &schema.Table{
		Name:       "tags",
		Columns:    TagsColumns,
		PrimaryKey: []*schema.Column{TagsColumns[0]},
	}
	// UsersColumns holds the columns for the "users" table.
	UsersColumns = []*schema.Column{
		{Name: "id", Type: field.TypeUUID, Unique: true},
		{Name: "username", Type: field.TypeString, Unique: true},
		{Name: "email", Type: field.TypeString, Unique: true},
		{Name: "password", Type: field.TypeString},
		{Name: "image", Type: field.TypeString, Default: "https://api.realworld.io/images/smiley-cyrus.jpeg"},
		{Name: "bio", Type: field.TypeString, Default: ""},
		{Name: "created_at", Type: field.TypeTime},
		{Name: "updated_at", Type: field.TypeTime},
	}
	// UsersTable holds the schema information for the "users" table.
	UsersTable = &schema.Table{
		Name:       "users",
		Columns:    UsersColumns,
		PrimaryKey: []*schema.Column{UsersColumns[0]},
		Indexes: []*schema.Index{
			{
				Name:    "user_username_email",
				Unique:  true,
				Columns: []*schema.Column{UsersColumns[1], UsersColumns[2]},
			},
		},
	}
	// UserFavoritesColumns holds the columns for the "user_favorites" table.
	UserFavoritesColumns = []*schema.Column{
		{Name: "id", Type: field.TypeUUID},
		{Name: "user_id", Type: field.TypeUUID},
		{Name: "article_id", Type: field.TypeUUID},
		{Name: "created_at", Type: field.TypeTime},
	}
	// UserFavoritesTable holds the schema information for the "user_favorites" table.
	UserFavoritesTable = &schema.Table{
		Name:       "user_favorites",
		Columns:    UserFavoritesColumns,
		PrimaryKey: []*schema.Column{UserFavoritesColumns[0]},
		Indexes: []*schema.Index{
			{
				Name:    "userfavorite_user_id_article_id",
				Unique:  true,
				Columns: []*schema.Column{UserFavoritesColumns[1], UserFavoritesColumns[2]},
			},
		},
	}
	// UserFollowsColumns holds the columns for the "user_follows" table.
	UserFollowsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeUUID},
		{Name: "created_at", Type: field.TypeTime},
		{Name: "follower_id", Type: field.TypeUUID, Unique: true},
		{Name: "followee_id", Type: field.TypeUUID},
	}
	// UserFollowsTable holds the schema information for the "user_follows" table.
	UserFollowsTable = &schema.Table{
		Name:       "user_follows",
		Columns:    UserFollowsColumns,
		PrimaryKey: []*schema.Column{UserFollowsColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "user_follows_users_follows",
				Columns:    []*schema.Column{UserFollowsColumns[2]},
				RefColumns: []*schema.Column{UsersColumns[0]},
				OnDelete:   schema.NoAction,
			},
			{
				Symbol:     "user_follows_users_followee",
				Columns:    []*schema.Column{UserFollowsColumns[3]},
				RefColumns: []*schema.Column{UsersColumns[0]},
				OnDelete:   schema.NoAction,
			},
		},
		Indexes: []*schema.Index{
			{
				Name:    "userfollow_follower_id_followee_id",
				Unique:  true,
				Columns: []*schema.Column{UserFollowsColumns[2], UserFollowsColumns[3]},
			},
		},
	}
	// Tables holds all the tables in the schema.
	Tables = []*schema.Table{
		ArticlesTable,
		ArticleTagsTable,
		CommentsTable,
		TagsTable,
		UsersTable,
		UserFavoritesTable,
		UserFollowsTable,
	}
)

func init() {
	UserFollowsTable.ForeignKeys[0].RefTable = UsersTable
	UserFollowsTable.ForeignKeys[1].RefTable = UsersTable
}
