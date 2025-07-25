.PHONY: build clean

build:
	go build -o bin/sclaude ./main.go

clean:
	rm -rf bin/
