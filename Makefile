BIN_NAME=prime-gen
NUMBERS=5000
JOBS=2
DEBUG=false

all: clean build

build:
	go build -o bin/$(BIN_NAME) .

clean:
	rm bin/*
	rmdir bin

run:
	bin/$(BIN_NAME) --numbers=$(NUMBERS) --jobs=$(JOBS) --debug=$(DEBUG)