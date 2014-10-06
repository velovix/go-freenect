#include "_cgo_export.h"
#include <libfreenect/libfreenect.h>

void logCallback(freenect_context *context, freenect_loglevel level, char *message) {

	logCallbackInterceptor(context, level, message);
}

void depthCallback(freenect_device *device, void *depth, uint32_t timestamp) {

	depthCallbackInterceptor(device, (uint16_t*)depth, timestamp);
}

void videoCallback(freenect_device *device, void *video, uint32_t timestamp) {

	videoCallbackInterceptor(device, (uint8_t*)video, timestamp);
}
