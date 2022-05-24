CREATE TABLE "pages" (
    "id" uuid PRIMARY KEY,
    "datetime" date NOT NULL,
    "content" text NOT NULL
);

CREATE INDEX ON "pages" ("datetime");

