export GO111MODULE=on

test-verbose:
	go test ./... -coverprofile=./coverage.txt -covermode=atomic -v

test-with-report:
	mkdir -p coverprofile
	go test ./... -coverprofile coverprofile/coverage.out
	go tool cover -html=coverprofile/coverage.out -o coverprofile/coverage.html
