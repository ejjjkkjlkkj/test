#ifndef ZERO_READER_H
#define ZERO_READER_H
#include "zero_core.h"
typedef struct { unsigned long spoken_sequence; unsigned interrupt_count; } ZeroReader;
void zero_reader_init(ZeroReader *r);
void zero_reader_consume(ZeroReader *r, const ZeroEvent *event);
void zero_reader_interrupt(ZeroReader *r);
#endif
