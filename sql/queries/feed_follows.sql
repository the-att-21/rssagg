-- name: CreateFeedFollow :one
INSERT INTO feed_follows (id, created_at, updated_at, users_id, feed_id) 
VALUES ($1, $2, $3, $4, $5) 
RETURNING *;  

-- name: GetFeedFollows :many
SELECT * FROM feed_follows WHERE users_id = $1;

-- name: DeleteFeedFollow :exec
DELETE FROM feed_follows WHERE id = $1 AND users_id = $2;