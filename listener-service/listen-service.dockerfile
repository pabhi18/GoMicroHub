FROM alpine:latest

WORKDIR /app

COPY listenApp .

CMD [ "/app/listenApp" ]