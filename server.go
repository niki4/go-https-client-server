package main

import (
	"fmt"
	"log"
	"net/http"
	"strings"
	"time"
)

func TelemetryRouterHandler(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	log.Println("Received new request from", r.RemoteAddr)
	log.Println("path:\t", r.URL.Path)
	log.Println("headers:", r.Header)
	for k, v := range r.Form {
		log.Println("key:\t", k)
		log.Println("value:\t", strings.Join(v, ""))
	}
	fmt.Fprintf(w, "Hi client")
}

func main() {
	srv := &http.Server{
		Addr:         ":9001",
		ReadTimeout:  60 * time.Second,
		WriteTimeout: 60 * time.Second,
		IdleTimeout:  1 * time.Second,
	}

	http.HandleFunc("/telemetry", TelemetryRouterHandler)

	log.Println("Server started and waiting for connections.")
	err := http.ListenAndServeTLS(srv.Addr,
		"cert.pem",
		"key.pem",
		nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}

	defer srv.Close()
}
