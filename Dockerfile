FROM golang:1.18 AS builder

ARG PORT
ARG DROPS_PRESCRIPTION

ADD . /pingpong
WORKDIR /pingpong

RUN go build -o /app ./cmd

FROM debian:buster
EXPOSE $PORT

WORKDIR /
COPY --from=builder /app /

ENTRYPOINT ["/app"]
