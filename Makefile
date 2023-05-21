build:
	go build -o checkout ./cmd/checkout/*.go

test:
	go test ./...

test-with-coverage:
	go test ./... -cover

test-generate-coverage-report:
	go test ./... -coverprofile=coverage.out
	go tool cover -html=coverage.out