#include "zero_media.h"

void zero_media_init(ZeroMedia *media, ZeroMediaKind kind, const char *source,
                     ZeroObservation *storage, size_t capacity)
{
    if (!media) return;
    media->kind = kind;
    media->source = source;
    media->observations = storage;
    media->observation_count = 0;
    media->observation_capacity = capacity;
    media->protected_interaction = (kind == ZERO_MEDIA_CAPTCHA);
}

int zero_media_add(ZeroMedia *media, ZeroObservation observation)
{
    if (!media || !media->observations ||
        media->observation_count >= media->observation_capacity) {
        return 0;
    }

    media->observations[media->observation_count++] = observation;
    return 1;
}

int zero_media_is_protected(const ZeroMedia *media)
{
    return media && media->protected_interaction;
}
