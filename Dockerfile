# syntax=docker/dockerfile:1

FROM golang:1.23.1-alpine

WORKDIR /app

COPY . ./

RUN go mod download

RUN CGO_ENABLED=0 GOOS=linux go build -o /sse-service

EXPOSE 80

CMD ["/sse-service"]