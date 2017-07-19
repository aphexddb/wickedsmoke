// dependencies managed via https://github.com/kardianos/govendor
package main

// TODO's
// - maybe charting https://github.com/wcharczuk/go-chart

import (
	"flag"
	"log"
	"time"
	// "github.com/mrmorphic/hwio"
)

var address = flag.String("address", ":8080", "http service address")
var tickMs = flag.Int("tickMs", 1000, "interval in ms to broadcast cook data to clients")

func main() {
	flag.Parse()
	log.SetFlags(log.Lshortfile)
	log.Println("Wickedsmoke starting")

	// create new hardware and cook
	hw := NewHardware()
	cook := NewCook(hw)

	// create the message hub
	hub := NewHub()
	go hub.Run()

	// broadcast the cook state every tick
	log.Printf("Broadcasting cook data to clients every %vms\n", *tickMs)
	ticker := time.NewTicker(time.Duration(*tickMs) * time.Millisecond)
	quit := make(chan struct{})
	go func() {
		for {
			select {
			case <-ticker.C:
				cook.UpdateFromHw(hw.Read())
				hub.Broadcast(cook.ToJSON())
			case <-quit:
				ticker.Stop()
				return
			}
		}
	}()
	defer close(quit)

	// Start HTTP server
	err := ServeHTTP(cook, hub, address)
	if err != nil {
		log.Fatal("http error: ", err)
	}
}
