lint:
	go mod vendor
	golangci-lint run -v

test:
	go test -v ./... -count=1 -failfast
