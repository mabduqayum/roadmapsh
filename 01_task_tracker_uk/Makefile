.PHONY: build run test clean

build:
	go build -o task-cli

run: build
	./task-cli

test:
	go test ./...

clean:
	rm -f task-cli
