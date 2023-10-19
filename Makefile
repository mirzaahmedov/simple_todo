.DEFAULT_GOAL := build

build_dir = ./bin
migration_dir = ./migrations
database_url = postgresql://localhost:5432/simple_todo?user=postgres&sslmode=disable

.PHONY = build run

build:
	@go build -o $(build_dir)/todo ./cmd/todo

run: build
	@exec $(build_dir)/todo

migration_init:
	@migrate create -dir $(migration_dir) -ext sql -seq init

migration_up:
	@migrate -path $(migration_dir) -database $(database_url) up

migration_down:
	@migrate -path $(migration_dir) -database $(database_url) down
