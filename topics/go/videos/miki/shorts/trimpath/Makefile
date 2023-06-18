all: build
	rm -f max

build:
	go build ./cmd/max
	- ./max
	go build -trimpath ./cmd/max
	- ./max
