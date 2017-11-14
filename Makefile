BUILD_DIR=$(shell pwd)/build
SRC_DIR=$(shell pwd)/src

SUFFIX=""

GOARCH=amd64

all: export GOPATH=$(shell pwd)
all: windows linux mac

fcgi: SUFFIX=.fcgi
fcgi: windows linux mac

windows:
	cd ${SRC_DIR}; \
	echo "Building webapi-dav-windows_amd64"; \
	env GOOS=windows GOARCH=amd64 go build -o ../build/webapi-dav-windows_${GOARCH}${SUFFIX}; \
	cd - >/dev/null

linux:
	cd ${SRC_DIR}; \
	echo "Building webapi-dav-linux_amd64"; \
	env GOOS=linux GOARCH=amd64 go build -o ../build/webapi-dav-linux_${GOARCH}${SUFFIX}
	cd - >/dev/null

mac:
	cd ${SRC_DIR}; \
    	echo "Building webapi-dav-mac_amd64"; \
    	env GOOS=darwin GOARCH=amd64 go build -o ../build/webapi-dav-mac_${GOARCH}${SUFFIX}; \
    	cd - >/dev/null

fmt:
	cd ${SRC_DIR}; \
	go fmt $$(go list -f '{{ .GoFiles }}' | grep *.go)

clean:
	cd ${BUILD_DIR}; \
	rm -rf webapi-dav-*

.PHONY: windows linux mac fmt clean
