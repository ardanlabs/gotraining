FROM		alpine
RUN             apk update && apk add make gcc linux-headers git perl musl-dev go
RUN		git clone https://github.com/xianyi/OpenBLAS && cd OpenBLAS && make && make PREFIX=/usr install
RUN		mkdir -p /go/src /go/bin /go/pkg
ENV		GOPATH=/go
RUN		go get github.com/gonum/blas github.com/sjwhitworth/golearn
