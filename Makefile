.PHONY: clean test

aoc:
	go build -o $@ main.go

clean:
	rm ./aoc

test:
	go test -v ./...
