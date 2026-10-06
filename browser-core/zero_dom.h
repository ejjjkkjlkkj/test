#ifndef ZERO_DOM_H
#define ZERO_DOM_H
#include <stddef.h>
#include "zero_core.h"
typedef struct { ZeroNode *nodes; ZeroNode **child_storage; size_t node_count; size_t capacity; } ZeroDocument;
void zero_document_init(ZeroDocument *d, ZeroNode *nodes, ZeroNode **children, size_t capacity);
int zero_document_parse(ZeroDocument *d, const char *input);
void zero_document_free(ZeroDocument *d);
#endif
