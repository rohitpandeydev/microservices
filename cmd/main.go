package main

import (
	"encoding/json"
	dlog "log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/rohitpandeydev/microservices/pkg/logger"
)

type Message struct {
	Message string `json:"message"`
}

var message = Message{Message: "Hello World"}

func home(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(message)
}

func main() {
	log := logger.New(logger.INFO)

	router := mux.NewRouter()
	router.HandleFunc("/helloworld", home).Methods("GET")
	log.Info("Starting library webservice on port 8080")
	dlog.Fatal(http.ListenAndServe(":8080", router))
}
