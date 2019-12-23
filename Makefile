GO    					:= go
PROJECT_NAME            ?= jotnar
TEST                    ?= $(shell go list ./... | grep -v '/vendor/')
TESTARGS                ?= -v -race

.PHONY: setup
setup:
	@echo ">> installing dependencies"
	@sudo curl -sfL https://raw.githubusercontent.com/golangci/golangci-lint/master/install.sh| sh -s -- -b $(go env GOPATH)/bin v1.21.0
	@$(GO) get -u "github.com/golang/tools/cmd/goimports"
	@$(GO) mod tidy

.PHONY: test
test:
	@echo ">> running tests"
	@$(GO) test $(TEST) $(TESTARGS)

.PHONY: fmt
fmt:
	@find . -name "*.go" | xargs goimports -w
	@find . -name "*.go" | xargs gofmt -w

.PHONY: lint
lint:
	@echo ">> linting code"
	@golangci-lint run

.PHONY: clean
clean:
	@echo ">> clean project"
	@rm -rf example *.o
