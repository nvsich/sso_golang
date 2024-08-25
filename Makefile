.PHONY: migration-up
migration-up:
	go run ./cmd/migrator --storage-path=./storage/sso.db --migrations-path=./migrations