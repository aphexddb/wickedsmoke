package main

import (
	"flag"
	"log"
	"time"
)

var address = flag.String("address", ":8080", "http service address")
var tickMs = flag.Int("tickMs", 1000, "interval in ms to broadcast cook data to clients")

func main() {
	flag.Parse()
	log.SetFlags(log.Lshortfile)
	log.Println("Wickedsmoke starting")

	// Create new connection to I2C bus on 2 line with address 0x27
	i2c, i2cErr := NewI2C(0x27, 2)
	if i2cErr != nil {
		log.Fatalln(i2cErr)
	}
	// Free I2C connection on exit
	defer i2c.Close()

	// create new hardware, cook and message hub
	hardware := NewHardware(i2c)
	cook := NewCook()
	hub := NewHub()

	// start message hub
	go hub.Run()

	// broadcast the cook state every tick
	log.Printf("Broadcasting cook data to clients every %vms\n", *tickMs)
	ticker := time.NewTicker(time.Duration(*tickMs) * time.Millisecond)
	quit := make(chan struct{})
	go func() {
		for {
			select {
			case <-ticker.C:
				// update all temp probes
				cook.Probes[0].SetTemperature(hardware.Probe0.Celsius)
				cook.Probes[1].SetTemperature(hardware.Probe1.Celsius)
				cook.Probes[2].SetTemperature(hardware.Probe2.Celsius)
				cook.Probes[3].SetTemperature(hardware.Probe3.Celsius)

				// update all temp probes with voltages
				cook.Probes[0].SetVoltage(hardware.Probe0.Voltage)
				cook.Probes[1].SetVoltage(hardware.Probe1.Voltage)
				cook.Probes[2].SetVoltage(hardware.Probe2.Voltage)
				cook.Probes[3].SetVoltage(hardware.Probe3.Voltage)

				// assume HW is ok if we are updating
				cook.SetHardwareStatus(true, "OK")

				hub.Broadcast(cook.ToJSON())
			case <-quit:
				ticker.Stop()
				return
			}
		}
	}()
	defer close(quit)

	// Start HTTP server
	httpErr := ServeHTTP(cook, hub, address)
	if httpErr != nil {
		log.Fatal("http error: ", httpErr)
	}
}
