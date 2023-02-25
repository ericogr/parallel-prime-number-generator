BIN_NAME=prime-gen
PRIMES=5000
JOBS=2
DEBUG=false

all: clean build

build:
	@go build -o bin/$(BIN_NAME) .

clean:
	@rm -rf bin

run: clean build
	@bin/$(BIN_NAME) --primes=$(PRIMES) --jobs=$(JOBS) --debug=$(DEBUG)