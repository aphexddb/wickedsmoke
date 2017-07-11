// dependencies managed via https://github.com/kardianos/govendor
package main

import (
	"flag"
	"log"
	"net/http"
	"time"
	// "github.com/mrmorphic/hwio"
)

var address = flag.String("address", ":8080", "http service address")
var hwPoll = flag.Int("hwPoll", 1000, "interval in ms to read values from hardware")
var broadcastPoll = flag.Int("broadcastPoll", 1000, "interval in ms to broadcast data")

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

	// create the message hub
	hub := NewHub()
	go hub.Run()

	// create a channel for reading hardware, then start polling hardware
	c := make(chan *Hardware)
	defer close(c)
	hw := NewHardware()
	go hw.Read(*hwPoll, c)

	// poll data channel
	go func() {
		for {
			// broadcast to hub
			hw := <-c
			hub.Broadcast(hw.ToJSON())
			time.Sleep(time.Duration(*broadcastPoll) * time.Millisecond)
		}
	}()

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
