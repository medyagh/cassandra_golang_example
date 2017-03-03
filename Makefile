BINARY=cassandra_golang_example

VERSION=1.0.0
BUILD=`date +%FT%T%z`

LDFLAGS=-ldflags "-X main.Version=${VERSION} -X main.Build=${BUILD}"

all: deps build test

deps:
	glide install


build:
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build ${LDFLAGS} -o ${BINARY}

mac:
	CGO_ENABLED=0 GOOS=darwin GOARCH=amd64 go build ${LDFLAGS} -o ${BINARY}


install:
	go install ${LDFLAGS}

clean:
	if [ -f ${BINARY} ] ; then rm ${BINARY} ; fi
	go clean
	docker-compose down

sclean: clean
	rm glide.lock

vet:
	go vet

up: clean deps build

	docker-compose build
	docker-compose up

.PHONY: build
