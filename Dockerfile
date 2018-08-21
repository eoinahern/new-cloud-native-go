FROM golang:1.7.4
MAINTAINER eoin.ahern

ENV SOURCES usr/local/go/src/github.com/eoinahern/new-cloud-native-go

COPY  . ${SOURCES}

RUN cd  ${SOURCES} && CGO_ENABLE=0 go install

ENV PORT 8080
EXPOSE 8080

ENTRYPOINT new-cloud-native-go
