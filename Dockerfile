FROM golang:1.11-alpine3.8
RUN apk add --update --no-cache \
           graphviz \
           ttf-freefont
WORKDIR /go/src/github.com/ardanlabs/gotraining
ENTRYPOINT [ "/bin/sh" ]
