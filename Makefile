.PHONY: run
run:
	rm -rf cover
	docker-compose down --remove-orphans --volumes
	docker-compose up --build --force-recreate

.PHONY: build
build:
	docker-compose build

.PHONY: test
test:
	rm -rf cover
	docker-compose down --remove-orphans --volumes
	docker-compose --profile test up --force-recreate --exit-code-from tests \
		--no-attach migrations --no-attach sso-envoy --no-attach shop-envoy \
		--no-attach branch-envoy --no-attach products-envoy --no-attach template-envoy \
		--no-attach logs-envoy