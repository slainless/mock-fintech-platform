#!/usr/bin/env bash
echo $(git describe --tags --exact-match 2> /dev/null || echo "v0.0.0") > $1
