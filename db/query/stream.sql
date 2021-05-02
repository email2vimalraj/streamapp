-- name: CreateStream :one
INSERT INTO streams (stream_name, stream_link, username)
VALUES ($1, $2, $3)
RETURNING *;
-- name: GetStream :one
SELECT *
FROM streams
WHERE stream_name = $1
LIMIT 1;
-- name: ListStreams :many
SELECT *
FROM streams
WHERE username = $1
ORDER BY id
LIMIT $2 OFFSET $3;