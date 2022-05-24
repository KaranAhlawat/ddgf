CREATE TABLE "advices_tags" (
    "advice_id" uuid NOT NULL,
    "tag_id" uuid NOT NULL
);

ALTER TABLE "advices_tags"
    ADD CONSTRAINT "pk_advices_tags" PRIMARY KEY ("advice_id", "tag_id"),
    ADD CONSTRAINT "fk_advice_id" FOREIGN KEY ("advice_id") REFERENCES "advices" ("id") ON UPDATE CASCADE ON DELETE CASCADE,
    ADD CONSTRAINT "fk_tag_id" FOREIGN KEY ("tag_id") REFERENCES "tags" ("id") ON UPDATE CASCADE ON DELETE CASCADE;

