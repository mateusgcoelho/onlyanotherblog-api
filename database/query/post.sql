-- name: CreatePost :one
INSERT INTO posts (id, title, content, user_id)
VALUES ($1, $2, $3, $4) RETURNING *;

-- name: GetPosts :many
SELECT
    p.id as post_id, p.title as post_title, p.content as post_content,
    p.created_at as post_created_at, p.updated_at as post_updated_at,
    u.username as username
FROM posts AS p
INNER JOIN users AS u ON u.id = p.user_id
WHERE p.id > $1
ORDER BY p.created_at DESC, p.id ASC
FETCH FIRST $2 ROWS ONLY;

-- name: GetPost :one
SELECT
    p.id as post_id, p.title as post_title, p.content as post_content,
    p.created_at as post_created_at, p.updated_at as post_updated_at,
    u.username as username
FROM posts AS p
INNER JOIN users AS u ON u.id = p.user_id
WHERE p.id = $1;
