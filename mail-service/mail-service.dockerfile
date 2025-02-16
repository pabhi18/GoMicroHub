FROM golang:1.23-alpine AS builder

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . . 

RUN GOOS=linux GOARCH=amd64 go build -o mailApp ./cmd/api

FROM alpine:latest

WORKDIR /app

COPY --from=builder /app/mailApp /app
COPY ./cmd/templates ./templates

RUN chmod +x /app/mailApp

CMD ["/app/mailApp"]


