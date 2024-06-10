all: build

build: clean
	go build -o=./bin/app ./cmd/sso/main.go

protoc: auth

auth:
	mkdir -p ./protos/gen/go/auth && protoc -I=./protos --go_out=./protos/gen/go/auth --go_opt=paths=source_relative --go-grpc_out=./protos/gen/go/auth --go-grpc_opt=paths=source_relative ./protos/auth.proto


clean:
	rm -fR ./bin