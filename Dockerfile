FROM golang:1.20.2-alpine

WORKDIR /app

RUN go install github.com/cosmtrek/air@latest
