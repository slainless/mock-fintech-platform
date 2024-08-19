#!/usr/bin/env bash
mkdir bin/swag
GOBIN="$(pwd)/bin/swag" go install github.com/swaggo/swag/cmd/swag@latest