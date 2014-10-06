package freenect

/*
#include <libfreenect.h>
#include <sys/time.h>

extern void logCallback(freenect_context *context, freenect_loglevel level, char *message);
*/
import "C"
import "time"
import "errors"
import "fmt"
import "syscall"
import "unsafe"

var contexts map[*C.freenect_context]*Context

// Context is an object that represents a Freenect context. Many operations are
// done through this object.
type Context struct {
	context     *C.freenect_context
	logCallback LogCallback
}

// LogCallback is a function type that Freenect uses to deliver log messages to
// the user.
type LogCallback func(*Context, LogLevel, string)

//export logCallbackInterceptor
func logCallbackInterceptor(context *C.freenect_context, level C.freenect_loglevel, message *C.char) {

	if val, ok := contexts[context]; ok && val.logCallback != nil {
		val.logCallback(val, LogLevel(level), C.GoString(message))
	}
}

// NewContext creates a new Context object.
func NewContext() (Context, error) {

	var context Context

	errCode := int(C.freenect_init(&context.context, nil))

	if errCode != 0 {
		return Context{}, errors.New("could not create context")
	}

	C.freenect_select_subdevices(context.context, C.FREENECT_DEVICE_MOTOR|C.FREENECT_DEVICE_CAMERA)

	return context, nil
}

// Destroy frees the Context object data.
func (context *Context) Destroy() error {

	errCode := int(C.freenect_shutdown(context.context))

	if errCode != 0 {
		return errors.New("could not destroy context")
	}

	return nil
}

// DeviceCount returns the number of devices connected to the system.
func (context *Context) DeviceCount() (int, error) {

	count := int(C.freenect_num_devices(context.context))

	if count < 0 {
		return 0, errors.New("error getting device count")
	}

	return count, nil
}

// OpenDevice opens a device of the given index and returns the corresponding
// Device object.
func (context *Context) OpenDevice(index int) (Device, error) {

	var device Device

	errCode := C.freenect_open_device(context.context, &device.device, C.int(index))

	if errCode != 0 {
		return Device{}, fmt.Errorf("could not open device", index)
	}

	return device, nil
}

// ProcessEvents looks for new data from the device or devices and sends the
// data to the callbacks.
func (context *Context) ProcessEvents(timeout time.Duration) error {

	var errCode int

	if timeout == 0 {
		errCode = int(C.freenect_process_events(context.context))
	} else {
		timeval := syscall.NsecToTimeval(timeout.Nanoseconds())
		cDuration := (*C.struct_timeval)(unsafe.Pointer(&timeval))
		errCode = int(C.freenect_process_events_timeout(context.context, cDuration))
	}

	if errCode != 0 {
		return errors.New("could not process events")
	}

	return nil
}

// SetLogLevel sets how verbose Freenect should be.
func (context *Context) SetLogLevel(level LogLevel) {

	C.freenect_set_log_level(context.context, C.freenect_loglevel(level))
}

// SetLogCallback sets the callback function Freenect will call when it has a
// log message. If this callback is set, Freenect will not automatically print
// log messages.
func (context *Context) SetLogCallback(callback LogCallback) {

	if _, ok := contexts[context.context]; !ok {
		contexts[context.context] = context
	}

	C.freenect_set_log_callback(context.context, (*[0]byte)(C.logCallback))
	context.logCallback = callback
}
