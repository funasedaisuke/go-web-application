.PHONY:  help build build-local up down logs ps test
.DEFAULT_GOAL := help

DOCKER_TAG := latest
build:
    dokcer build -t gowebapp:${DOCKER_TAGA} --target deploy ./
build-local:
	docker compose build --no-cache
up:
	docker compose up -d
down:
	docker compose down
logs:
	docker compose logs -f
ps:
	docker compose ps
test:
	go test -race -shuffle=on ./...
migrate:
	mysqldef -u todo -p todo -h localhost -P 33306 todo < ./_tools/mysql/schema.sql
dry-migrate:
	mysqldef -u todo -p todo -h localhost -P 33306 todo --dry-run < ./_tools/mysql/schema.sql

generate:
	go generate ./...

help:
    @grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) |\
                awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-20s\033[0m %s\n",$$1,$$$2"}