FROM golang:1.7.4
MAINTAINER eoin.ahern

COPY  . ./github.com/eoinahern/new-cloud-native-go

RUN cd ./github.com/eoinahern/new-cloud-native-go && CGO_ENABLE=0 go install

ENV PORT 8080
EXPOSE 8080

ENTRYPOINT new-cloud-native-go
