.PHONY: run
run:
	rm -rf cover
	docker-compose down --remove-orphans --volumes
	docker-compose up --force-recreate