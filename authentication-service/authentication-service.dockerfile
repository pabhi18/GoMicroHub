FROM golang:1.23-alpine AS builder

WORKDIR /app

COPY . .

RUN GOOS=linux GOARCH=amd64 go build -o authApp .

FROM alpine:latest

WORKDIR /app

COPY --from=builder /app/authApp /app

RUN chmod +x /app/authApp

CMD ["/app/authApp"]
