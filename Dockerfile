FROM golang:1.9-alpine3.6
RUN apk add --update --no-cache \
           graphviz \
           ttf-freefont
WORKDIR /go/src/github.com/ardanlabs/gotraining
ENTRYPOINT [ "/bin/sh" ]
