OBJTYPE=$(shell uname -m | sed 's;.*i[3-6]86.*;386;; s;.*amd64.*;x86_64;; s;.*armv.*;arm;g;')
OBJTYPE!=uname -m | sed 's;.*i[3-6]86.*;386;; s;.*amd64.*;x86_64;; s;.*armv.*;arm;g;'
SYSNAME=$(shell uname)
SYSNAME!=uname

REPO=github.com/azmodb/bricktop
GENERATOR=go run $(GOPATH)/src/$(REPO)/internal/generator/main.go

INDENT=clang-format -i
#MAKE=make -w
MAKE=make
GO=go
CC=cc
AR=ar
RANLIB=ranlib
INSTALL=install
CFLAGS+=-O2 -Wall
LDFLAGS+=
ARFLAGS=rsc

PREFIX=/usr/local
