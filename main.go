package main

import (
	"encoding/json"
	"log"
	"net/http"
)

type Data struct {
	Foo string `json:"foo"`
	Bar int    `json:"bar"`
}

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		var data Data
		if err := json.NewDecoder(r.Body).Decode(&data); err != nil {
			log.Println(err)
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		log.Printf("%+v", data)

		if err := json.NewEncoder(w).Encode(data); err != nil {
			log.Println(err)
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
	})
	log.Println("start")
	http.ListenAndServe(":8080", nil)
}
