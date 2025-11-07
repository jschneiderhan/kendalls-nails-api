package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/jschneiderhan/kendalls-nails-api/api"
)

func main() {
	http.HandleFunc("/testing", handler.Handler)

	port := ":8080"
	fmt.Printf("Server starting on port %s\n", port)

	if err := http.ListenAndServe(port, nil); err != nil {
		log.Fatal(err)
	}
}
