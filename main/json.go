package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type Data struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}

func handleJSON(w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)
	var data Data
	err := decoder.Decode(&data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	fmt.Fprintf(w, "Name: %s\\n", data.Name)
	fmt.Fprintf(w, "Age: %d\\n", data.Age)
}
