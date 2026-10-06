#include <assert.h>
#include "zero_media.h"

int main(void)
{
    ZeroObservation storage[4];
    ZeroMedia media;

    zero_media_init(&media, ZERO_MEDIA_CHART, "chart.png", storage, 4);

    assert(!zero_media_is_protected(&media));
    assert(zero_media_add(&media, (ZeroObservation){
        ZERO_OBSERVATION_AXIS, ZERO_MEDIA_CHART, 1, 0.99, 0, "horizontal axis"
    }));
    assert(zero_media_add(&media, (ZeroObservation){
        ZERO_OBSERVATION_SERIES, ZERO_MEDIA_CHART, 2, 0.91, 1, "series A"
    }));
    assert(media.observation_count == 2);

    zero_media_init(&media, ZERO_MEDIA_CAPTCHA, "challenge.png", storage, 4);
    assert(zero_media_is_protected(&media));

    return 0;
}
