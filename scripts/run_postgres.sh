#!/usr/bin/env bash
mkdir var/postgres
chmod 777 var/postgres
docker run \
  --rm \
  --name mock_fintech_postgres \
  -e POSTGRES_PASSWORD=1234 \
  -e POSTGRES_DB=mock_fintech \
  -p 7777:5432 \
  -p 7776:80 \
  -v $(pwd)/var/postgres:/var/lib/postgresql/data \
  -d \
  postgres