FROM golang:1.19.2-alpine3.13 AS builder

ENV GOPATH="$HOME/go"

WORKDIR $GOPATH/src/github.com/vitorsavian/github-api

COPY . $GOPATH/src/github.com/vitorsavian/github-api

RUN apk update && apk upgrade && apk add curl git && apk add gcc libc-dev

RUN go build -ldflags '-linkmode=external'

RUN go get -u github.com/swaggo/swag/cmg/swag

RUN swag init

#########################################################################

FROM alpine:3.13

ENV GOPATH="$HOME/go"

WORKDIR /app

RUN apk update

COPY --from=builder $GOPATH/src/github.com/vitorsavian/github-api .