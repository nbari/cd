package main

import (
	"fmt"
	"github.com/nbari/violetear"
	"log"
	"net/http"
)

var version, githash string

func catchAll(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello, %s+%s", version, githash)
}

func main() {
	router := violetear.New()
	router.LogRequests = true

	router.HandleFunc("*", catchAll)

	log.Fatal(http.ListenAndServe(":8080", router))
}
