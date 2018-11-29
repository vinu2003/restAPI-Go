FROM golang:1.7.4-alpine
MAINTAINER Vinodhini

ENV HOME /Users/vinodhinibalusamy
ENV SOURCES /Users/vinodhinibalusamy/go/src/package/CloudNativeGo
ENV GOPATH $HOME/go

COPY . ${SOURCES}

RUN cd ${SOURCES} && CGO_ENABLED=0 go install -a

ENV PORT 8995
EXPOSE 8995

ENTRYPOINT Cloud-Native-Go