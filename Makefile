#!/usr/bin/make -f

GO=go
GB=gb

darwin-amd64:
	env GOOS=darwin GOARCH=amd64 $(GB) build

linux-amd64:
	env GOOS=linux GOARCH=amd64 $(GB) build

all: test clean darwin-amd64 linux-amd64

clean:
	rm -fR pkg bin

test:
	$(GB) test -v
