FROM golang:1.17

ADD .. .

RUN go run server/server.go

ENTRYPOINT ["/go/bin/server"]

EXPOSE 5300