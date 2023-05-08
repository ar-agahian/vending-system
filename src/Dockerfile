# syntax=docker/dockerfile:1

FROM golang:1.20 AS build-stage

ENV GO111MODULE=on

WORKDIR /app

COPY go.mod ./

COPY . ./

RUN CGO_ENABLED=0 GOOS=linux go build -a -o /vending-system

FROM alpine AS build-release-stage

WORKDIR /

COPY --from=build-stage /vending-system /vending-system

ENTRYPOINT ["/vending-system"]
