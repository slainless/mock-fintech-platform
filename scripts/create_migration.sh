#!/usr/bin/env bash
bin/migrate/migrate create -ext sql -dir db/migrations "$@"