.PHONY:
.SILENT:

build:
	go build -o ./.bin/server cmd/main.go

run: build
	./.bin/server
	