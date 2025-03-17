FROM golang:1.24-alpine AS builder

RUN apk add --no-cache git ca-certificates && update-ca-certificates

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN GOOS=linux GOARCH=amd64 go build -o users-service ./cmd/api

RUN ls -la /app/users-service

FROM alpine:latest

WORKDIR /root/

COPY --from=builder /app/users-service /users-service

RUN chmod +x /users-service

COPY --from=builder /app/kit/config /app/kit/config

ENV CONF_DIR=/app/kit/config
ENV SCOPE=stage

RUN chmod -R 755 /app/kit/config

CMD ["/users-service"]