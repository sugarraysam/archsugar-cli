#!/usr/bin/env bash
#
# Generate all protobuf bindings.
# Run from repository root.
set -e
set -u

URL="https://github.com/golangci/golangci-lint"

if [[ ! "$(golangci-lint --version 2>/dev/null)" ]]; then
    echo 2>"Please install a golangci-lint version compatible with ${GOLANGCI_LINT_VERSION}"
    echo 2>"Follow insttructions here: ${URL}"
    exit 1
fi

golangci-lint run
