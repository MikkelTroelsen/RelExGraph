package main

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"
)

type Relation struct {
	subject string
	label   string
	object  string
}

func callRelik(text string) []Relation {
	payload := map[string]interface{}{
		"text": text,
	}
	jsonData, err := json.Marshal(payload)
	if err != nil {
		panic(err)
	}

	resp, err := http.Post(
		"http://127.0.0.1:5000/get-relations",
		"application/json",
		bytes.NewBuffer(jsonData),
	)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}

	// Parse the response: [["subject", "label", "object"], ...]
	var raw [][]string
	if err := json.Unmarshal(body, &raw); err != nil {
		panic(err)
	}
	relations := make([]Relation, len(raw))
	for i, r := range raw {
		if len(r) != 3 {
			panic("invalid relation format")
		}
		relations[i] = Relation{
			subject: r[0],
			label:   r[1],
			object:  r[2],
		}
	}
	return relations
}
