package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	// "github.com/rohitpandeydev/microservice/pkg/logger"
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
	//  log := logger
	router := mux.NewRouter()
	router.HandleFunc("/", home).Methods("GET")
	fmt.Println("Server running on port 8080")
	http.ListenAndServe(":8080", router)
}
