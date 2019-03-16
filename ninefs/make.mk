OBJTYPE=$(shell uname -m | sed 's;.*i[3-6]86.*;386;; s;*x86_64*;amd64;; s;.*armv.*;arm;g;')
OBJTYPE!=uname -m | sed 's;.*i[3-6]86.*;386;; s;*x86_64*;amd64;; s;.*armv.*;arm;g;'
SYSNAME=$(shell uname | tr 'A-Z' 'a-z')
SYSNAME!=uname | tr 'A-Z' 'a-z'

INDENT=clang-format -i
DOCKER=docker
MAKE=make
CC=cc
AR=ar
RANLIB=ranlib
INSTALL=install
CFLAGS+=-O2 -Wall -Wno-parentheses
LDFLAGS+=
ARFLAGS=rsc

PREFIX=/usr/local
VERSION=0.1.0
