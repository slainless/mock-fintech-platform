#!/usr/bin/env bash
echo "Starting dev environment"

echo "Running postgres container"
POSTGRES_CONTAINER="$(./scripts/run_postgres.sh)"
echo "Postgres container = ${POSTGRES_CONTAINER}"

echo "Running pgadmin container"
PGADMIN_CONTAINER="$(./scripts/run_pgadmin.sh)"
echo "pgadmin container = ${PGADMIN_CONTAINER}"

cleanup() {
  echo "SIGINT received, cleaning up..."
  docker stop ${POSTGRES_CONTAINER} ${PGADMIN_CONTAINER}
  # in case the container is persisted...
  docker rm ${POSTGRES_CONTAINER} ${PGADMIN_CONTAINER}
  exit 0
}

trap cleanup SIGINT
echo "Press CTRL+C to stop"
while true; do
  read
done