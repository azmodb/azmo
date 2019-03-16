include ../make.mk

.PHONY: all build clean install uninstall fmt generate

CFLAGS+=-Wno-unused-function
LIB=libninep.a

HFILES=$(wildcard *.h)
CFILES=$(wildcard *.c)
OFILES=$(CFILES:.c=.o)

all: build

build: $(LIB)

$(LIB): $(OFILES)
	$(AR) $(ARFLAGS) $(LIB) $(OFILES)
	$(RANLIB) $(LIB)

generate:
	@$(GENERATOR) -h -root .
	@$(GENERATOR) -c -root .

fmt:
	@$(INDENT) $(CFILES) $(HFILES)

install: build
	@echo -n

uninstall:
	@echo -n

clean:
	rm -f $(OFILES) $(LIB)

$(OFILES): $(HFILES)

%.o: %.c 
	$(CC) $(CFLAGS) -c $<
