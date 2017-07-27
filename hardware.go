package main

import (
	"encoding/json"
	"log"
	"math/rand"
)

var r = rand.New(rand.NewSource(99))

// Probe represents a physical temprature probe
type Probe struct {
	Voltage float32 `json:"voltage"`
	Celsius float32 `json:"celsius"`
}

// Hardware represents all of the the temprature probes
type Hardware struct {
	// adc    *RaspPiOAnalog
	i2c    *I2C
	Probe0 Probe `json:"probe0"`
	Probe1 Probe `json:"probe1"`
	Probe2 Probe `json:"probe2"`
	Probe3 Probe `json:"probe3"`
}

// NewHardware creates a new instance of the hardware
func NewHardware(i2c *I2C) *Hardware {
	log.Println("hardware: TODO: NewHardware -> use NewRaspPiOAnalog???")

	return &Hardware{
		// adc:    NewRaspPiOAnalog(),
		i2c:    i2c,
		Probe0: Probe{Voltage: 0, Celsius: 0},
		Probe1: Probe{Voltage: 0, Celsius: 0},
		Probe2: Probe{Voltage: 0, Celsius: 0},
		Probe3: Probe{Voltage: 0, Celsius: 0},
	}
}

// Read gets the current value of probe data
func (h *Hardware) Read() *Hardware {
	log.Println("hardware: Reading hardware")

	// // ask adc hardware for voltages
	// channels, err := h.adc.GetChannels()
	// if err != nil {
	// 	// log.Fatalf("hardware: error reading channels: %s\n", err)
	// 	log.Printf("hardware: error reading adc channels: %s\n", err)
	// } else {
	// 	for i := 0; i < len(channels); i++ {
	// 		log.Printf("adc channel %v voltage: %v\n", i, channels[i])

	// 		// TODO - update probe data
	// 	}
	// }

	// read from i2c bus
	var buffer []byte
	bytes, readErr := h.i2c.Read(buffer)
	if readErr != nil {
		log.Fatalln(readErr)
	}
	log.Println(bytes)

	log.Println("hardware: TODO: Read -> get real HW voltage values, not random")
	h.Probe0.Celsius = r.Float32() * 100
	h.Probe1.Celsius = r.Float32() * 100
	h.Probe2.Celsius = r.Float32() * 100
	h.Probe3.Celsius = r.Float32() * 100

	return h
}

// ToJSON returns the current probe values to JSON
func (h *Hardware) ToJSON() []byte {
	bytes, _ := json.Marshal(h)
	return bytes
}
