#!/bin/sh
set -e

host="$1"
shift
until pg_isready -h "$host" -p 5432 -U "$POSTGRES_USER"; do
  echo "Esperando a Postgres en $host..."
  sleep 2
done

exec "$@"
