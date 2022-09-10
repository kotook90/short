.PHONY: build
build:
	go build -o bin/main cmd/shortner/main.go


.PHONY: run
run:
	go run cmd/shortner/main.go

.PHONY: test
test:
	go test ./...

PHONY: lint
lint:
	golangci-lint run ./...

.PHONY: pc
pc:
	pre-commit run --all-files

compile:
	echo "Compiling for every OS and Platform"
    # 32-Bit Systems
    # FreeBDS
    GOOS=freebsd GOARCH=386 go build -o bin/main-freebsd-386 main.go
    # MacOS
    GOOS=darwin GOARCH=386 go build -o bin/main-darwin-386 main.go
    # Linux
    GOOS=linux GOARCH=386 go build -o bin/main-linux-386 main.go
    # Windows
    GOOS=windows GOARCH=386 go build -o bin/main-windows-386 main.go
        # 64-Bit
    # FreeBDS
    GOOS=freebsd GOARCH=amd64 go build -o bin/main-freebsd-amd64 main.go
    # MacOS
    GOOS=darwin GOARCH=amd64 go build -o bin/main-darwin-amd64 main.go
    # Linux
    GOOS=linux GOARCH=amd64 go build -o bin/main-linux-amd64 main.go
    # Windows
    GOOS=windows GOARCH=amd64 go build -o bin/main-windows-amd64 main.go

PHONY: clean
clean:
	rm -rf  ./bin



all:build run test lint pc clean
