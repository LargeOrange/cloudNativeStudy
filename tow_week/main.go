package main

import (
	"io"
	"log"
	"net/http"
)

const OS_VERSION = "1.0.0"

func main() {
	http.HandleFunc("/healthz", healthz)
	//Use the default DefaultServeMux. • ListenAndService
	err := http.ListenAndServe(":80", nil)
	if err != nil {
		log.Fatal(err)
	}
}

func healthz(w http.ResponseWriter, r *http.Request) {
	customHeader := r.Header.Get("custom_header")
	w.Header().Set("custom_header", customHeader)
	w.Header().Set("os_version", OS_VERSION)
	io.WriteString(w, "ok")
}
