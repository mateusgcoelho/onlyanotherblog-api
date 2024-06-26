// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.25.0
// source: post.sql

package database

import (
	"context"

	"github.com/jackc/pgx/v5/pgtype"
)

const createPost = `-- name: CreatePost :one
INSERT INTO posts (id, title, content, user_id)
VALUES ($1, $2, $3, $4) RETURNING id, title, content, user_id, created_at, updated_at
`

type CreatePostParams struct {
	ID      string      `json:"id"`
	Title   pgtype.Text `json:"title"`
	Content pgtype.Text `json:"content"`
	UserID  pgtype.Text `json:"user_id"`
}

func (q *Queries) CreatePost(ctx context.Context, arg CreatePostParams) (Post, error) {
	row := q.db.QueryRow(ctx, createPost,
		arg.ID,
		arg.Title,
		arg.Content,
		arg.UserID,
	)
	var i Post
	err := row.Scan(
		&i.ID,
		&i.Title,
		&i.Content,
		&i.UserID,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}

const getPost = `-- name: GetPost :one
SELECT
    p.id as post_id, p.title as post_title, p.content as post_content,
    p.created_at as post_created_at, p.updated_at as post_updated_at,
    u.username as username
FROM posts AS p
INNER JOIN users AS u ON u.id = p.user_id
WHERE p.id = $1
`

type GetPostRow struct {
	PostID        string           `json:"post_id"`
	PostTitle     pgtype.Text      `json:"post_title"`
	PostContent   pgtype.Text      `json:"post_content"`
	PostCreatedAt pgtype.Timestamp `json:"post_created_at"`
	PostUpdatedAt pgtype.Timestamp `json:"post_updated_at"`
	Username      pgtype.Text      `json:"username"`
}

func (q *Queries) GetPost(ctx context.Context, id string) (GetPostRow, error) {
	row := q.db.QueryRow(ctx, getPost, id)
	var i GetPostRow
	err := row.Scan(
		&i.PostID,
		&i.PostTitle,
		&i.PostContent,
		&i.PostCreatedAt,
		&i.PostUpdatedAt,
		&i.Username,
	)
	return i, err
}

const getPosts = `-- name: GetPosts :many
SELECT
    p.id as post_id, p.title as post_title, p.content as post_content,
    p.created_at as post_created_at, p.updated_at as post_updated_at,
    u.username as username
FROM posts AS p
INNER JOIN users AS u ON u.id = p.user_id
WHERE p.id > $1
ORDER BY p.created_at DESC, p.id ASC
FETCH FIRST $2 ROWS ONLY
`

type GetPostsParams struct {
	ID    string `json:"id"`
	Limit int32  `json:"limit"`
}

type GetPostsRow struct {
	PostID        string           `json:"post_id"`
	PostTitle     pgtype.Text      `json:"post_title"`
	PostContent   pgtype.Text      `json:"post_content"`
	PostCreatedAt pgtype.Timestamp `json:"post_created_at"`
	PostUpdatedAt pgtype.Timestamp `json:"post_updated_at"`
	Username      pgtype.Text      `json:"username"`
}

func (q *Queries) GetPosts(ctx context.Context, arg GetPostsParams) ([]GetPostsRow, error) {
	rows, err := q.db.Query(ctx, getPosts, arg.ID, arg.Limit)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []GetPostsRow{}
	for rows.Next() {
		var i GetPostsRow
		if err := rows.Scan(
			&i.PostID,
			&i.PostTitle,
			&i.PostContent,
			&i.PostCreatedAt,
			&i.PostUpdatedAt,
			&i.Username,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const getPostsOfUser = `-- name: GetPostsOfUser :many
SELECT
    p.id as post_id, p.title as post_title, p.content as post_content,
    p.created_at as post_created_at, p.updated_at as post_updated_at,
    u.username as username
FROM posts AS p
INNER JOIN users AS u ON u.id = p.user_id
WHERE u.username = $1 OR u.id = $1
ORDER BY p.created_at DESC
`

type GetPostsOfUserRow struct {
	PostID        string           `json:"post_id"`
	PostTitle     pgtype.Text      `json:"post_title"`
	PostContent   pgtype.Text      `json:"post_content"`
	PostCreatedAt pgtype.Timestamp `json:"post_created_at"`
	PostUpdatedAt pgtype.Timestamp `json:"post_updated_at"`
	Username      pgtype.Text      `json:"username"`
}

func (q *Queries) GetPostsOfUser(ctx context.Context, username pgtype.Text) ([]GetPostsOfUserRow, error) {
	rows, err := q.db.Query(ctx, getPostsOfUser, username)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []GetPostsOfUserRow{}
	for rows.Next() {
		var i GetPostsOfUserRow
		if err := rows.Scan(
			&i.PostID,
			&i.PostTitle,
			&i.PostContent,
			&i.PostCreatedAt,
			&i.PostUpdatedAt,
			&i.Username,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}
