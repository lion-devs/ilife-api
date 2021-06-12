FROM golang:1.16.4-alpine3.13

#MAINTAINER hugh.12345zxcvb@gmail.com
#MAINTAINER gene31105@gmail.com

ARG PORT

ENV GIN_MODE=release

WORKDIR /src/

COPY . /src

# Disabling cgo results in a fully static binary that can run without C libs
RUN set -eux;\
    CGO_ENABLED=0 go build -a -v -o /bin/ilife-api /src/cmd

EXPOSE $PORT

CMD /bin/ilife-api
