#!/bin/sh

# Run database migrations
/app/main migrate

# Start the server
exec /app/main