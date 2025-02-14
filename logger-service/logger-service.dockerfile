FROM golang:1.23-alpine AS builder

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . . 

RUN GOOS=linux GOARCH=amd64 go build -o loggerApp ./cmd/logger

FROM alpine:latest

WORKDIR /app

COPY --from=builder /app/loggerApp /app

RUN chmod +x /app/loggerApp

CMD ["/app/loggerApp"]




