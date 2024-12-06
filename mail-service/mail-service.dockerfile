FROM alpine:latest

WORKDIR /app

COPY mailApp .
COPY ./cmd/templates ./templates

CMD [ "/app/mailApp" ]
