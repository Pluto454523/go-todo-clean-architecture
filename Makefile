build:
	go build -o ./tmp/main cmd/generics_server/*.go

run: build
	@export $$(grep -v '^#' .env | xargs); ./tmp/main

gen-http-fiber:
	go run cmd/generate_fiber_interface/main.go

generate-grpc:
	go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
	go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest
	protoc --go_out=. --go_opt=paths=source_relative \
		--go-grpc_out=. --go-grpc_opt=paths=source_relative \
		./src/interface/grpc_server/**/*.proto

test: unit-test coverage-test benchmark-test

unit-test:
	go test -v ./internal/...

coverage-test:
	go test -coverprofile cover.out ./internal/...

benchmark-test:
	go test -bench=. -benchtime=10s -count 3 ./internal/...

coverage-test-html: coverage-test
	go tool cover -html=cover.out

integration-test:
	go test -v ./test/... -p 1