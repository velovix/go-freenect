package freenect

/*
#include <libfreenect/libfreenect.h>
*/
import "C"

// LogLevel represents a logging verbosity level.
type LogLevel int

// Constants representing logging verbosity levels, ordered from least to most
// verbose.
const (
	LogFatal   = LogLevel(C.FREENECT_LOG_FATAL)
	LogError   = LogLevel(C.FREENECT_LOG_ERROR)
	LogWarning = LogLevel(C.FREENECT_LOG_WARNING)
	LogNotice  = LogLevel(C.FREENECT_LOG_NOTICE)
	LogInfo    = LogLevel(C.FREENECT_LOG_INFO)
	LogDebug   = LogLevel(C.FREENECT_LOG_DEBUG)
	LogSpew    = LogLevel(C.FREENECT_LOG_SPEW)
)

// Resolution represents a video/depth stream resolution format.
type Resolution int

// Constants representing various video/depth resolution levels.
const (
	ResolutionLow    = Resolution(C.FREENECT_RESOLUTION_LOW)
	ResolutionMedium = Resolution(C.FREENECT_RESOLUTION_MEDIUM)
	ResolutionHigh   = Resolution(C.FREENECT_RESOLUTION_HIGH)
)

// DepthFormat represents a data representation of depth information.
type DepthFormat int

// Constants representing various data representations of depth information.
const (
	DepthFormat11Bit       = DepthFormat(C.FREENECT_DEPTH_11BIT)
	DepthFormat10Bit       = DepthFormat(C.FREENECT_DEPTH_10BIT)
	DepthFormat11BitPacked = DepthFormat(C.FREENECT_DEPTH_11BIT_PACKED)
	DepthFormat10BitPacked = DepthFormat(C.FREENECT_DEPTH_10BIT_PACKED)
	DepthFormatRegistered  = DepthFormat(C.FREENECT_DEPTH_REGISTERED)
	DepthFormatMM          = DepthFormat(C.FREENECT_DEPTH_MM)
)

// VideoFormat represents a data representation of color information.
type VideoFormat int

// Constants representing various data representations of video information.
const (
	VideoFormatRGB           = VideoFormat(C.FREENECT_VIDEO_RGB)
	VideoFormatBayer         = VideoFormat(C.FREENECT_VIDEO_BAYER)
	VideoFormatIR8Bit        = VideoFormat(C.FREENECT_VIDEO_IR_8BIT)
	VideoFormatIR10Bit       = VideoFormat(C.FREENECT_VIDEO_IR_10BIT)
	VideoFormatIR10BitPacked = VideoFormat(C.FREENECT_VIDEO_IR_10BIT_PACKED)
	VideoFormatYUVRGB        = VideoFormat(C.FREENECT_VIDEO_YUV_RGB)
	VideoFormatYUVRaw        = VideoFormat(C.FREENECT_VIDEO_YUV_RAW)
)

// LEDColor represents an LED color option.
type LEDColor int

// Constants representing LED color options.
const (
	LEDColorOff            = LEDColor(C.LED_OFF)
	LEDColorGreen          = LEDColor(C.LED_GREEN)
	LEDColorRed            = LEDColor(C.LED_RED)
	LEDColorYellow         = LEDColor(C.LED_YELLOW)
	LEDColorBlinkGreen     = LEDColor(C.LED_BLINK_GREEN)
	LEDColorBlinkRedYellow = LEDColor(C.LED_BLINK_RED_YELLOW)
)

// TiltStatus represents a tilt motor status code.
type TiltStatus int

// Constants representing tilt motor status codes.
const (
	TiltStatusStopped = TiltStatus(C.TILT_STATUS_STOPPED)
	TiltStatusLimit   = TiltStatus(C.TILT_STATUS_LIMIT)
	TiltStatusMoving  = TiltStatus(C.TILT_STATUS_MOVING)
)
