package temper

import (
	"fmt"
	"log"

	"github.com/google/gousb"
)

type DeviceTemperature struct {
	Device      *gousb.Device
	Temperature float64
}

// GetTemperature is a function that tries to find first USB TEMPer device on maching and read
// the current temperature from there
// Returns temperature in Celcius and error if it occured
func GetTemperature() (DeviceTemperature, error) {
	ctx := gousb.NewContext()
	defer ctx.Close()

	// Iterate through available Devices, finding all that match a known VID/PID.
	vid, pid := gousb.ID(0x0c45), gousb.ID(0x7401)
	devs, err := ctx.OpenDevices(func(desc *gousb.DeviceDesc) bool {
		// this function is called for every device present.
		// Returning true means the device should be opened.
		return desc.Vendor == vid && desc.Product == pid
	})

	for _, d := range devs {
		defer d.Close()
	}
	if err != nil {
		log.Fatalf("OpenDevices(): %v", err)
		return DeviceTemperature{}, err
	}
	if len(devs) == 0 {
		log.Fatalf("no devices found matching VID %s and PID %s", vid, pid)
		return DeviceTemperature{}, fmt.Errorf("no devices found matching VID %s and PID %s", vid, pid)
	}

	// Pick the first device found.
	dev := devs[0]

	return getDeviceTemperature(dev)
}

// GetTemperatures is a function that tries to find all USB TEMPer devices on maching and read
// the current temperatures from all devices
// Returns temperatures in Celcius and error if it occured
func GetTemperatures() ([]DeviceTemperature, error) {
	ctx := gousb.NewContext()
	defer ctx.Close()

	// Iterate through available Devices, finding all that match a known VID/PID.
	vid, pid := gousb.ID(0x0c45), gousb.ID(0x7401)
	devs, err := ctx.OpenDevices(func(desc *gousb.DeviceDesc) bool {
		// this function is called for every device present.
		// Returning true means the device should be opened.
		return desc.Vendor == vid && desc.Product == pid
	})

	for _, d := range devs {
		defer d.Close()
	}
	if err != nil {
		log.Fatalf("OpenDevices(): %v", err)
		return []DeviceTemperature{}, err
	}
	if len(devs) == 0 {
		log.Fatalf("no devices found matching VID %s and PID %s", vid, pid)
		return []DeviceTemperature{}, fmt.Errorf("no devices found matching VID %s and PID %s", vid, pid)
	}

	var temperatures []DeviceTemperature

	for _, dev := range devs {
		temp, err := getDeviceTemperature(dev)
		if err == nil {
			temperatures = append(temperatures, temp)
		}
	}
	return temperatures, nil
}

func getDeviceTemperature(dev *gousb.Device) (DeviceTemperature, error) {
	// Switch the configuration to #1.
	cfg, err := dev.Config(1)
	if err != nil {
		log.Fatalf("%s.Config(1): %v", dev, err)
		return DeviceTemperature{}, err
	}
	defer cfg.Close()

	// In the config #1, claim interface #1 with alt setting #0.
	intf, err := cfg.Interface(1, 0)
	if err != nil {
		log.Fatalf("%s.Interface(1, 0): %v", cfg, err)
		return DeviceTemperature{}, err
	}
	defer intf.Close()

	// In this interface open endpoint 0x82 for reading.
	epIn, err := intf.InEndpoint(0x82)
	if err != nil {
		log.Fatalf("%s.InEndpoint(0x82): %v", intf, err)
		return DeviceTemperature{}, err
	}

	// Send device a control request with standard parameters and data
	_, err = dev.Control(
		0x21, 0x09, 0x0200, 0x01, []byte{0x01, 0x80, 0x33, 0x01, 0x00, 0x00, 0x00, 0x00},
	)
	if err != nil {
		return DeviceTemperature{}, err
	}

	// Read response from USB temper device
	buf := make([]byte, 8)
	if _, err = epIn.Read(buf); err != nil {
		return DeviceTemperature{}, err
	}

	return DeviceTemperature{Device: dev, Temperature: float64(buf[2]) + float64(buf[3])/256}, nil
}
