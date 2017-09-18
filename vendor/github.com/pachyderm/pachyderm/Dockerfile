FROM ubuntu:14.04
MAINTAINER jdoliner@pachyderm.io

RUN \
  apt-get update -yq && \
  apt-get install -yq --no-install-recommends \
    build-essential \
    ca-certificates \
    cmake \
    curl \
    fuse \
    git \
    libssl-dev \
    mercurial \
    pkg-config && \
  apt-get clean && \
  rm -rf /var/lib/apt
RUN \
  curl -fsSL https://get.docker.com/builds/Linux/x86_64/docker-1.12.1.tgz | tar -C /bin -xz docker/docker --strip-components=1 && \
  chmod +x /bin/docker
RUN \
  curl -sSL https://storage.googleapis.com/golang/go1.8.linux-amd64.tar.gz | tar -C /usr/local -xz && \
  mkdir -p /go/bin
ENV PATH /go/bin:/usr/local/go/bin:$PATH
ENV GOPATH /go
ENV GO15VENDOREXPERIMENT 1
RUN go get github.com/kisielk/errcheck github.com/golang/lint/golint
RUN mkdir -p /go/src/github.com/pachyderm/pachyderm
ADD . /go/src/github.com/pachyderm/pachyderm/
WORKDIR /go/src/github.com/pachyderm/pachyderm
