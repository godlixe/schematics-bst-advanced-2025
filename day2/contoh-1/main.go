package main

import (
	"fmt"
	"net/http"
)

func jsonHelloHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	fmt.Fprint(w, `{"data":"hello"}`)
}

func textHelloHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Halo dari REST API Go!")
}

func main() {
	http.HandleFunc("/", textHelloHandler)
	http.ListenAndServe(":8080", nil)
}
