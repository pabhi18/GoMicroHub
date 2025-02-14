FROM golang:1.23-alpine AS builder

WORKDIR /app

COPY . .

RUN GOOS=linux GOARCH=amd64 go build -o listenApp .

FROM alpine:latest

WORKDIR /app

COPY --from=builder /app/listenApp /app

RUN chmod +x /app/listenApp

CMD ["/app/listenApp"]
