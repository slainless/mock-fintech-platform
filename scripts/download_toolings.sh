#!/usr/bin/env bash

export os="$(go env GOOS)"
export arch="$(go env GOARCH)"

mkdir bin/migrate
curl -L https://github.com/golang-migrate/migrate/releases/download/v4.17.1/migrate.$os-$arch.tar.gz | \
  tar xvz -C bin/migrate
