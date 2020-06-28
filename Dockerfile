FROM alpine:3.12

RUN apk --no-cache add alpine-sdk perl curl

RUN curl -sSLo hey "https://storage.googleapis.com/hey-release/hey_linux_amd64" && \
chmod +x hey && mv hey /usr/local/bin/hey

WORKDIR /
COPY ./bin/microdemo .
RUN chmod 777 ./microdemo

ENTRYPOINT ["./microdemo"]