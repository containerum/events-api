FROM golang:1.10-alpine as builder
RUN apk add --update make git
WORKDIR /go/src/github.com/containerum/events-api
COPY . .
RUN VERSION=$(git describe --abbrev=0 --tags) make build-for-docker

FROM alpine:3.7

COPY --from=builder /tmp/events-api /
ENV CH_RESOURCE_DEBUG="true" \
    CH_RESOURCE_TEXTLOG="true" \
    CH_RESOURCE_MONGO_ADDR="http://mongo:27017"

EXPOSE 1213

CMD "/resource-service"
