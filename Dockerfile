# Build stage
FROM golang:1.25-alpine AS builder
WORKDIR /app
COPY . .
RUN go build -o main main.go

# Run stage
FROM alpine:3.19
WORKDIR /app
COPY --from=builder /app/main .
COPY config.json .

EXPOSE 8091
CMD ["/app/main"]