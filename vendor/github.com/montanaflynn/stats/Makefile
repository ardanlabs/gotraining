default: deps lint test

deps:
	go get github.com/alecthomas/gometalinter
	gometalinter --install

lint:
	go fmt
	gometalinter 

test:
	go test -race

check: lint test

benchmark:
	go test -bench=. -benchmem

coverage:
	go test -cover

report:
	go test -coverprofile=coverage.out
	go tool cover -html="coverage.out"

.PHONY: default deps lint test check benchmark coverage report