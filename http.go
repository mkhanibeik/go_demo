package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
)

func httpTest() {
	// get request
	resp, err := http.Get("http://httpbin.org/get")
	if err != nil {
		log.Fatalf("ERROR cannot call httpbin.org")
	}
	defer resp.Body.Close()
	io.Copy(os.Stdout, resp.Body)

	fmt.Println("--------")
	// post request
	job := &Job{
		User:   "Milad Khanibeik",
		Action: "punch",
		Count:  10,
	}
	var buf bytes.Buffer
	encoder := json.NewEncoder(&buf)
	if err := encoder.Encode(job); err != nil {
		log.Fatalf("ERROR cannot encode job - %s", err)
	}

	resp, err = http.Post("https://httpbin.org/post", "application/json", &buf)
	if err != nil {
		log.Fatalf("ERROR cannot call httpbin.org")
	}
	defer resp.Body.Close()
	io.Copy(os.Stdout, resp.Body)
}

type Job struct {
	User   string `json:"user"`
	Action string `json:"action"`
	Count  int    `json:"count"`
}
