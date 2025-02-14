FROM golang:1.23-alpine AS builder

WORKDIR /app

COPY . .

RUN GOOS=linux GOARCH=amd64 go build -o brokerApp ./cmd/api

FROM alpine:latest

WORKDIR /app

COPY --from=builder /app/brokerApp /app

RUN chmod +x /app/brokerApp

ENTRYPOINT ["/app/brokerApp"]


