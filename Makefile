TARGETS := lint test build install uninstall clean
.PHONY: $(TARGETS)

export SHELL := /bin/bash
export GO111MODULE := on

BINARY := archsugar

lint:
	@golangci-lint run

test:
	@go clean -testcache
	@go test -cover -tags integration ./...
	@go test -race >/dev/null 2>&1

build:
	@goreleaser release --skip-publish --snapshot --rm-dist

install: build
	@sudo install -Dm 755 dist/archsugar-cli_linux_amd64/$(BINARY) /usr/local/bin/$(BINARY)
	@/usr/local/bin/$(BINARY) completion | sudo install -Dm 644 /dev/stdin /usr/share/zsh/site-functions/_$(BINARY)

uninstall:
	-@sudo rm -f /usr/local/bin/$(BINARY) /usr/share/zsh/site-functions/_$(BINARY) > /dev/null 2>&1

FILES_TO_CLEAN := $(shell find . -type d -name dist)
clean:
	@echo "Cleaning files: $(FILES_TO_CLEAN)"
	@rm -fr $(FILES_TO_CLEAN)
	@go mod tidy
