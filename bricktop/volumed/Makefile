include ../make.mk

.PHONY: all build clean install uninstall fmt

CMD=volumed
HFILES=

CFLAGS+=-I../libninep
LDFLAGS+=-L../libninep -lninep -lpthread

CFILES=$(wildcard *.c)
OFILES=$(CFILES:.c=.o)

all: build

build: $(CMD)

$(CMD): $(OFILES)
	$(CC) -o $(CMD) $(OFILES) $(LDFLAGS)

fmt:
	@$(INDENT) $(CFILES) $(HFILES)

install: build
	@echo -n

uninstall:
	@echo -n

clean:
	rm -f $(OFILES) $(CMD)

$(OFILES): $(HFILES)

%.o: %.c 
	$(CC) $(CFLAGS) -c $<
