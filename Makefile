.PHONY: build
build: db_up
	# go build -v ./cmd/client
	go build -v ./cmd/server

.DEFAULT_GOAL := build