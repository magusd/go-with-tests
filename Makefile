PHONY: test coverage benchmark

test:
	go test ./...

coverage:
	go test -cover ./...

benchmark:
	go test -bench=.