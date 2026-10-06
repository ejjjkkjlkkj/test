CC ?= cc
CFLAGS ?= -std=c11 -Wall -Wextra -Wpedantic -O2

all: zero-browser-core zero-media-test

zero-browser-core: zero_core.o zero_dom.o zero_reader.o main.o
	$(CC) $(CFLAGS) -o $@ $^

zero-media-test: zero_media.o zero_media_test.o
	$(CC) $(CFLAGS) -o $@ $^

zero_core.o: zero_core.c zero_core.h
	$(CC) $(CFLAGS) -c zero_core.c

zero_dom.o: zero_dom.c zero_dom.h zero_core.h
	$(CC) $(CFLAGS) -c zero_dom.c

zero_reader.o: zero_reader.c zero_reader.h zero_core.h
	$(CC) $(CFLAGS) -c zero_reader.c

zero_media.o: zero_media.c zero_media.h
	$(CC) $(CFLAGS) -c zero_media.c

zero_media_test.o: zero_media_test.c zero_media.h
	$(CC) $(CFLAGS) -c zero_media_test.c

main.o: main.c zero_core.h
	$(CC) $(CFLAGS) -c main.c

clean:
	rm -f *.o zero-browser-core zero-media-test
