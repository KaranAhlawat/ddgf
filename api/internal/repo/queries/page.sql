-- name: InsertPage :one
INSERT INTO "pages" ("id", "datetime", "content")
VALUES ($1, $2, $3)
RETURNING *;

-- name: SelectPage :one
SELECT *
FROM "pages"
WHERE "id" = $1
LIMIT 1;

-- name: SelectPages :many
SELECT *
FROM "pages"
ORDER BY "datetime";

-- name: UpdatePage :exec
UPDATE "pages"
SET "content" = $1,
"datetime" = $2
WHERE "id" = $3;

-- name: DeletePage :exec
DELETE FROM "pages"
WHERE "id" = $1;
