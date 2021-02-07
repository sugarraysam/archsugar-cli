export SHELL := /bin/bash
export GO111MODULE := on
export GOOS := linux
export GOARCH := amd64
export GOLANGCI_LINT_VERSION := 1.36.0


TEST_TARGETS := deps lint test sonar pre-commit
.PHONY: $(TEST_TARGETS)

deps:
	@go mod tidy
	@go mod verify

lint:
	@./scripts/golangci-lint.sh

test:
	@go test -covermode=count -coverprofile=.coverage.out ./...
	@go tool cover -func=.coverage.out
	@go test -race >/dev/null 2>&1

sonar:
	@./scripts/sonarqube.sh

pre-commit: $(TEST_TARGETS)


BUILD_TARGETS := build install uninstall clean
.PHONY: $(BUILD_TARGETS)

BINARY := archsugar
PROJECT := github.com/sugarraysam/archsugar
VERSION := 1.0.0
COMMIT := $(shell git rev-parse --short HEAD)
BUILD_TIME := $(shell date -u +'%Y-%m-%dT%H:%M:%S%Z')
LDFLAGS := -s -w -X $(PROJECT)/version.Version=$(VERSION) \
				 -X $(PROJECT)/version.Commit=$(COMMIT) \
				 -X $(PROJECT)/version.BuildTime=$(BUILD_TIME)

build: deps
	@CGO_ENABLED=0 go build -a -installsuffix cgo -o _build/$(BINARY) -ldflags "$(LDFLAGS)"

install: build
	@sudo install -Dm 755 _build/$(BINARY) /usr/bin/$(BINARY)
	@/usr/bin/$(BINARY) completion | sudo install -Dm 644 /dev/stdin /usr/share/zsh/site-functions/_$(BINARY)

uninstall:
	-@sudo rm -f /usr/bin/$(BINARY) /usr/share/zsh/site-functions/_$(BINARY) > /dev/null 2>&1

FILES_TO_CLEAN := $(shell find . -type d -name _build)
clean:
	@echo "Cleaning files: $(FILES_TO_CLEAN)"
	-@rm -fr $(FILES_TO_CLEAN)
	-@go clean -testcache
	-@docker rm -f $$(docker ps -aq) > /dev/null 2>&1
