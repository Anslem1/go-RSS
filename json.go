package main

import (
	"encoding/json"
	"log"
	"net/http"
)

func respondWithError(w http.ResponseWriter, code int, msg string) {
	if code > 499 {
		log.Printf("Responding with %v error: %v", code, msg)
	}
	type errorResponse struct {
		Error string `json:"error"`
	}
	jsonResponse(w, code, errorResponse{
		Error: msg,
	})
}

func jsonResponse(w http.ResponseWriter, code int, payload interface{}) {
	data, err := json.Marshal(payload)

	if err != nil {
		log.Fatal(err)
		log.Println("failed to marshal", payload)
		w.WriteHeader(500)
		return
	}
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(data)
}
