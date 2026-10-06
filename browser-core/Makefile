CC ?= cc
CFLAGS ?= -std=c11 -Wall -Wextra -Wpedantic -O2
all: zero-browser-core
zero-browser-core: main.o zero_core.o
	$(CC) $(CFLAGS) -o $@ $^
main.o: main.c zero_core.h
	$(CC) $(CFLAGS) -c main.c
zero_core.o: zero_core.c zero_core.h
	$(CC) $(CFLAGS) -c zero_core.c
clean:
	rm -f *.o zero-browser-core
