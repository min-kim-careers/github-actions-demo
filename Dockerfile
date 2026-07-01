FROM golang:alpine AS base

WORKDIR /go/src/app

ADD . .

RUN go mod init && go build -o go-helloworld

EXPOSE 6111

CMD ["./helloworld"]