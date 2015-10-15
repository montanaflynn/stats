.PHONY: all

default: format test

docs:
	godoc `pwd`

format: 
	go fmt

test: lint
	go test -race 
	
benchmark:
	go test -bench=. -benchmem

coverage:
	go test -coverprofile=coverage.out
	go tool cover -html="coverage.out"

lint: format
	go get github.com/alecthomas/gometalinter
	gometalinter --install
	gometalinter 

check: lint test
