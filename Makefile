.PHONY: build start generate test

build:
	go mod vendor
	docker build --tag obuddies .
	rm -rf vendor

start:
	docker run -p 8080:8080 --env-file .env  obuddies

generate:
	go run github.com/99designs/gqlgen generate

test:
	go mod vendor
	docker-compose up --build --abort-on-container-exit
	docker-compose down
	rm -rf vendor
	