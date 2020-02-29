package main

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

func main() {
	router := mux.NewRouter()
	router.HandleFunc("/hello/{name}", helloHandler)

	log.Fatal(http.ListenAndServe(":8080", router))
}

const ContentTypeApplicationJson = "application/json"

type HelloResponseDto struct {
	Greeting string `json:"greeting"`
}

func helloHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	name := vars["name"]
	w.Header().Add("Content-Type", ContentTypeApplicationJson)

	responseDto := HelloResponseDto{Greeting: "hello, " + name}
	encoder := json.NewEncoder(w)
	encoder.SetEscapeHTML(false)
	err := encoder.Encode(responseDto)
	if err != nil {
		log.Printf("error while encoding json response: %v", err)
		w.WriteHeader(http.StatusInternalServerError)
	}
}
