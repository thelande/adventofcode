BIN=aoc2023

ifeq ($(OS),Windows_NT)
	BIN := $(BIN).exe
endif

all: $(BIN)

$(BIN): *.go days/**/*.go ../common/*.go
	go build -o $(BIN) .

.PHONY: clean
clean:
	rm $(BIN)
