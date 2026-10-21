package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "Hello from the backend!")
	})

	log.Println("Backend running on :9000")
	log.Fatal(http.ListenAndServe(":9000", nil))
}
