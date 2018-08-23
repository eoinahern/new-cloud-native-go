FROM golang:1.7.4
MAINTAINER eoin.ahern

COPY  . /go/src/github.com/eoinahern/new-cloud-native-go

WORKDIR /go/src/github.com/eoinahern/new-cloud-native-go
RUN  CGO_ENABLED=0 && go install

ENV PORT 8080
EXPOSE 8080

ENTRYPOINT new-cloud-native-go
