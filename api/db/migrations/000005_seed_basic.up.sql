INSERT INTO "pages" ("id", "content", "datetime")
    VALUES ('c3820a56-b9d9-46c2-9c69-d98883b53eb7', 'Page One : Lorem Ipsum Set Dolorfjalsdfja', '2022-05-25'), ('6dba00a4-762f-455c-a185-32950c7f514f', 'Page Two : Lorem Ipsum Set Dolor', '2022-05-20');

INSERT INTO "advices" ("id", "content")
    VALUES ('5dd2b406-b7e4-4291-9add-41abcae06a00', 'Always test your code'), ('d721a76f-43aa-40af-bb67-ed5595cff192', 'It takes time, but it is worth it');

INSERT INTO "tags" ("id", "tag")
    VALUES ('524bb97d-f78f-41dc-842b-728aee1e2ad7', 'default'), ('e5ba43e5-8d00-4a0b-958c-f03561c102f7', 'custom');

INSERT INTO "advices_tags" ("advice_id", "tag_id")
    VALUES ('5dd2b406-b7e4-4291-9add-41abcae06a00', '524bb97d-f78f-41dc-842b-728aee1e2ad7'), ('5dd2b406-b7e4-4291-9add-41abcae06a00', 'e5ba43e5-8d00-4a0b-958c-f03561c102f7'), ('d721a76f-43aa-40af-bb67-ed5595cff192', 'e5ba43e5-8d00-4a0b-958c-f03561c102f7');

