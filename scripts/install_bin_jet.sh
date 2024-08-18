#!/usr/bin/env bash
mkdir bin/jet
GOBIN="$(pwd)/bin/jet" go install github.com/go-jet/jet/v2/cmd/jet@latest