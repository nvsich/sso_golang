.PHONY: migration-up
migration-up:
	go run ./cmd/migrator --storage-path=./storage/sso.db --migrations-path=./migrations

.PHONY: run-local
run-local:
	go run ./cmd/sso --config=./config/local.yaml