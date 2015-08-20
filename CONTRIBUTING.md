# Contributing

Pull request are always welcome and very much appreciated! To make things go as smoothly as possible please consider following the steps below before sending a pull request:

1. Squash unnecassary commits with `git rebase -i`.

2. Format the code with `go fmt` to ensure proper spacing and such.

3. Run the unit tests with `go test` to make sure there are no breaking changes.

4. It would be great if you wrote tests for your work. You can generate coverage reports easily with the following commands:

```
# Generate coverage data
go test -coverprofile=coverage.out

# View coverage by function
go tool cover -func="coverage.out"

# View coverage line by line
go tool cover -html="coverage.out"
```

![Correlation](http://imgs.xkcd.com/comics/correlation.png)