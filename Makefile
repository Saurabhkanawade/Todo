NAME :=to

VERSION := v1.0.0

.PHONY: setup
setup:  ## Installs all dependencies
	go mod tidy -compat=1.21
	which sqlboiler || (go install github.com/volatiletech/sqlboiler/v4@v4.14.2 && go install github.com/volatiletech/sqlboiler/v4/drivers/sqlboiler-psql@v4.14.2)


.PHONY: test
test: ## runs test with coverage (doesn't produce a report)
	go test -v -coverpkg=./... -coverprofile=bin/coverage.txt ./...

.PHONY: test-no-cover
test-no-cover: ## runs test without coverage
	go test -v ./...

.PHONY: cover
cover: test ## Run all the tests and opens the coverage report
		go tool cover -html=bin/coverage.txt

install-mockery: ## installs mockery
	which mockery || go get github.com/vektra/mockery/v2/.../

mocks: ## regenerates mockery mocks
	mockery --all --outpkg=mocks --output=./mocks --inpackage --dir=./internal/

install-lint: ## installs golangci-lint linter
	which golangci-lint || go install github.com/golangci/golangci-lint/cmd/golangci-lint@v1.60.3

lint: ## runs golangci-lint linter
	golangci-lint run --config golangci.yml --timeout 6m0s

swagger-install: ## install go swagger tool
	which swagger || go install github.com/go-swagger/go-swagger/cmd/swagger@v0.31.0

swagger-generate: ## generate swagger.json file
	swagger generate spec -o ./swagger-ui/swagger.json --scan-models

swagger-validate: swagger-generate ## validate the swagger docs
	swagger validate ./swagger-ui/swagger.json

swagger-serve: swagger-generate ## host local swagger page for debugging
	swagger serve ./swagger-ui/swagger.json -F swagger

.PHONY: pre-commit
pre-commit: test-no-cover lint swagger-validate    ## call before committing and pushing changes
	@echo 'pre-commit checks completed without error. You are clear to commit to your branch'

.PHONY: help
help:  ## displays this message
	@grep -E '^[ a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | \
		awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-15s\033[0m %s\n", $$1, $$2}'

.PHONY: version
version:  ## displays the version
	@echo $(VERSION)

run: ## run the web server locally
	 go build -o bin/ && ./bin/$(NAME)

bench:  ## run all benchmark tests
	go test -v ./... -bench=. -run=xxx -benchmem

.PHONY: code-gen-db
code-gen-db:  ## generates code from db/sqlboiler.toml
	sqlboiler psql
	grep -rl "DeletedAt" internal/dbmodels | xargs sed -i "" 's/DeletedAt/DeletedTS/g'  # fixes sqlboiler bug
	grep -rl "deleted_at" internal/dbmodels | xargs sed -i "" 's/deleted_at/deleted_ts/g'

.PHONY: vendor
vendor: ## run go mod tidy and vendor
	go mod tidy -compat=1.23 && go mod vendor
