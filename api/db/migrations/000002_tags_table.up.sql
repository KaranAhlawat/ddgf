CREATE TABLE "tags" (
    "id" UUID PRIMARY KEY,
    "tag" VARCHAR(50) NOT NULL UNIQUE
);

CREATE INDEX ON "tags" ("tag");