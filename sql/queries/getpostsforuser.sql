-- name: GetPostsForUser :many
SELECT * FROM posts
ORDER BY published_at NULLS FIRST
LIMIT $1;

