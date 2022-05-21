lint:
	go mod vendor
	golangci-lint run -v

test:
	go mod vendor
	go test -v ./... -count=1 -failfast

integration-test:
	go mod vendor
	go test -v ./... -count=1 -failfast -tags=integration
