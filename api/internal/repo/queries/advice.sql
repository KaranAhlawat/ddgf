-- name: InsertAdvice :one
INSERT INTO "advices" ("id", "content")
VALUES ($1, $2)
RETURNING *;

-- name: SelectAdvice :one
SELECT *
FROM "advices"
WHERE "id" = $1
LIMIT 1;

-- name: SelectAdvices :many
SELECT *
FROM "advices"
ORDER BY "id";

-- name: UpdateAdvice :exec
UPDATE "advices"
SET "content" = $1
WHERE "id" = $2;

-- name: DeleteAdvice :exec
DELETE FROM "advices"
WHERE "id" = $1;
