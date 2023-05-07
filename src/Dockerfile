# syntax=docker/dockerfile:1

FROM golang:1.20

ENV GO111MODULE=on

WORKDIR /app

COPY go.mod ./

COPY *.go ./

RUN CGO_ENABLED=0 GOOS=linux go build -a -o /vending-system

CMD [ "/vending-system" ]
