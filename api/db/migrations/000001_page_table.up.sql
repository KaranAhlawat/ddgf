CREATE TABLE "pages" (
    "id" UUID PRIMARY KEY,
    "datetime" DATE NOT NULL,
    "content" TEXT NOT NULL
);

CREATE INDEX ON "pages" ("datetime");
