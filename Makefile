all: test build run

modules:
	go mod tidy

go-build: modules
	go build  -mod=readonly  -o bin/short/short.exe ./cmd/short/.

build:
	docker-compose build

build-docker:
	docker build -t url-short .

run:
	docker-compose up

run-docker:
	docker run -p 8000:8000 url-short

stop:
	docker stop url-short

go-test:
	go test ./...

test:
	newman run test/go.postman_collection.json