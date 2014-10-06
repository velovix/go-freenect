package freenect

/*
#include <libfreenect/libfreenect.h>

extern void depthCallback(freenect_device *device, void *depth, uint32_t timestamp);
extern void videoCallback(freenect_device *device, void *video, uint32_t timestamp);
*/
import "C"
import "errors"
import "unsafe"
import "encoding/binary"
import "fmt"

var devices map[*C.freenect_device]*Device

type Device struct {
	device        *C.freenect_device
	depthCallback DepthCallback
	videoCallback VideoCallback
}

type DepthCallback func(device *Device, depth []uint16, timestamp uint32)
type VideoCallback func(device *Device, video []byte, timestamp uint32)

//export depthCallbackInterceptor
func depthCallbackInterceptor(device *C.freenect_device, depth *uint16, timestamp uint32) {

	if val, ok := devices[device]; ok && val.depthCallback != nil {
		attributes := C.freenect_get_current_depth_mode(devices[device].device)

		depthSliceData := C.GoBytes(unsafe.Pointer(depth), C.int(attributes.bytes))
		depthSlice := make([]uint16, int(attributes.width)*int(attributes.height))
		for i := 0; i < len(depthSliceData); i += 2 {
			depthSlice[i/2] = uint16(binary.LittleEndian.Uint16(depthSliceData[i : i+2]))
		}
		val.depthCallback(val, depthSlice, timestamp)
	}
}

//export videoCallbackInterceptor
func videoCallbackInterceptor(device *C.freenect_device, video *byte, timestamp uint32) {

	if val, ok := devices[device]; ok && val.videoCallback != nil {
		attributes := C.freenect_get_current_video_mode(devices[device].device)

		videoSlice := C.GoBytes(unsafe.Pointer(video), C.int(attributes.bytes))
		val.videoCallback(val, videoSlice, timestamp)
	}
}

func (device *Device) Destroy() {

	C.freenect_close_device(device.device)
}

func (device *Device) SetDepthCallback(callback DepthCallback) {

	if _, ok := devices[device.device]; !ok {
		devices[device.device] = device
	}

	C.freenect_set_depth_callback(device.device, (*[0]byte)(C.depthCallback))
	device.depthCallback = callback
}

func (device *Device) SetVideoCallback(callback VideoCallback) {

	if _, ok := devices[device.device]; !ok {
		devices[device.device] = device
	}

	C.freenect_set_video_callback(device.device, (*[0]byte)(C.videoCallback))
	device.videoCallback = callback
}

func (device *Device) StartDepthStream(resolution Resolution, format DepthFormat) error {

	errCode := C.freenect_set_depth_mode(device.device,
		C.freenect_find_depth_mode(C.freenect_resolution(resolution), C.freenect_depth_format(format)))

	if errCode != 0 {
		return errors.New("could not find depth mode")
	}

	errCode = C.freenect_start_depth(device.device)

	if errCode != 0 {
		return errors.New("could not start depth stream")
	}

	return nil
}

func (device *Device) StopDepthStream() error {

	errCode := C.freenect_stop_depth(device.device)

	if errCode != 0 {
		return errors.New("could not stop depth stream")
	}

	return nil
}

func (device *Device) StartVideoStream(resolution Resolution, format VideoFormat) error {

	errCode := C.freenect_set_video_mode(device.device,
		C.freenect_find_video_mode(C.freenect_resolution(resolution), C.freenect_video_format(format)))

	if errCode != 0 {
		return errors.New("could not find video mode")
	}

	errCode = C.freenect_start_video(device.device)

	if errCode != 0 {
		return errors.New("could not start video stream")
	}

	return nil
}

func (device *Device) StopVideoStream() error {

	errCode := C.freenect_stop_video(device.device)

	if errCode != 0 {
		return errors.New("could not stop video stream")
	}

	return nil
}

func (device *Device) SetLED(color LEDColor) error {

	errCode := C.freenect_set_led(device.device, C.freenect_led_options(color))

	if errCode != 0 {
		return errors.New("could not set LED color")
	}

	return nil
}

func (device *Device) GetTilt() float64 {

	return 0
}

func (device *Device) SetTilt(degrees float64) error {

	errCode := C.freenect_set_tilt_degs(device.device, C.double(degrees))

	if errCode < 0 {
		return fmt.Errorf("could not set tilt to ", degrees)
	}

	return nil
}
