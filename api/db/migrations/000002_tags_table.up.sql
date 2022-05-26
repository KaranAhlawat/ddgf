CREATE TABLE "tags" (
    "id" uuid PRIMARY KEY,
    "tag" varchar(50) NOT NULL UNIQUE
);

CREATE INDEX ON "tags" ("tag");

