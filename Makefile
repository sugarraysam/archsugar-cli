TARGETS := lint test build install uninstall clean
.PHONY: $(TARGETS)

export SHELL := /bin/bash
export GO111MODULE := on
export GOOS := linux
export GOARCH := amd64
export GOLANGCI_LINT_VERSION := 1.36.0
export BINARY := archsugar
export PROJECT := github.com/sugarraysam/archsugar-cli
export VERSION := 1.0.0
export COMMIT := $(shell git rev-parse --short HEAD)
export BUILD_TIME := $(shell date -u +'%Y-%m-%dT%H:%M:%S%Z')
export LDFLAGS := -s -w -X $(PROJECT)/version.Version=$(VERSION) \
				 -X $(PROJECT)/version.Commit=$(COMMIT) \
				 -X $(PROJECT)/version.BuildTime=$(BUILD_TIME)

lint:
	@golangci-lint run

test:
	@go clean -testcache
	@go test -covermode=count -coverprofile=.coverage.out -tags integration ./...
	@go test -race >/dev/null 2>&1

build:
	@CGO_ENABLED=0 go build -a -installsuffix cgo -o _build/$(BINARY) -ldflags "$(LDFLAGS)"

install: build
	@sudo install -Dm 755 _build/$(BINARY) /usr/bin/$(BINARY)
	@/usr/bin/$(BINARY) completion | sudo install -Dm 644 /dev/stdin /usr/share/zsh/site-functions/_$(BINARY)

uninstall:
	-@sudo rm -f /usr/bin/$(BINARY) /usr/share/zsh/site-functions/_$(BINARY) > /dev/null 2>&1

FILES_TO_CLEAN := $(shell find . -type d -name _build)
clean:
	@echo "Cleaning files: $(FILES_TO_CLEAN)"
	@rm -fr $(FILES_TO_CLEAN)
	@go mod tidy
