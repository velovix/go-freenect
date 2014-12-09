// Package freenect implements a Go binding for the libfreenect library.
package freenect

/*
#cgo CFLAGS: -I/usr/local/include/libfreenect -I/usr/include/libfreenect
#cgo LDFLAGS: -lfreenect
#include <libfreenect.h>
*/
import "C"

func init() {

	contexts = make(map[*C.freenect_context]*Context)
	devices = make(map[*C.freenect_device]*Device)
}
