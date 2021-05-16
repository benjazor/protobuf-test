generate-proto:
	protoc -I=./proto/. --go_out=./pkg ./proto/*.proto

run:
	go run .