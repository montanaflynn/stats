.PHONY: all

default: test lint

format: 
	go fmt .

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
	git tag ${TAG}
	git-chglog $(TAG) | tail -n +4 | sed '1s/^/$(TAG)\n/gm' > release-notes.txt
	git push origin master ${TAG}
	hub release create --copy -F release-notes.txt ${TAG}

