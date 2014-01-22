package cuda

import (
	"fmt"
	"testing"
)

func TestAll(t *testing.T) {
	err := Init(0)
	fmt.Printf("Initialized: %d\n", err)

	driverVersion, err := DriverGetVersion()
	if (err == 0) {
		fmt.Printf("Driver Version: %d\n", driverVersion)
	}

	device, err := DeviceGet(0)
	if (err == 0) {
		fmt.Printf("Device: %d\n", device)
	}

	count, err := DeviceGetCount()
	if (err == 0) {
		fmt.Printf("Device count: %d\n", count)
	}

	major, minor, err := DeviceComputeCapability(device)
	if (err == 0) {
		fmt.Printf("Compute Capability: %d.%d\n", major, minor)
	}
}
