export PATH := $(PATH):$(shell go env GOPATH)/bin

.PHONY: all
all: gomod generate build

.PHONY: clean-gen
clean-gen:
	rm -rf gen

.PHONY: generate
generate: clean-gen
	mkdir -p gen/go
	go generate ./...

.PHONY: build
build:
	go build -o bin/pim-sys cmd/main.go

docker-%:
	docker run --rm -v $$PWD:$$PWD:rw -w $$PWD main-image make $*

.PHONY: bash
docker-bash:
	docker run --rm -it main-image bash

.PHONY: build-docker-image
build-docker-image:
	docker build -t main-image:latest .

.PHONY: gomod
gomod: 
	go mod tidy
