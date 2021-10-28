include .bingo/Variables.mk
FILES_TO_FMT  ?= $(shell find . -path ./vendor -prune -o -name '*.go' -print)
FILES_TO_TEST ?= $(shell go list ./... | grep -v /vendor/)

help: ## Displays help.
	@awk 'BEGIN {FS = ":.*##"; printf "\nUsage:\n  make \033[36m<target>\033[0m\n\nTargets:\n"} /^[a-z0-9A-Z_-]+:.*?##/ { printf "  \033[36m%-10s\033[0m %s\n", $$1, $$2 }' $(MAKEFILE_LIST)

.PHONY: format
format: ## Runs all format targets.
format: go-format

.PHONY: go-format
go-format: ## Formats Go code including imports.
go-format: $(GOIMPORTS)
	@echo ">> formatting go code"
	@gofmt -s -w $(FILES_TO_FMT)
	@$(GOIMPORTS) -w $(FILES_TO_FMT)

.PHONY: lint
lint: ## Runs all lint targets.
lint: go-lint

.PHONY: go-lint
go-lint: ## Lints Go code.
go-lint: $(GOLANGCI_LINT)
	@echo ">> linting all of the Go files"
	@$(GOLANGCI_LINT) run

.PHONY: bench
bench: ## Benchmarks Go code.
bench: $(BENCHSTAT)
	@echo ">> running benchmarks"
	@mkdir -p .test
	@if [ -f .test/new.txt ]; then mv .test/new.txt .test/old.txt; fi
	@$(GO) test -bench=. -benchmem -count=3 $(FILES_TO_TEST) | tee .test/new.txt
	@$(BENCHSTAT) .test/old.txt .test/new.txt

.PHONY: test
test: ## Tests Go code.
test:
	@echo ">> running unit tests"
	@$(GO) test $(FILES_TO_TEST);

.PHONY: go-build
go-build: ## Creates a new Go build.
go-build: $(GORELEASER)
	$(GORELEASER) build --snapshot --rm-dist

.PHONY: go-release
go-release: ## Creates a new Go release.
go-release: $(GORELEASER)
	$(GORELEASER) release --snapshot --rm-dist
