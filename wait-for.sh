#!/bin/sh

set -e

host="$1"
shift
cmd="$@"

until nc -z -v -w30 "$host" 2>/dev/null; do
  echo "⏳ Waiting for $host to be available..."
  sleep 3
done

echo "✅ $host is available! Running the application..."
exec $cmd
