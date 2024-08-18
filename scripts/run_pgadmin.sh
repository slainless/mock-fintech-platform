#!/usr/bin/env bash
mkdir var/pgadmin
chmod 777 var/pgadmin
docker run \
  --rm \
  -p 7776:80 \
  -e PGADMIN_DEFAULT_EMAIL="user@concreteai.io" \
  -e PGADMIN_DEFAULT_PASSWORD="1234" \
  -v $(pwd)/var/pgadmin:/var/lib/pgadmin \
  -d \
  dpage/pgadmin4
