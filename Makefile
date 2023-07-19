PHONY: test coverage benchmark

test:
	go test ./...

coverage:
	go test -cover ./...

benchmark:
	go test -bench=.

errcheck:
	#go install github.com/kisielk/errcheck@latest
	errcheck .