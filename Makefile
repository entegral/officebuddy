build:
	go mod vendor
	docker build --tag obuddies .
	rm -rf vendor

start:
	docker run -p 127.0.0.1:8080:8080 obuddies