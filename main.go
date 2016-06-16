package main

import (
	"fmt"
	"github.com/nbari/violetear"
	"log"
	"net/http"
)

var version, githash string

func catchAll(w http.ResponseWriter, r *http.Request) {
	html := `
	<html>
	<head>
		<meta http-equiv="refresh" content="3" />
	</head>
	<body>
	<h3>Hallo</h3>
	<small>%s</small>`
	fmt.Fprintf(w, html, version)
}

func main() {
	router := violetear.New()
	router.LogRequests = true

	router.HandleFunc("*", catchAll)

	log.Fatal(http.ListenAndServe(":8000", router))
}
