# Build stage
FROM golang:1.25-alpine AS builder
WORKDIR /app
COPY . .
RUN go build -o main main.go
RUN apk add --no-cache curl && \
    mkdir -p /app/migration && \
    curl -L https://github.com/golang-migrate/migrate/releases/download/v4.15.2/migrate.linux-amd64.tar.gz | tar xz -C /app/migration && \
    chmod +x /app/migration/migrate

# Run stage
FROM alpine:3.19
WORKDIR /app
COPY --from=builder /app/main .
COPY config.json .
COPY db/migration ./db/migration
COPY start.sh .
COPY --from=builder /app/migration/migrate ./migration/migrate
RUN chmod +x /app/start.sh

EXPOSE 8091
ENTRYPOINT ["/app/start.sh"]