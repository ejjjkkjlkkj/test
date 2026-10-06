#ifndef ZERO_MEDIA_H
#define ZERO_MEDIA_H

#include <stddef.h>
#include "zero_core.h"

typedef enum {
    ZERO_MEDIA_IMAGE = 1,
    ZERO_MEDIA_PHOTO,
    ZERO_MEDIA_GRAPHIC,
    ZERO_MEDIA_CHART,
    ZERO_MEDIA_DIAGRAM,
    ZERO_MEDIA_ICON,
    ZERO_MEDIA_CAPTCHA,
    ZERO_MEDIA_UNKNOWN
} ZeroMediaKind;

typedef enum {
    ZERO_OBSERVATION_TEXT_REGION = 1,
    ZERO_OBSERVATION_OBJECT,
    ZERO_OBSERVATION_RELATION,
    ZERO_OBSERVATION_AXIS,
    ZERO_OBSERVATION_SERIES,
    ZERO_OBSERVATION_POINT,
    ZERO_OBSERVATION_LABEL,
    ZERO_OBSERVATION_CONTROL,
    ZERO_OBSERVATION_PROTECTED
} ZeroObservationKind;

typedef struct {
    ZeroObservationKind kind;
    ZeroMediaKind media_kind;
    unsigned region_id;
    double confidence;
    int uncertain;
    const char *description;
} ZeroObservation;

typedef struct {
    ZeroMediaKind kind;
    const char *source;
    ZeroObservation *observations;
    size_t observation_count;
    size_t observation_capacity;
    int protected_interaction;
} ZeroMedia;

void zero_media_init(ZeroMedia *media, ZeroMediaKind kind, const char *source,
                     ZeroObservation *storage, size_t capacity);
int zero_media_add(ZeroMedia *media, ZeroObservation observation);
int zero_media_is_protected(const ZeroMedia *media);

#endif
