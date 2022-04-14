FROM golang:1.18-alpine AS builder
ENV PORT

ADD . /pingpong
WORKDIR /pingpong

RUN go build -o /app

FROM debian:buster
EXPOSE PORT

WORKDIR /
COPY --from=builder /app /

ENTRYPOINT ["/app"]
