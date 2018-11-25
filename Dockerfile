FROM golang:1.11-alpine as base
RUN apk add make
WORKDIR ${GOPATH}/src/github.com/bombergame/chats-service
COPY . .
RUN go build . && mv ./chats-service /tmp/service

FROM alpine:latest
WORKDIR /tmp
COPY --from=base /tmp/service .
ENTRYPOINT ./service --http_port=80
EXPOSE 80
