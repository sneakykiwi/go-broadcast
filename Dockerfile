FROM golang:1.16.5-buster

RUN mkdir /app

ADD . /app

WORKDIR /app


RUN go build .