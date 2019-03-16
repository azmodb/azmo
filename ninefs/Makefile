include make.mk

.PHONY: all build install uninstall clean fmt image

all: build

build:
	@$(MAKE) -C libnpfs build
	@$(MAKE) -C ninefs build

fmt:
	@$(MAKE) -C libnpfs fmt
	@$(MAKE) -C ninefs fmt

install:
	@$(MAKE) -C libnpfs install
	@$(MAKE) -C ninefs install

uninstall:
	@$(MAKE) -C libnpfs uninstall
	@$(MAKE) -C ninefs uninstall

clean:
	@$(MAKE) -C libnpfs clean
	@$(MAKE) -C ninefs clean

image:
	@$(DOCKER) build --tag=ninefs:latest \
		--tag=ninefs:$(VERSION) .
