-- name: CreatePost :one
INSERT INTO posts (id, title, content, user_id)
VALUES ($1, $2, $3, $4) RETURNING *;
