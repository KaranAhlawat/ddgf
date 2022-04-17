-- name: InsertPage :one
INSERT INTO "diary" ("page_id", "datetime", "content")
VALUES ($1, $2, $3)
RETURNING *;

-- name: SelectPage :one
SELECT *
FROM "diary"
WHERE "page_id" = $1
LIMIT 1;

-- name: SelectPages :many
SELECT *
FROM "diary"
ORDER BY "datetime";

-- name: UpdatePage :exec
UPDATE "diary"
SET "content" = $1,
"datetime" = $2
WHERE "page_id" = $3;

-- name: DeletePage :exec
DELETE FROM "diary"
WHERE "page_id" = $1;
