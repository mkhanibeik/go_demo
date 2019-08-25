package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"os"
)

type Request struct {
	Login  string  `json:"user"`
	Type   string  `json:"type"`
	Amount float64 `json:"amount"`
}

var data = `
{
	"user": "Milad Khanibeik",
	"type": "deposit",
	"amount": 1235.3
}
`

func jsonTest() {
	reader := bytes.NewBufferString(data)
	decoder := json.NewDecoder(reader)

	req := &Request{}
	if err := decoder.Decode(req); err != nil {
		log.Fatalf("ERROR cannot decode - %s", err)
	}

	fmt.Printf("get: %+v\n", req)

	// create response
	prevBalance := 34345.2 // loaded from db
	resp := map[string]interface{}{
		"ok":      true,
		"balance": prevBalance + req.Amount,
	}

	encoder := json.NewEncoder(os.Stdout)
	if err := encoder.Encode(resp); err != nil {
		log.Fatalf("ERROR cannot encode -%s", err)
	}
}
