.PHONY: default build
APPS        := healthchecker
BLDDIR      ?= bin

.EXPORT_ALL_VARIABLES:
GO111MODULE  = on

default: build

build: clean $(APPS)

$(BLDDIR)/%:
	go build $(LDFLAGS) -o $@ ./cmd/$*

$(APPS): %: $(BLDDIR)/%

deps:
	@echo 'Installing go modules...'
	@go mod download

clean:
	@mkdir -p $(BLDDIR)
	@for app in $(APPS) ; do \
		rm -f $(BLDDIR)/$$app ; \
	done

format:
	@echo 'Formatting the code...'
	@gofmt -w .
	@goimports -local "github.com/teneta-io/healthchecker" -w .

lint-revive:
	@echo 'Linting with revive...'
	@revive -formatter stylish -config=revive.toml ./...

codegen:
	@echo 'Generating mocks...'
	@go generate ./...

lint-golangci: 
	@echo 'Linting with golangci...'
	@golangci-lint run ./pkg/...

lint: format lint-revive lint-golangci
