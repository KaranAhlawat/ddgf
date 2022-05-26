-- name: InsertTag :one
INSERT INTO "tags" ("id", "tag")
VALUES ($1, $2)
RETURNING *;

-- name: SelectTag :one
SELECT *
FROM "tags"
WHERE "id" = $1
LIMIT 1;

-- name: SelectTags :many
SELECT *
FROM "tags"
ORDER BY "tag";

-- name: DeleteTag :exec
DELETE FROM "tags"
WHERE "id" = $1;