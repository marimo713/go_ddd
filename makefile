.PHONY: build
build:
	go build -o cmd/api/main cmd/api/main.go cmd/api/wire_gen.go
run:
	docker-compose -f build/docker/docker-compose.yaml up
generate:
	go generate -x ./...
wire:
	wire ./cmd/api
test:
	go test ./...
