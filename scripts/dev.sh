#!/usr/bin/env bash
echo "Starting dev environment"

destroy() {
  echo "Cleaning dev containers"
  docker stop mock_fintech_pgadmin mock_fintech_postgres
  # in case the container is persisted...
  docker rm mock_fintech_pgadmin mock_fintech_postgres
}

destroy
echo "Running postgres container"
./scripts/run_postgres.sh
echo "Running pgadmin container"
./scripts/run_pgadmin.sh

cleanup() {
  echo "SIGINT received, cleaning up..."
  destroy
  exit 0
}

trap cleanup SIGINT
echo "Press CTRL+C to stop"
while true; do
  read
done