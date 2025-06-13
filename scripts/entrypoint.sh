#!/bin/sh

set -e

echo "Waiting for Postgres to be ready..."
until pg_isready -h $DB_HOST -p $DB_PORT -U $DB_USER; do
  sleep 2
done

echo "Postgres is ready. Running migrations..."

DB_PASSWORD=$(cat /run/secrets/db_password)

GOOSE_DBSTRING="postgres://$DB_USER:$DB_PASSWORD@$DB_HOST:$DB_PORT/$DB_NAME?sslmode=disable"

export GOOSE_DRIVER=postgres
export GOOSE_DBSTRING=$GOOSE_DBSTRING

GOOSE_MIGRATION_DIR=/bin/app/sql/migrations goose up

echo "Starting application..."
exec /bin/app/ocrolus-task
