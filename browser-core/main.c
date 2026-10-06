#include <stdio.h>
#include "zero_core.h"
static ZeroNode heading={2,ZERO_HEADING,"ZERO",NULL,0,NULL,NULL,0};
static ZeroNode link={3,ZERO_LINK,"Example","https://example.com",0,NULL,NULL,0};
static ZeroNode *children[]={&heading,&link};
static ZeroNode root={1,ZERO_DOCUMENT,"document",NULL,0,NULL,children,2};
static void show(const char*a){const ZeroEvent*e=zero_last_event();printf("%s event=%d target=%lu sequence=%lu\n",a,(int)e->kind,e->target,e->sequence);}
int main(void){ZeroCore c;zero_core_init(&c,&root);zero_focus_first(&c);show("FIRST");while(zero_focus_next(&c))show("NEXT");while(zero_focus_previous(&c))show("PREVIOUS");return 0;}
