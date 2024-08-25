.PHONY: migration-up
migration-up:
	go run ./cmd/migrator --storage-path=./storage/sso.db --migrations-path=./migrations

.PHONY: run-local
run-local:
	go run ./cmd/sso --config=./config/local.yaml


# tests
.PHONY: migration-up-test
migration-up-test:
	go run ./cmd/migrator/main.go --storage-path=./storage/sso.db --migrations-path=./tests/migrations --migrations-table=migrations_test