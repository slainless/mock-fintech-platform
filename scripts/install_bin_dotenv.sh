#!/usr/bin/env bash
mkdir bin/dotenv
GOBIN="$(pwd)/bin/dotenv" go install github.com/joho/godotenv/cmd/godotenv@latest