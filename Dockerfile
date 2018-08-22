FROM golang:1.7.4
MAINTAINER eoin.ahern

#ENV SOURCES github.com/eoinahern/new-cloud-native-go

COPY  ./new-cloud-native-go  /app/new-cloud-native-go


RUN  cd /app/new-cloud-native-go && go install

ENV PORT 8080
EXPOSE 8080

ENTRYPOINT /app/new-cloud-native-go
