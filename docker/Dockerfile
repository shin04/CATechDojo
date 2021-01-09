FROM golang:1.15-alpine

ENV GO111MODULE=on

RUN apk update && apk add git bash
RUN apk add build-base

RUN go get github.com/pilu/fresh

CMD ["fresh"]