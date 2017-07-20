package main

import (
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

// SetTempRequest is a request to set temprature for a probe
type SetTempRequest struct {
	Temp float32 `json:"temp"`
}

// CmdResponse is a response to requests
type CmdResponse struct {
	Message string `json:"message"`
}

// serveHomepage returns an HTTP page
func serveHomepage(w http.ResponseWriter, r *http.Request) {
	// homepage
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

// setTargetTemp sets a target temp for a probe channel
func setTargetTemp(cook *Cook, w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	channel, ok := vars["channel"]
	if !ok {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Printf("missing channel")
		http.NotFound(w, r)
		return
	}

	body, _ := ioutil.ReadAll(r.Body)
	setTempRequest := &SetTempRequest{}
	err := json.Unmarshal(body, setTempRequest)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Println(err)
		return
	}

	c, _ := strconv.Atoi(channel)
	cook.Probes[c].TargetTemp = setTempRequest.Temp

	cmdResp := &CmdResponse{
		Message: fmt.Sprintf("Target temp on channel %s set to %v", channel, setTempRequest.Temp),
	}
	resp, _ := json.Marshal(cmdResp)
	io.WriteString(w, string(resp))
}

// middleware enables cross origin requests
func middleware(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		if r.Method == "OPTIONS" {
			w.Header().Set("Access-Control-Allow-Methods", "POST, GET")
			w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Access-Control-Allow-Headers, Authorization, X-Requested-With")
			return
		}
		h.ServeHTTP(w, r)
	})
}

// ServeHTTP starts a webs server
func ServeHTTP(cook *Cook, hub *Hub, address *string) error {
	log.Printf("HTTP listening on: %s\n", *address)

	r := mux.NewRouter()

	r.HandleFunc("/", serveHomepage).Methods("GET")

	r.HandleFunc("/probe/{channel}/setTargetTemp", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		setTargetTemp(cook, w, r)
	})).Methods("POST")

	r.HandleFunc("/cook/start", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		cook.Start()
		cmdResp := &CmdResponse{
			Message: "Cook started",
		}
		resp, _ := json.Marshal(cmdResp)
		io.WriteString(w, string(resp))
	})).Methods("POST")
	r.HandleFunc("/cook/stop", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		cook.Stop()
		cmdResp := &CmdResponse{
			Message: "Cook stopped",
		}
		resp, _ := json.Marshal(cmdResp)
		io.WriteString(w, string(resp))
	})).Methods("POST")

	r.HandleFunc("/ws", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		serveWs(hub, w, r)
	}))

	http.Handle("/", r)
	return http.ListenAndServe(*address, middleware(r))
}
