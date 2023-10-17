.DEFAULT_GOAL := build

build_dir = ./bin

.PHONY = build run

build:
	@go build -o $(build_dir)/todo ./cmd/todo

run: build
	@exec $(build_dir)/todo
