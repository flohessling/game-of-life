BINARY_NAME=gol

build-binary:
	CGO_ENABLED=0 go build -o dist/$(BINARY_NAME) -ldflags "-s -w" cmd/main.go

dep:
	go mod download
