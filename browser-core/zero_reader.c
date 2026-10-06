#include "zero_reader.h"
#include <stdio.h>
void zero_reader_init(ZeroReader*r){r->spoken_sequence=0;r->interrupt_count=0;}
void zero_reader_consume(ZeroReader*r,const ZeroEvent*e){if(!r||!e)return;r->spoken_sequence=e->sequence;printf("ACCESS event=%d target=%lu sequence=%lu\n",(int)e->kind,e->target,e->sequence);}
void zero_reader_interrupt(ZeroReader*r){if(r)r->interrupt_count++;}
