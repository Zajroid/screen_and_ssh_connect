BINARY_NAME=main.out

all: build test

build:
	go build -o ${BINARY_NAME} make_screenshot.go connect_to_server.go main.go

test:
	go test -v main.go

run:
	go build -o ${BINARY_NAME} main.go
	./${BINARY_NAME}

clean:
	go clean
	rm ${BINARY_NAME}

deps:
	go get github.com/kbinani/screenshot
	go get golang.org/x/crypto/ssh
	go get gopkg.in/ini.v1