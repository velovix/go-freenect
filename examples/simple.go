// This is a simple example of a program that uses go-freenect. The program
// initializes a Kinect and counts how many depth frames it recieves in 10
// seconds.
package main

import (
	"fmt"
	"github.com/velovix/go-freenect"
	"os"
	"time"
)

var (
	context  freenect.Context
	kinect   freenect.Device
	frameCnt int
)

func onDepthFrame(device *freenect.Device, depth []uint16, timestamp uint32) {

	frameCnt++
}

func init() {

	var err error

	// Create a freenect context.
	context, err = freenect.NewContext()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	// Set the log level to be fairly verbose.
	context.SetLogLevel(freenect.LogDebug)

	// Open the kinect device
	if cnt, _ := context.DeviceCount(); cnt == 0 {
		fmt.Println("could not find any devices")
		os.Exit(1)
	}
	kinect, err = context.OpenDevice(0)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	// Set the depth callback. This function is called whenever a depth frame is retrieved
	kinect.SetDepthCallback(onDepthFrame)

	// Start the depth camera and begin streaming
	err = kinect.StartDepthStream(freenect.ResolutionMedium, freenect.DepthFormatMM)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func main() {

	initTime := time.Now()

	// Waits until 10 seconds have passed
	for time.Since(initTime).Seconds() < 10.0 {
		// Process freenect events
		err := context.ProcessEvents(0)
		if err != nil {
			fmt.Println(err)
			break
		}
	}

	fmt.Println("Processed", frameCnt, "frames in 10 seconds.")

	// Clean up objects
	kinect.StopDepthStream()
	kinect.Destroy()
	context.Destroy()
}
