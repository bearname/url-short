.PHONY: all
all: test build run

.PHONY: modules
modules:
	go mod tidy

.PHONY: download
download:
	go mod download

.PHONY: go-build
go-build: modules
	go build  -mod=readonly  -o bin/short/short.exe ./cmd/short/.

.PHONY: build
build:
	docker-compose build

.PHONY: build-docker
build-docker:
	docker build -t url-short .

.PHONY: run
run:
	docker-compose up

.PHONY: run-docker
run-docker:
	docker run -p 8000:8000 url-short

.PHONY: stop
stop-docker:
	docker stop url-short

.PHONY: go-test
go-test:
	go test ./...

.PHONY: test
test:
	newman run test/go.postman_collection.json