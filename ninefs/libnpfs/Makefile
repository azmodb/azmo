include ../make.mk

.PHONY: all build clean install uninstall fmt

CFLAGS+=-I../include
LIB=libnpfs.a

HFILES=$(wildcard *.h)
CFILES=$(wildcard *.c)
OFILES=$(CFILES:.c=.o)

all: build

build: $(LIB)

$(LIB): $(OFILES)
	$(AR) $(ARFLAGS) $(LIB) $(OFILES)
	$(RANLIB) $(LIB)

fmt:
	@$(INDENT) $(CFILES) $(HFILES)

install: build
	@echo -n

uninstall:
	@echo -n

clean:
	rm -f $(OFILES) $(LIB)

$(OFILES): $(HFILES)

.c.o:
	$(CC) $(CFLAGS) -o $@ -c $<

%.o: %.c
	$(CC) $(CFLAGS) -o $@ -c $<
