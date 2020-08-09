.PHONY: run
run: clean build
	go run ./build.go -run deploy

.PHONY: build
build:
	go run ./build.go -fast build

.PHONY: release
release:
	go run ./build.go build

.PHONY: clean
clean:
	go run ./build.go clean
