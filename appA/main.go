package main

import (
	"fmt"
	"net/http"
	"os"
)

// This variable is injected at build time
var environment string = ""

// Serving hello URL
func hello(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("hello request received (%s %s %s)\n", r.RemoteAddr, r.Method, r.URL)
	fmt.Fprintf(w, `<h1 style="font-size:120px;">Hello %s!<h1/>`, environment)
}

func main() {
	listenAddress := ":8080"
	http.HandleFunc("/", hello)
	fmt.Printf("Listening on %s\n", listenAddress)
	if err := http.ListenAndServe(listenAddress, nil); err != nil {
		fmt.Printf("Server failed: %s\n", err)
		os.Exit(1)
	}
}
