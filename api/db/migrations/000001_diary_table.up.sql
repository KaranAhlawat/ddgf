CREATE TABLE "diary" (
    "page_id" UUID PRIMARY KEY,
    "datetime" DATE NOT NULL,
    "content" TEXT NOT NULL
);

CREATE INDEX ON "diary" ("datetime");