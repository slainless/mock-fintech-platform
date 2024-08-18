#!/usr/bin/env bash
bin/jet/jet -schema=public -dsn="${POSTGRESQL_URL}" -path=./pkg/internal/artifact/database "$@"