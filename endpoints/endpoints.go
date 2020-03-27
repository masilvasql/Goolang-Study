package main

import (
	"encoding/json"
	"log"
	"net/http"
)

type Profile struct {
	Name    string
	Hobbies []string
}

func main() {
	http.HandleFunc("/first", hello)
	log.Print("EXEC")
	log.Fatal(http.ListenAndServe(":3001", nil))

}

func hello(w http.ResponseWriter, r *http.Request) {
	profile := Profile{"Alex", []string{"snowboarding", "programming"}}
	js, err := json.Marshal(profile)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(js)
}
