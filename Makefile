.PHONY: all

format: 
	go fmt

test:
	go test -race 
	
check: format test

benchmark:
	go test -bench=. -benchmem

coverage:
	go test -coverprofile=coverage.out
	go tool cover -html="coverage.out"

lint: format
	golangci-lint run .

changelog:
	git-chglog -o CHANGELOG.md

docs:
	godoc2md github.com/montanaflynn/stats | sed -e s#src/target/##g > DOCUMENTATION.md

release:
	RELEASE_NOTES:=$(shell git-chglog $(TAG))
	$(RELEASE_NOTES)
	git tag ${TAG}
	git push origin master ${TAG}
	hub release create -m $(RELEASE_NOTES) ${TAG}

default: lint test
