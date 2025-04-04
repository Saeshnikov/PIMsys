.PHONY: all
all: gomod generate build

.PHONY: clean-gen
clean-gen:
	rm -rf gen
	rm -rf docs/gen

.PHONY: generate
generate: clean-gen
	mkdir -p gen/go
	go generate ./proto/...
	go generate ./internal/...
	go generate .

.PHONY: build
build: cmd/*
	@for file in $^ ; do \
                srvsrv="`echo $$file | cut -f 2 -d '/'`"; \
				echo cmd/$$srvsrv/main.go;\
				go build -cover -o bin/$$srvsrv cmd/$$srvsrv/main.go; \
    done

docker-%:
	docker run --rm -v $$PWD:$$PWD:rw -w $$PWD main-image make $*

.PHONY: bash
docker-bash:
	docker run --rm -v $$PWD:$$PWD:rw -w $$PWD -it main-image bash

.PHONY: build-docker-images
build-docker-images:
	docker build -f docker/main-image -t main-image:latest .

.PHONY: gomod
gomod: 
	go mod tidy

test-%:
	rm -rf cover
	docker-compose down --remove-orphans --volumes
	docker-compose --profile $* up --force-recreate --abort-on-container-exit --exit-code-from test-$*
	docker-compose down --remove-orphans --volumes
	
.PHONY: cover
docker-cover:
	docker run --rm -v $$PWD:$$PWD:rw -w $$PWD main-image:latest mkdir -p /cover_merged && \
		go tool covdata textfmt -i=cover -o cover.out && \
    	go tool cover -html cover.out -o cover.html
