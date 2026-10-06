#ifndef ZERO_CORE_H
#define ZERO_CORE_H
#include <stddef.h>
typedef enum { ZERO_DOCUMENT, ZERO_HEADING, ZERO_PARAGRAPH, ZERO_LINK, ZERO_BUTTON, ZERO_TEXTBOX, ZERO_CHECKBOX, ZERO_RADIO, ZERO_LIST, ZERO_LIST_ITEM, ZERO_TABLE, ZERO_ROW, ZERO_CELL, ZERO_DIALOG, ZERO_ALERT } ZeroRole;
typedef struct ZeroNode { unsigned long id; ZeroRole role; const char *name; const char *value; unsigned state; struct ZeroNode *parent; struct ZeroNode **children; size_t child_count; } ZeroNode;
typedef enum { ZERO_EVENT_DOCUMENT_START, ZERO_EVENT_DOCUMENT_END, ZERO_EVENT_FOCUS_ENTER, ZERO_EVENT_FOCUS_LEAVE, ZERO_EVENT_VALUE_CHANGE, ZERO_EVENT_STATE_CHANGE, ZERO_EVENT_ALERT, ZERO_EVENT_ERROR, ZERO_EVENT_UNKNOWN } ZeroEventKind;
typedef struct { ZeroEventKind kind; unsigned long target; unsigned long sequence; } ZeroEvent;
typedef struct { ZeroNode *root; ZeroNode *focus; unsigned long next_sequence; } ZeroCore;
void zero_core_init(ZeroCore *core, ZeroNode *root);
int zero_focus_first(ZeroCore *core);
int zero_focus_next(ZeroCore *core);
int zero_focus_previous(ZeroCore *core);
const ZeroEvent *zero_last_event(void);
#endif
