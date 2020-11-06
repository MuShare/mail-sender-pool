FROM golang:1.14 as build

ARG VERSION

ADD . /go/src/github.com/MuShare/mail-sender-pool

WORKDIR /go/src/github.com/MuShare/mail-sender-pool

RUN  export GO111MODULE=on GOPROXY=https://proxy.golang.org && \
  go build -ldflags="-X 'main.VERSION=${VERSION}'" -o mail-sender-pool main.go

FROM ubuntu:20.04

RUN apt-get update && apt-get install -y wget

ENV DOCKERIZE_VERSION v0.6.1
RUN wget https://github.com/jwilder/dockerize/releases/download/$DOCKERIZE_VERSION/dockerize-linux-amd64-$DOCKERIZE_VERSION.tar.gz \
    && tar -C /usr/local/bin -xzvf dockerize-linux-amd64-$DOCKERIZE_VERSION.tar.gz \
    && rm dockerize-linux-amd64-$DOCKERIZE_VERSION.tar.gz

COPY --from=build /go/src/github.com/MuShare/mail-sender-pool/mail-sender-pool /usr/bin/

ENTRYPOINT ["/usr/bin/mail-sender-pool"]
