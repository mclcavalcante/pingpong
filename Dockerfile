FROM golang:1.18-alpine AS builder

WORKDIR /go/src/github.com/lpegoraro/pingpong-go-module
COPY . .

RUN go build
