package main

import (
	"encoding/json"
	"fmt"
	"log"
	"time"
)

// TargetTemp is used to set a target temprature
type TargetTemp struct {
	Channel    int     `json:"channel"`
	TargetTemp float32 `json:"targetTemp"`
}

// CookProbe represents a physical temprature probe used in cooking
type CookProbe struct {
	Channel       int     `json:"channel"`
	Celsius       float32 `json:"celsius"`
	Fahrenheit    float32 `json:"fahrenheit"`
	Voltage       float32 `json:"voltage"`
	TargetTemp    float32 `json:"targetTemp"`
	TargetReached bool    `json:"targetReached"`
}

// SetVoltage updates the raw voltage value of a probe
func (p *CookProbe) SetVoltage(voltage float32) {
	p.Voltage = voltage
}

// SetTemperature updates the temprature of a probe
func (p *CookProbe) SetTemperature(celsius float32) {
	p.Celsius = celsius
	p.Fahrenheit = CelsiusToFahrenheit(celsius)

	// check for target temprature
	if p.TargetTemp > 0 && celsius > p.TargetTemp {
		p.TargetReached = true
	} else {
		p.TargetReached = false
	}
}

// Cook represents cooking parameters and associated probes
type Cook struct {
	HardwareStatus string      `json:"hardwareStatus"`
	HardwareOK     bool        `json:"hardwareOK"`
	Label          string      `json:"label"`
	UptimeSince    time.Time   `json:"uptimeSince"`
	StartTime      *time.Time  `json:"startTime"`
	StopTime       *time.Time  `json:"stopTime"`
	Cooking        bool        `json:"cooking"`
	Probes         []CookProbe `json:"cookProbes"`
}

// NewCook creates a new instance of the cook
func NewCook() *Cook {
	c := &Cook{
		Label:          "Cook",
		UptimeSince:    time.Now(),
		HardwareStatus: "Starting",
	}

	// create cook probes
	c.Probes = append(c.Probes, CookProbe{
		Channel:    0,
		TargetTemp: -1,
	})
	c.Probes = append(c.Probes, CookProbe{
		Channel:    1,
		TargetTemp: -1,
	})
	c.Probes = append(c.Probes, CookProbe{
		Channel:    2,
		TargetTemp: -1,
	})
	c.Probes = append(c.Probes, CookProbe{
		Channel:    3,
		TargetTemp: -1,
	})
	return c
}

// SetHardwareStatus sets the hardware status
func (c *Cook) SetHardwareStatus(ok bool, status string) {
	c.HardwareOK = ok
	c.HardwareStatus = status
}

// Start starts a cook
func (c *Cook) Start() {
	now := time.Now()
	c.Cooking = true
	c.StartTime = &now
	c.StopTime = nil
	log.Println("Cooking started")
}

// Stop starts a cook
func (c *Cook) Stop() {
	now := time.Now()
	c.Cooking = false
	c.StopTime = &now
	log.Println("Cooking stopped")
}

// SetTargetTemp updates target temprature for a channel
func (c *Cook) SetTargetTemp(channel int, temp float32) error {
	if len(c.Probes) < channel {
		return fmt.Errorf("No probe at Channel %v", channel)
	}
	c.Probes[channel].TargetTemp = temp
	return nil
}

// ToJSON returns the current cook values to JSON
func (c *Cook) ToJSON() []byte {
	bytes, _ := json.Marshal(c)
	return bytes
}
