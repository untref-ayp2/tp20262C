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
	@golangci-lint version 2>/dev/null | grep -q 'version 2' || { echo 'Error: se requiere golangci-lint v2.x. Instalalo: https://golangci-lint.run/welcome/install/'; exit 1; }
	golangci-lint run

clean:
	rm -f output/*.txt