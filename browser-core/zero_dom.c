#include "zero_dom.h"
#include <stdlib.h>
#include <string.h>
static ZeroRole role_for(const char *s){if(!strcmp(s,"h1")||!strcmp(s,"h2")||!strcmp(s,"h3"))return ZERO_HEADING;if(!strcmp(s,"a"))return ZERO_LINK;if(!strcmp(s,"button"))return ZERO_BUTTON;if(!strcmp(s,"input"))return ZERO_TEXTBOX;if(!strcmp(s,"p"))return ZERO_PARAGRAPH;if(!strcmp(s,"li"))return ZERO_LIST_ITEM;if(!strcmp(s,"ul")||!strcmp(s,"ol"))return ZERO_LIST;return ZERO_DOCUMENT;}
void zero_document_init(ZeroDocument*d,ZeroNode*n,ZeroNode**c,size_t cap){d->nodes=n;d->child_storage=c;d->node_count=0;d->capacity=cap;}
static int add(ZeroDocument*d,ZeroRole r,const char*name,const char*value){if(d->node_count>=d->capacity)return 0;ZeroNode*n=&d->nodes[d->node_count];memset(n,0,sizeof(*n));n->id=(unsigned long)d->node_count+1;n->role=r;n->name=name;n->value=value;d->node_count++;return 1;}
int zero_document_parse(ZeroDocument*d,const char*in){if(!d||!in)return 0;const char*p=in;while(*p){if(*p!='<'){p++;continue;}const char*s=++p;while(*p&&*p!='>')p++;if(!*p)break;size_t len=(size_t)(p-s);char tag[32];if(len>=sizeof(tag))len=sizeof(tag)-1;memcpy(tag,s,len);tag[len]=0;if(tag[0]=='/') {p++;continue;}char*sp=strchr(tag,' ');if(sp)*sp=0;ZeroRole r=role_for(tag);if(r!=ZERO_DOCUMENT&&!add(d,r,tag,NULL))return 0;p++;}return 1;}
void zero_document_free(ZeroDocument*d){if(d){d->nodes=NULL;d->child_storage=NULL;d->node_count=0;d->capacity=0;}}
