.PHONY: run
run:
	rm -rf cover
	docker-compose down --remove-orphans --volumes
	docker-compose up --build --force-recreate

.PHONY: build
build: build-ui

.PHONY: build-ui
build-ui:
	docker build -t ui ./frontend/pim-sys
