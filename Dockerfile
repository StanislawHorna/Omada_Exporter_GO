# Build stage
FROM golang:1.24 AS builder

WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
COPY . .

# Build statically linked binary
RUN CGO_ENABLED=0 GOOS=linux go build -o main .

# Minimal runtime stage
FROM alpine:latest
ARG APP_VERSION='0.0.0'

ENV LOKI_APP_VERSION=${APP_VERSION}
ENV LOKI_ENV=prod

RUN apk --no-cache add ca-certificates

WORKDIR /root/

COPY --from=builder /app/main .

CMD ["./main"]