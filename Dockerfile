FROM golang:1.8

MAINTAINER pan@cloudtogo.cn

COPY ./  /go/src/

RUN cd /go/src && \
     mv sources.list /etc/apt/sources.list && \
     rm -f Dockerfile && \
     apt-get update && apt-get upgrade && apt-get install apache2-utils

CMD cd /go/src && go run ab.go
