FROM golang:latest
LABEL maintainer="Vyacheslav Kryuchenko <v.kryuchenko@corp.mail.ru>"
ENV GOOS="linux"
ADD . /sources
WORKDIR /sources

ENTRYPOINT ["go", "build", "lfserv.go"]
