test:
	go test -v -race ./...

benchmark:
	go test -run=none -bench=. -benchmem ./...

coverage:
	go test -covermode=atomic -coverprofile=c.out ./...
	go tool cover -html=c.out -o cover.html


.PHONY: test benchmark coverage
