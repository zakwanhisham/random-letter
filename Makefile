all: build

build: main.go
	@echo "Building....."
	@go build -o bin/randl main.go

install: build
	@echo "Installing to `~/.local/bin`"
	@cp -iv bin/randl ~/.local/bin

clean:
	@echo "Cleaning....."
	@rm -rvf bin/*

.PHONY: all build clean
