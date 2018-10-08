FROM golang:1.10-alpine as builder
RUN apk add --update make git
WORKDIR /go/src/github.com/containerum/events-api
COPY . .
RUN VERSION=$(git describe --abbrev=0 --tags) make build-for-docker

FROM alpine:3.7

COPY --from=builder /tmp/events-api /
ENV MONGO_ADDR="mongo-mongodb.mongo.svc.cluster.local:27017" \
    MONGO_DB="events"

EXPOSE 1667

CMD "/events-api"
