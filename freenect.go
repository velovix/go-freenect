package freenect

/*
#cgo LDFLAGS: -lfreenect
#include <libfreenect/libfreenect.h>
*/
import "C"

func init() {

	contexts = make(map[*C.freenect_context]*Context)
	devices = make(map[*C.freenect_device]*Device)
}
