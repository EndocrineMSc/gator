-- name: CreateFeedFollow :one
WITH ff AS (
    INSERT INTO feed_follows (id, created_at, updated_at, user_id, feed_id)
        VALUES ($1, $2, $3, $4, $5)
        RETURNING *)
SELECT ff.id, ff.created_at, ff.updated_at, ff.user_id, ff.feed_id, u.name AS user_name, f.name AS feed_name
FROM ff
         INNER JOIN feeds AS f ON ff.feed_id = f.id
         INNER JOIN users AS u ON ff.user_id = u.id;

-- name: GetFeedFollowsForUser :many
SELECT ff.*, u.name AS user_name, f.name AS feed_name
FROM feed_follows AS ff
INNER JOIN users AS u ON ff.user_id = u.id
INNER JOIN feeds AS f ON ff.feed_id = f.id
WHERE u.id = $1;

-- name: DeleteByUserIdAndFeedId :exec
DELETE FROM feed_follows WHERE user_id = $1 AND feed_id = $2;