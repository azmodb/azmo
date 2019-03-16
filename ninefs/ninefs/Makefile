include ../make.mk

.PHONY: all build clean install uninstall fmt

CFLAGS+=-I../include
LDFLAGS+=-L../libnpfs -lnpfs -lpthread
CMD=ninefs

HFILES=$(wildcard *.h)
CFILES=$(wildcard *.c)
OFILES=$(CFILES:.c=.o)

all: build

build: $(CMD)

$(CMD): $(OFILES)
	$(CC) -o $(CMD) $(OFILES) $(LDFLAGS)

fmt:
	@$(INDENT) $(CFILES) $(HFILES)

install: build
	$(INSTALL) -m 755 $(CMD) $(PREFIX)/bin

uninstall:
	rm -f $(PREFIX)/bin/$(CMD)

clean:
	rm -f $(OFILES) $(CMD)

$(OFILES): $(HFILES)

.c.o:
	$(CC) $(CFLAGS) -o $@ -c $<

%.o: %.c
	$(CC) $(CFLAGS) -o $@ -c $<
