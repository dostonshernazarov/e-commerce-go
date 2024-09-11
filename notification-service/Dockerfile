FROM golang:1.23.0-alpine AS build

WORKDIR /src

COPY . .

WORKDIR /src/cmd/notification

RUN go mod download

RUN go build -o /bin/ntfk

FROM alpine:latest

RUN apk --no-cache add ca-certificates

COPY --from=build /bin/ntfk /bin/ntfk

CMD [ "/bin/ntfk" ]
