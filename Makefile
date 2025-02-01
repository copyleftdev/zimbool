# Makefile for Zimbool CLI Tool

# Use GOBIN if set; otherwise default to $(HOME)/.local/bin.
PREFIX ?= $(if $(GOBIN),$(GOBIN),$(HOME)/.local)
BINDIR := $(PREFIX)/bin
BINARY := zimbool

.PHONY: all build install clean fmt vet test run

all: build

build:
	go build -o $(BINARY) .

install: build
	@mkdir -p $(BINDIR)
	@cp $(BINARY) $(BINDIR)/$(BINARY)
	@echo "Installed $(BINARY) to $(BINDIR)"

clean:
	@rm -f $(BINARY)

fmt:
	go fmt ./...

vet:
	go vet ./...

test:
	go test ./...

run: build
	./$(BINARY)
