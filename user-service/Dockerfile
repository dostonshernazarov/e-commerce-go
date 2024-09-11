FROM golang:1.22.6-alpine3.20 AS builder

RUN mkdir app

COPY . /app

WORKDIR /app

RUN go build -o user-service ./cmd/main.go

FROM alpine:3.20

WORKDIR /app

COPY --from=builder /app .

CMD ["/app/user-service"]
