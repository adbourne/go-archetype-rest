#
default: test

clean:
	rm -rf bin/*; rm -rf pkg/*

build:
	go install

test:
	go test