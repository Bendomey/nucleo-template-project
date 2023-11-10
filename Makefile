run-dev: ## install reflex(go package) and run dev server
	scripts/run-dev.sh

run:
	scripts/run.sh

utest: ## run unit tests (does not require mock dataset)
	go test -count=1 ./test/unit

itest: ## run integration tests (requires mock dataset to be running)
	go test -count=1 ./test/integration

lint: ## run linter
	golangci-lint run ./...