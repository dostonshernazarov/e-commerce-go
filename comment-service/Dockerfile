FROM golang:1.22.3

WORKDIR /app

COPY . .

RUN go build -o main cmd/main.go

CMD ["./main"]
