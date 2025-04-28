.PHONY: run
run:
	rm -rf cover
	docker-compose down --remove-orphans --volumes
	docker-compose up --build --force-recreate

.PHONY: build
build:
	docker-compose build
