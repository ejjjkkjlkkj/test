#include "zero_core.h"
static ZeroEvent last_event;
static void emit(ZeroEventKind k,unsigned long t,unsigned long s){last_event.kind=k;last_event.target=t;last_event.sequence=s;}
static ZeroNode *next_node(ZeroNode *n){if(!n)return NULL;if(n->child_count)return n->children[0];while(n->parent){ZeroNode *p=n->parent;for(size_t i=0;i+1<p->child_count;i++)if(p->children[i]==n)return p->children[i+1];n=p;}return NULL;}
static ZeroNode *prev_node(ZeroNode *n){if(!n||!n->parent)return NULL;ZeroNode *p=n->parent;for(size_t i=0;i<p->child_count;i++)if(p->children[i]==n){if(i==0)return p;n=p->children[i-1];while(n->child_count)n=n->children[n->child_count-1];return n;}return p;}
void zero_core_init(ZeroCore *c,ZeroNode *r){c->root=r;c->focus=NULL;c->next_sequence=1;emit(ZERO_EVENT_DOCUMENT_START,r?r->id:0,c->next_sequence++);}
int zero_focus_first(ZeroCore *c){if(!c||!c->root)return 0;if(c->focus)emit(ZERO_EVENT_FOCUS_LEAVE,c->focus->id,c->next_sequence++);c->focus=c->root;if(c->root->child_count)c->focus=c->root->children[0];emit(ZERO_EVENT_FOCUS_ENTER,c->focus->id,c->next_sequence++);return 1;}
int zero_focus_next(ZeroCore *c){if(!c)return 0;if(!c->focus)return zero_focus_first(c);ZeroNode*n=next_node(c->focus);if(!n)return 0;emit(ZERO_EVENT_FOCUS_LEAVE,c->focus->id,c->next_sequence++);c->focus=n;emit(ZERO_EVENT_FOCUS_ENTER,n->id,c->next_sequence++);return 1;}
int zero_focus_previous(ZeroCore *c){if(!c)return 0;if(!c->focus)return zero_focus_first(c);ZeroNode*n=prev_node(c->focus);if(!n)return 0;emit(ZERO_EVENT_FOCUS_LEAVE,c->focus->id,c->next_sequence++);c->focus=n;emit(ZERO_EVENT_FOCUS_ENTER,n->id,c->next_sequence++);return 1;}
const ZeroEvent *zero_last_event(void){return &last_event;}
