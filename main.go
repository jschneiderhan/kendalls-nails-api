package main

import (
	"fmt"
	"log"
	"net/http"
)

func testingHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "hello world")
}

func main() {
	http.HandleFunc("/testing", testingHandler)

	port := ":8080"
	fmt.Printf("Server starting on port %s\n", port)

	if err := http.ListenAndServe(port, nil); err != nil {
		log.Fatal(err)
	}
}
