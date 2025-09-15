FROM golang:1.25 AS builder
WORKDIR /app

COPY . .
RUN go mod download
RUN go mod tidy

WORKDIR /app
RUN mkdir /app/build

RUN cd /app/cmd/api && CGO_ENABLED=0 GOOS=linux go build -o ../../build/api

FROM alpine:latest AS meli-api
RUN apk --no-cache add ca-certificates

WORKDIR /root/
COPY --from=builder /app/build/api .
