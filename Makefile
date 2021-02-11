GO        = go
GO_BUILD  = $(GO) build
GO_FORMAT = $(GO) fmt
GO_LINT   = golangci-lint
GO_TEST   = $(GO) test
GO_VET    = $(GO) vet

BINARY_PATH = "./bin/kumade"

ifeq (, $(shell which golangci-lint))
	$(warning "could not find golangci-lint in $(PATH), run: curl -sfL https://install.goreleaser.com/github.com/golangci/golangci-lint.sh | sh")
endif

.PHONY: clean fmt lint test vet

default: all
all: fmt lint vet test build

build:
	$(GO_BUILD) -o $(BINARY_PATH) -v
clean:
	rm $(BINARY_PATH)
fmt:
	$(GO_FORMAT) ./...
lint:
	$(GO_LINT) run -v
test:
	$(GO_TEST) -v ./...
vet:
	$(GO_VET) ./...
