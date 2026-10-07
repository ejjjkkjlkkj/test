#include "zero_core.h"
#include <stdint.h>
#include <stdio.h>
#include <stdlib.h>

static uint64_t mix(uint64_t x) {
    x += UINT64_C(0x9e3779b97f4a7c15);
    x = (x ^ (x >> 30)) * UINT64_C(0xbf58476d1ce4e5b9);
    x = (x ^ (x >> 27)) * UINT64_C(0x94d049bb133111eb);
    return x ^ (x >> 31);
}

static int belongs(const ZeroNode *root, const ZeroNode *needle) {
    if (!root || !needle) return 0;
    if (root == needle) return 1;
    for (size_t i = 0; i < root->child_count; ++i)
        if (belongs(root->children[i], needle)) return 1;
    return 0;
}

static int check_event(const ZeroCore *core, unsigned long before_seq) {
    const ZeroEvent *e = zero_last_event();
    if (!e) return 0;
    if (e->sequence < before_seq) return 0;
    if (core->focus && !belongs(core->root, core->focus)) return 0;
    return 1;
}

int main(void) {
    const char *shard_s = getenv("ZERO_SHARD");
    const char *count_s = getenv("ZERO_CASES_PER_SHARD");
    uint64_t shard = shard_s ? strtoull(shard_s, NULL, 10) : 0;
    uint64_t count = count_s ? strtoull(count_s, NULL, 10) : 1000000;

    ZeroNode leaf1 = {1, ZERO_LINK, "a", "", 0, NULL, NULL, 0};
    ZeroNode leaf2 = {2, ZERO_BUTTON, "b", "", 0, NULL, NULL, 0};
    ZeroNode leaf3 = {3, ZERO_TEXTBOX, "c", "", 0, NULL, NULL, 0};
    ZeroNode *children[] = {&leaf1, &leaf2, &leaf3};
    ZeroNode root = {0, ZERO_DOCUMENT, "root", "", 0, NULL, children, 3};
    leaf1.parent = leaf2.parent = leaf3.parent = &root;

    ZeroCore core;
    zero_core_init(&core, &root);
    if (zero_last_event()->kind != ZERO_EVENT_DOCUMENT_START) return 10;

    for (uint64_t i = 0; i < count; ++i) {
        uint64_t id = (shard << 24) | i;
        uint64_t r = mix(id);
        unsigned long before = core.next_sequence;

        switch ((unsigned)(r & 3u)) {
            case 0: (void)zero_focus_first(&core); break;
            case 1: (void)zero_focus_next(&core); break;
            case 2: (void)zero_focus_previous(&core); break;
            default:
                /* Reinitialization is deliberately forbidden inside a case. */
                break;
        }

        if (!check_event(&core, before) && core.next_sequence != before) {
            fprintf(stderr, "FAIL case=%llu sequence=%lu\n",
                    (unsigned long long)id, core.next_sequence);
            return 20;
        }
        if (core.next_sequence < before) {
            fprintf(stderr, "FAIL sequence regression case=%llu\n",
                    (unsigned long long)id);
            return 21;
        }
    }

    printf("PASS shard=%llu cases=%llu\n",
           (unsigned long long)shard, (unsigned long long)count);
    return 0;
}
