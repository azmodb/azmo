include ../make.mk

.PHONY: all build clean install uninstall fmt generate

all: build

build:
	@$(GO) build

generate:
	@$(GENERATOR) -go -root .

fmt:
	@$(GO) fmt

install: build
	@echo -n

uninstall:
	@echo -n

clean:
	@echo -n
