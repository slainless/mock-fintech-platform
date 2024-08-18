#!/usr/bin/env bash
mkdir .data/postgres
chmod 777 .data/postgres
docker run \
  --rm \
  --name mock_fintech_postgres \
  -e POSTGRES_PASSWORD=1234 \
  -e POSTGRES_DB=mock_fintech \
  -p 7777:5432 \
  -p 7776:80 \
  -v $(pwd)/.data/postgres:/var/lib/postgresql/data \
  -d \
  postgres