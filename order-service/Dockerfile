FROM golang:1.22 AS builder

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN go build -o main ./cmd

FROM gcr.io/distroless/base

COPY --from=builder /app/main /app/

WORKDIR /app

ENTRYPOINT ["/app/main"]