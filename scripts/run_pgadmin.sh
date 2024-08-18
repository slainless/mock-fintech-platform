#!/usr/bin/env bash
mkdir .data/pgadmin
chmod 777 .data/pgadmin
docker run \
  --rm \
  --name mock_fintech_pgadmin \
  --network container:mock_fintech_postgres \
  -e PGADMIN_DEFAULT_EMAIL="user@concreteai.io" \
  -e PGADMIN_DEFAULT_PASSWORD="1234" \
  -v $(pwd)/.data/pgadmin:/var/lib/pgadmin \
  -d \
  dpage/pgadmin4
