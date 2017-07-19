package main

import (
	"encoding/json"
	"fmt"
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
	TargetTemp    float32 `json:"targetTemp"`
	TargetReached bool    `json:"targetReached"`
}

// Cook represents cooking parameters and associated probes
type Cook struct {
	hardware       *Hardware
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
func NewCook(hardware *Hardware) *Cook {
	c := &Cook{
		hardware:       hardware,
		Label:          "Cook",
		UptimeSince:    time.Now(),
		HardwareStatus: "Starting",
	}
	c.Probes = append(c.Probes, CookProbe{
		Celsius: c.hardware.Probe0.Celsius,
		Channel: 0,
	})
	c.Probes = append(c.Probes, CookProbe{
		Celsius: c.hardware.Probe1.Celsius,
		Channel: 1,
	})
	c.Probes = append(c.Probes, CookProbe{
		Celsius: c.hardware.Probe2.Celsius,
		Channel: 2,
	})
	c.Probes = append(c.Probes, CookProbe{
		Celsius: c.hardware.Probe3.Celsius,
		Channel: 3,
	})
	return c
}

// SetLabel sets a label for the cook
func (c *Cook) SetLabel(label string) {
	c.Label = label
}

// SetHardwareStatus sets the hardware status
func (c *Cook) SetHardwareStatus(ok bool) {
	c.HardwareOK = ok
	if ok {
		c.HardwareStatus = "OK"
	} else {
		c.HardwareStatus = "Not OK"
	}
}

// Start starts a cook
func (c *Cook) Start() {
	now := time.Now()
	c.Cooking = true
	c.StartTime = &now
}

// Stop starts a cook
func (c *Cook) Stop() {
	now := time.Now()
	c.Cooking = false
	c.StopTime = &now
}

// SyncFromHardware updates probe temps from Hardware
func (c *Cook) SyncFromHardware() {
	c.Probes[0].Celsius = c.hardware.Probe0.Celsius
	c.Probes[1].Celsius = c.hardware.Probe1.Celsius
	c.Probes[2].Celsius = c.hardware.Probe2.Celsius
	c.Probes[3].Celsius = c.hardware.Probe3.Celsius
}

// SetTargetTemp updates target temprature for a channel
func (c *Cook) SetTargetTemp(channel int, temp float32) error {
	if len(c.Probes) < channel {
		return fmt.Errorf("No probe at Channel %v", channel)
	}
	c.Probes[channel].TargetTemp = temp
	return nil
}

// UpdateProbeTemps sets the hardware
func (c *Cook) UpdateProbeTemps(hw *Hardware) {
	c.hardware.Probe0.Voltage = hw.Probe0.Voltage
	c.hardware.Probe0.Celsius = hw.Probe0.Celsius

	c.hardware.Probe1.Voltage = hw.Probe1.Voltage
	c.hardware.Probe1.Celsius = hw.Probe1.Celsius

	c.hardware.Probe2.Voltage = hw.Probe2.Voltage
	c.hardware.Probe2.Celsius = hw.Probe2.Celsius

	c.hardware.Probe3.Voltage = hw.Probe3.Voltage
	c.hardware.Probe3.Celsius = hw.Probe3.Celsius
}

// ToJSON returns the current cook values to JSON
func (c *Cook) ToJSON() []byte {
	bytes, _ := json.Marshal(c)
	return bytes
}