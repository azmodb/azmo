include make.mk

.PHONY: all build install uninstall clean fmt

all: build

build:
	@$(MAKE) -C libninep build
	@$(MAKE) -C volumed build

fmt:
	@$(MAKE) -C libninep fmt
	@$(MAKE) -C volumed fmt

install:
	@$(MAKE) -C libninep install
	@$(MAKE) -C volumed install

uninstall:
	@$(MAKE) -C libninep uninstall
	@$(MAKE) -C volumed uintstall

clean:
	@$(MAKE) -C libninep clean
	@$(MAKE) -C volumed clean
