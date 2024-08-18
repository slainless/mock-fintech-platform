#!/usr/bin/env bash

export OS="$(go env GOOS)"
export ARCH="$(go env GOARCH)"

mkdir bin/migrate
curl -L https://github.com/golang-migrate/migrate/releases/download/v4.17.1/migrate.$OS-$ARCH.tar.gz | \
  tar xvz -C bin/migrate