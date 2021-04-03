FROM golang:1.16.3-alpine3.13

#MAINTAINER hugh.12345zxcvb@gmail.com
#MAINTAINER gene31105@gmail.com

ARG PORT

ENV GIN_MODE=release

WORKDIR /src/

COPY . /src

RUN set -eux;\
    CGO_ENABLED=0 go build -v -o /bin/test /src/cmd/ilife-api

EXPOSE $PORT

CMD /bin/test
