FROM golang:1.23-alpine AS builder

WORKDIR /app

# Copy only go.mod first
COPY go.mod ./
# Initialize go.mod and download dependencies
RUN go mod download || true
RUN go mod tidy

COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -o frontApp ./cmd/web

FROM alpine:latest

WORKDIR /app

COPY --from=builder /app/frontApp /app
COPY ./cmd/web/templates ./templates

RUN chmod +x /app/frontApp

EXPOSE 9090

CMD ["/app/frontApp"] 