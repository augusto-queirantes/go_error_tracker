FROM golang:1.20.2-alpine

WORKDIR /app

# Hot reload
RUN go install github.com/cosmtrek/air@latest

# Linter
RUN go install github.com/golangci/golangci-lint/cmd/golangci-lint@v1.52.2
