FROM golang:latest

RUN mkdir -p /go/src/github.com/ardanlabs
#
# downloading the latest gotraining material so that it allows to
# run the container without mapping to any local gotraining copy
# e.g.
#       docker build -t ardanlabs-gotraining .
#       docker run -it ardanlabs-gotraining
#
RUN git clone https://github.com/ardanlabs/gotraining /go/src/github.com/ardanlabs/gotraining

RUN mkdir -p /go/src/github.com/ardanlabs/gotraining

WORKDIR /go/src/github.com/ardanlabs/gotraining

CMD /bin/bash
