.PHONY: test
test:
	go test -v -race ./...

.PHONY: benchmark
benchmark:
	go test -benchmem -bench=. ./...

.PHONY: coverage
coverage:
	go test -covermode=atomic -coverprofile=c.out ./...
	go tool cover -html=c.out -o cover.html
