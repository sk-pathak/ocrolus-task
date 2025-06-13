#!/bin/sh

set -e

echo "Waiting for Postgres to be ready..."
until pg_isready -h $DB_HOST -p $DB_PORT -U $DB_USER; do
  sleep 2
done

echo "Postgres is ready. Running migrations..."

# Read the password from the secrets file
DB_PASSWORD=$(cat /run/secrets/db_password)

# Construct the database connection string using environment variables
GOOSE_DBSTRING="postgres://$DB_USER:$DB_PASSWORD@$DB_HOST:$DB_PORT/$DB_NAME?sslmode=disable"

# Export the connection string for goose
export GOOSE_DRIVER=postgres
export GOOSE_DBSTRING=$GOOSE_DBSTRING

# Run migrations
GOOSE_MIGRATION_DIR=/bin/app/sql/migrations goose up

echo "Starting application..."
exec /bin/app/ocrolus-task