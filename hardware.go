package main

import (
	"encoding/json"
	"log"
	"math/rand"
	"time"
)

var r = rand.New(rand.NewSource(99))

// Probe represents a physical temprature probe
type Probe struct {
	Voltage    float32 `json:"voltage"`
	Celsius    float32 `json:"c"`
	Fahrenheit float32 `json:"f"`
}

// Hardware represents all of the the temprature probes
type Hardware struct {
	adc    *RaspPiOAnalog
	Probe0 Probe `json:"probe0"`
	Probe1 Probe `json:"probe1"`
	Probe2 Probe `json:"probe2"`
	Probe3 Probe `json:"probe3"`
}

// NewHardware creates a new instance of the hardware
func NewHardware() *Hardware {
	log.Println("hardware: TODO: NewHardware -> use GPIO")
	return &Hardware{
		adc:    NewRaspPiOAnalog(),
		Probe0: Probe{Voltage: 0, Celsius: 0, Fahrenheit: 0},
		Probe1: Probe{Voltage: 0, Celsius: 0, Fahrenheit: 0},
		Probe2: Probe{Voltage: 0, Celsius: 0, Fahrenheit: 0},
		Probe3: Probe{Voltage: 0, Celsius: 0, Fahrenheit: 0},
	}
}

// Read checks the hardware on a cycle and updates the channel
func (h *Hardware) Read(intervalMs int, c chan *Hardware) {
	log.Printf("hardware: Reading hardware every %vms\n", intervalMs)
	for {
		// ask adc hardware for voltages
		channels, err := h.adc.GetChannels()
		if err != nil {
			log.Fatalf("hardware: error reading channels: %s\n", err)
		} else {
			for i := 0; i < len(channels); i++ {
				log.Printf("adc channel %v voltage: %v\n", i, channels[i])

				// TODO - update probe data
			}
		}

		log.Println("hardware: TODO: Read -> get real HW values")
		h.Probe0.Celsius = r.Float32() * 100
		h.Probe1.Celsius = r.Float32() * 100
		h.Probe2.Celsius = r.Float32() * 100
		h.Probe3.Celsius = r.Float32() * 100

		// convert all °C tempratures to °F
		h.Probe0.Fahrenheit = CelsiusToFahrenheit(h.Probe0.Celsius)
		h.Probe1.Fahrenheit = CelsiusToFahrenheit(h.Probe1.Celsius)
		h.Probe2.Fahrenheit = CelsiusToFahrenheit(h.Probe2.Celsius)
		h.Probe3.Fahrenheit = CelsiusToFahrenheit(h.Probe3.Celsius)

		c <- h
		time.Sleep(time.Duration(intervalMs) * time.Millisecond)
	}
}

// CelsiusToFahrenheit converts celsius to fahrenheit
// T(°F) = T(°C) × 9/5 + 32
func CelsiusToFahrenheit(celsius float32) float32 {
	return celsius*9/5 + 32
}

// FahrenheitToCelsius converts fahrenheit to celsius
// T(°C) = (T(°F) - 32) × 5/9
func FahrenheitToCelsius(fahrenheit float32) float32 {
	return (fahrenheit - 32) * 5 / 9
}

// ToJSON returns the current probe values to JSON
func (h *Hardware) ToJSON() []byte {
	bytes, _ := json.Marshal(h)
	return bytes
}
