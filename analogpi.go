package main

import (
	"log"
	"os"

	"github.com/mrmorphic/hwio"
)

const (
	// DeviceAddress is the default address for the Analog Zero Get using `i2cdetect -y`
	DeviceAddress = 0x68
)

// RaspPiOAnalog represents a RaspPiO Analog Zero
type RaspPiOAnalog struct {
	device hwio.I2CDevice
}

// NewRaspPiOAnalog creates a new RaspPiO Analog Zero device
func NewRaspPiOAnalog() *RaspPiOAnalog {
	hwio.SetDriver(new(hwio.RaspberryPiDTDriver))

	module, e := hwio.GetModule("i2c")
	if e != nil {
		log.Fatalf("analogpi: could not get i2c module: %s\n", e)
		os.Exit(1)
	}
	i2c := module.(hwio.I2CModule)
	device := i2c.GetDevice(DeviceAddress)
	return &RaspPiOAnalog{device: device}
}

// GetChannels retuns an array of
func (t *RaspPiOAnalog) GetChannels() ([8]float32, error) {
	arr := [8]float32{0, 0, 0, 0, 0, 0, 0, 0}

	buffer, e := t.device.Read(0x00, 2)
	if e != nil {
		return arr, e
	}
	log.Printf("analogpi: buffer: %v\n", buffer)

	// MSB := buffer[0]
	// LSB := buffer[1]

	// /* Convert 12bit int using two's compliment */
	// /* Credit: http://bildr.org/2011/01/tmp102-arduino/ */
	// temp := ((int(MSB) << 8) | int(LSB)) >> 4

	// // divide by 16, since lowest 4 bits are fractional.
	// return float32(temp) * 0.0625, nil

	// TODO - populate this!!!

	return arr, nil
}
