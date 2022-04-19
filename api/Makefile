include .env

postgresup:
	podman run --name $(PODMAN_NAME) -p 5432:5432 -e POSTGRES_USER=$(POSTGRES_USER) -e POSTGRES_PASSWORD=$(POSTGRES_PASSWORD) -v "/home/karan/Secondary/go_post_pod:/var/lib/postgresql/data:Z" -dt postgres:latest

postgresdown:
	podman stop $(PODMAN_NAME) && podman rm $(PODMAN_NAME) && sudo rm -rf ~/Secondary/go_post_pod && mkdir ~/Secondary/go_post_pod

createdb:
	podman exec -it $(PODMAN_NAME) createdb --username=$(POSTGRES_USER) --owner=$(POSTGRES_USER) ddgf_db_dev

dropdb:
	podman exec -it $(PODMAN_NAME) dropdb ddgf_db_dev

migrateup:
	migrate -path db/migrations -database $(POSTGRES_STRING) -verbose up

migratedown:
	migrate -path db/migrations -database $(POSTGRES_STRING) -verbose down

migratenew:
	read -p "migration: " migration ; migrate create -ext sql -dir db/migrations -seq $$migration

.PHONY: postgres createdb dropdb migrateup migratedown postgresdown migratenew
