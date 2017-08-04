package main

import (
	"encoding/json"
	"log"
	"math/rand"
)

// TODO nuke this
var r = rand.New(rand.NewSource(99))

// Probe represents a physical temprature probe
type Probe struct {
	Voltage float32 `json:"voltage"`
	Celsius float32 `json:"celsius"`
}

// Hardware represents all of the the temprature probes
type Hardware struct {
	adc    *ADC
	Probe0 Probe `json:"probe0"`
	Probe1 Probe `json:"probe1"`
	Probe2 Probe `json:"probe2"`
	Probe3 Probe `json:"probe3"`
}

// NewHardware creates a new instance of the hardware
func NewHardware(adc *MCP3008) *Hardware {
	log.Println("hardware: TODO: NewHardware -> use NewRaspPiOAnalog???")

	return &Hardware{
		adc:    adc,
		Probe0: Probe{Voltage: 0, Celsius: 0},
		Probe1: Probe{Voltage: 0, Celsius: 0},
		Probe2: Probe{Voltage: 0, Celsius: 0},
		Probe3: Probe{Voltage: 0, Celsius: 0},
	}
}

// Read gets the current value of probe data
func (h *Hardware) Read() *Hardware {
	log.Println("hardware: Reading hardware")

	// read from adc
	h.adc.Read()

	h.Probe0.Voltage = float32(h.adc.values[0])
	h.Probe1.Voltage = float32(h.adc.values[1])
	h.Probe2.Voltage = float32(h.adc.values[2])
	h.Probe3.Voltage = float32(h.adc.values[3])
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
