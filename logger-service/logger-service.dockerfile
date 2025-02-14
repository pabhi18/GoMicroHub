FROM golang:1.23-alpine AS builder

WORKDIR /app

COPY . .

RUN GOOS=linux GOARCH=amd64 go build -o loggerApp .

FROM alpine:latest

WORKDIR /app

COPY --from=builder /app/loggerApp /app

RUN chmod +x /app/loggerApp

CMD ["/app/loggerApp"]



