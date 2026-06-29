package main

import (
	"fmt"
	"net/http"
	"os"
)

func helloWorld(w http.ResponseWriter, r *http.Request) {
	greeting := os.Getenv("GREETING")
	fmt.Fprintf(w, "%s", greeting)
}

func main() {
	http.HandleFunc("/", helloWorld)
	http.ListenAndServe(":6112", nil)
}
