-- name: InsertAdviceTagEntry :one
INSERT INTO "advices_tags" ("advice_id", "tag_id")
    VALUES ($1, $2)
RETURNING
    *;

-- name: SelectAllEntries :many
SELECT
    "at"."advice_id", "at"."tag_id", "t"."tag"
FROM
    "advices_tags" "at"
    JOIN "tags" "t" ON "at"."tag_id" = "t"."tag"
ORDER BY
    "advice_id";

-- name: SelectTagsForAdvice :many
SELECT
    "at"."tag_id",
    "t"."tag"
FROM
    "advices_tags" "at"
    JOIN "tags" "t" ON "at"."tag_id" = "t"."id"
WHERE
    "at"."advice_id" = $1;

-- name: SelectAdvicesForTag :many
SELECT
    "at"."advice_id",
    "a"."content"
FROM
    "advices_tags" "at"
    JOIN "advices" "a" ON "at"."advice_id" = "a"."id"
WHERE
    "at"."tag_id" = $1;

-- name: DeleteTagFromAdvice :exec
DELETE FROM "advices_tags"
WHERE "advice_id" = $1
    AND "tag_id" = $2;

