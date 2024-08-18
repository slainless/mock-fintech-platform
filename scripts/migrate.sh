#!/usr/bin/env bash
bin/migrate/migrate -database "${POSTGRESQL_URL}" -path db/migrations "$@"