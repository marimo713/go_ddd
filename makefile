.PHONY: build
build:
	go build -o cmd/api/main cmd/api/main.go
run:
	docker-compose -f build/docker/docker-compose.yaml up
