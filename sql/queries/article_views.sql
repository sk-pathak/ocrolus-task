-- name: UpsertArticleView :exec
INSERT INTO article_views (user_id, article_id, viewed_at)
VALUES ($1, $2, $3)
ON CONFLICT (user_id, article_id)
DO UPDATE SET viewed_at = EXCLUDED.viewed_at;

-- name: GetRecentlyViewedArticles :many
SELECT a.*
FROM article_views av
JOIN articles a ON av.article_id = a.id
WHERE av.user_id = $1
ORDER BY av.viewed_at DESC
LIMIT 15;

-- name: DeleteOldArticleViews :exec
DELETE FROM article_views
WHERE article_views.user_id = $1
  AND id NOT IN (
    SELECT id
    FROM article_views
    WHERE user_id = $1
    ORDER BY viewed_at DESC
    LIMIT $2
  );
