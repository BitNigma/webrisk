.PHONE: build 

.PHONY: gen

gen :
	templ generate

build : gen
		go build -o ./webrisk cmd/main.go

.PHONE : test
test :
	go test -v -race -timeout 30s ./...

.DEFAULT_GOAL := build