.PHONY: run build test test-v fmt lint clean

run:
	go run ./cmd/simulador data/basico/entrada.json

build:
	go build ./...

test:
	go test ./...

test-v:
	go test -v ./...

fmt:
	go fmt ./...

lint:
	golangci-lint run

clean:
	rm -f output/*.txt