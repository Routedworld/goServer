package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {

	log.Println("starting: Building Apps For K8s app")

	http.HandleFunc("/", handler)
	http.ListenAndServe(":8000", nil)

}

func handler(w http.ResponseWriter, r *http.Request) {

	log.Println("Request received from %s", r.RemoteAddr)
	fmt.Fprintf(w, "Building Apps for K8s app says Hi")

}
