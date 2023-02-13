
.PHONY: fetch
fetch:
	go run ./cmd/coinfetcher/main.go

.PHONY: gofmt
gofmt: ## Run gofumpt over the codebase. gofumpt must be installed and in your path.
	gofumpt -l -w .

.PHONY: lint
lint: ## Run golangci-lint. golangci-lint must be installed and in your path.
	golangci-lint run --modules-download-mode vendor

.PHONY: test
test:
	go test -count=1 -cover ./...

.PHONY: docker_up
docker_up:
	docker-compose -f docker-compose.yml up -d

.PHONY: dbc
dbc: ## Connect to local MySQL database.
	dev/local_database_conn.sh

.PHONY: migrate
migrate:
	go run cmd/migrate/migrate.go -action up
