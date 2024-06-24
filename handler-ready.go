package main

import "net/http"


func handlerReady (w http.ResponseWriter, r *http.Request) {
	jsonResponse(w, 200, struct{}{})
}  