
HOST="$DB_HOST"
PORT="$DB_PORT"

echo "Waiting for $HOST:$PORT to be available..."
while ! nc -z "$HOST" "$PORT"; do
  echo "Waiting for $HOST:$PORT..."
  sleep 1
done

echo "✅ MySQL is up at $HOST:$PORT. Starting the app..."

exec "$@"
