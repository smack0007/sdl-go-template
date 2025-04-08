REPO_PATH := $(dir $(realpath $(lastword $(MAKEFILE_LIST))))

build:
	go build -o ./bin/app ./app

clean:
	go clean
	rm -rf ./bin

run:
	go run ./app