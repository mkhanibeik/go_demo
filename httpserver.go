package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"sync"
)

type MathRequest struct {
	Op    string  `json:"op"`
	Left  float64 `json:"left"`
	Right float64 `json:"right"`
}

type MathResponse struct {
	Error  string  `json:"error"`
	Result float64 `json:"result"`
}

func helloHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello World!")
}

func mathHandler(w http.ResponseWriter, r *http.Request) {
	// Decode request
	defer r.Body.Close()
	dec := json.NewDecoder(r.Body)
	req := &MathRequest{}

	if err := dec.Decode(req); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// Do work
	resp := &MathResponse{}
	switch req.Op {
	case "+":
		resp.Result = req.Left + req.Right
	case "-":
		resp.Result = req.Left - req.Right
	case "*":
		resp.Result = req.Left * req.Right
	case "/":
		if req.Right == 0 {
			resp.Error = "Division by 0"
		} else {
			resp.Result = req.Left / req.Right
		}
	}

	// Encode and return result
	w.Header().Set("Content-Type", "application/json")
	if resp.Error != "" {
		w.WriteHeader(http.StatusBadRequest)
	}

	enc := json.NewEncoder(w)
	if err := enc.Encode(resp); err != nil {
		log.Printf("cannot enccode %v - %s\n", resp, err)
	}
}

type Entry struct {
	Key   string      `json:"key"`
	Value interface{} `json:"value"`
}

var (
	keyValueDB = map[string]interface{}{}
	dbLock     sync.Mutex
)

func getData(w http.ResponseWriter, r *http.Request) {
	key := r.URL.Path[4:] // Trim the /db/ prefix

	dbLock.Lock()
	defer dbLock.Unlock()

	value, ok := keyValueDB[key]
	if !ok {
		http.Error(w, fmt.Sprintf("Key %q not found", key), http.StatusNotFound)
		return
	}

	entry := &Entry{
		Key:   key,
		Value: value,
	}

	sendResponse(entry, w)
}

func sendResponse(entry *Entry, w http.ResponseWriter) {
	w.Header().Set("Content-Type", "application/json")
	enc := json.NewEncoder(w)
	if err := enc.Encode(entry); err != nil {
		log.Printf("error encoding %+v - %s", entry, err)
	}
}

func assignData(w http.ResponseWriter, r *http.Request) {
	// Decode
	defer r.Body.Close()
	dec := json.NewDecoder(r.Body)
	entry := &Entry{}
	if err := dec.Decode(entry); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// Store data
	_, exist := keyValueDB[entry.Key]
	if exist {
		http.Error(w, "Key/value already existing in db", http.StatusBadRequest)
		return
	}

	keyValueDB[entry.Key] = entry.Value
	sendResponse(entry, w)
}

func startHTTPServer() {
	fmt.Println("Starting an http server listening on port 8081")
	http.HandleFunc("/hello", helloHandler)
	http.HandleFunc("/math", mathHandler)

	http.HandleFunc("/db", assignData)
	http.HandleFunc("/db/", getData)

	if err := http.ListenAndServe(":8081", nil); err != nil {
		log.Fatal(err)
	}
}
