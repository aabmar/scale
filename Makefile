
init:
	go get github.com/karalabe/hid

run:
	go run src/hid.go

build:
	mkdir -p bin
	go build -o bin/hid src/hid.go

