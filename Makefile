

.PHONY: test
make test-unit:
	go test -v ./... -tags=unit