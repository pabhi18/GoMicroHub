FROM golang:1.23-alpine AS builder

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . . 

RUN GOOS=linux GOARCH=amd64 go build -o authApp ./cmd/auth

FROM alpine:latest

WORKDIR /app

COPY --from=builder /app/authApp /app

RUN chmod +x /app/authApp

CMD ["/app/authApp"]
