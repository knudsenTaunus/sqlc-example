SQLC_VERSION:=1.10.0

generate:
	docker run -v $$(pwd)/datasource:/srv -u $$(id -u):$$(id -g) -w /srv kjconroy/sqlc:$(SQLC_VERSION) generate

init:
	docker run -v $$(pwd)/datasource:/srv -u $$(id -u):$$(id -g) -w /srv kjconroy/sqlc:$(SQLC_VERSION) init