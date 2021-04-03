FROM golang:alpine AS build-env

#MAINTAINER hugh.12345zxcvb@gmail.com

ENV GIN_MODE=release
ENV PORT=8080

# Alpine Linux package management
RUN apk update
RUN apk add git

COPY . /go/src/github.com/lion-devs/ilife-api

EXPOSE $PORT

ENTRYPOINT cd src/github.com/lion-devs/ilife-api/cmd/ilife-api/ && go run main.go

#docker build -t ilifeapi -f build/Dockerfile .
#docker run -itd -p 8080:8080 ilifeapi
