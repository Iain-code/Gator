-- name: MarkFeedFetched :exec
UPDATE feeds 
SET last_fetched_at = Now(), updated_at = Now()
WHERE feeds.id = $1;

-- name: GetNextFeedToFetch :one
SELECT * FROM feeds
ORDER BY updated_at NULLS FIRST 
LIMIT 1;