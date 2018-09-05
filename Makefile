default: lint test

deps:
	go get github.com/alecthomas/gometalinter
	gometalinter --install

doc:
	godoc `pwd`

webdoc:
	godoc -http=:44444

format:
	go fmt

lint: format
	gometalinter --disable gocyclo deadcode

test:
	go test -race

check: format test

benchmark:
	go test -bench=. -benchmem

coverage:
	go test -coverprofile=coverage.out
	go tool cover -html="coverage.out"
