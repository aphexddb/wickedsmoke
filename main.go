// dependencies managed via https://github.com/kardianos/govendor
package main

// TODO's
// - maybe charting https://github.com/wcharczuk/go-chart

import (
	"flag"
	"log"
	"net/http"
	"time"
	// "github.com/mrmorphic/hwio"
)

var address = flag.String("address", ":8080", "http service address")
var tickMs = flag.Int("tickMs", 1000, "interval in ms to broadcast cook data to clients")
var hwPollMs = flag.Int("hwPollMs", 1000, "interval in ms to read values from hardware")

// serveHomepage returns an HTTP page
func serveHomepage(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		log.Printf("HTTP 404 %s\n", r.URL)
		http.Error(w, "Not found", 404)
		return
	}
	if r.Method != "GET" {
		http.Error(w, "Method not allowed", 405)
		return
	}
	log.Printf("HTTP GET %s\n", r.URL)
	http.ServeFile(w, r, "index.html")
}

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
				cook.SyncFromHardware()
				hub.Broadcast(cook.ToJSON())
			case <-quit:
				ticker.Stop()
				return
			}
		}
	}()
	defer close(quit)

	// create a data channel for reading hardware, then start polling hardware
	cHardware := make(chan *Hardware)
	defer close(cHardware)
	go hw.Read(*hwPollMs, cHardware)

	// poll data channel
	go func() {
		for {
			// update hardware in our cook
			hw := <-cHardware
			cook.UpdateProbeTemps(hw)

			// assume HW is ok if we are reading data
			cook.SetHardwareStatus(true)

			time.Sleep(time.Duration(*hwPollMs) * time.Millisecond)
		}
	}()

	// Start HTTP server
	log.Printf("HTTP listening on: %s\n", *address)
	http.HandleFunc("/", serveHomepage)
	http.HandleFunc("/ws", func(w http.ResponseWriter, r *http.Request) {
		serveWs(hub, w, r)
	})
	err := http.ListenAndServe(*address, nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
