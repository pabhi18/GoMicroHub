FROM golang:1.23-alpine AS builder

WORKDIR /app

COPY . .

RUN GOOS=linux GOARCH=amd64 go build -o mailApp .

FROM alpine:latest

WORKDIR /app

COPY --from=builder /app/mailApp /app
COPY ./cmd/templates ./templates

RUN chmod +x /app/mailApp

CMD ["/app/mailApp"]

