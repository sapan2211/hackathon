package main

import (
	"fmt"
	"html"
	"log"
	"net/http"
)

func main() {
	str := "Hello World. This is  a simple respones from the server side"
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "%s , %q", str, html.EscapeString(r.URL.Path))
	})

	log.Fatal(http.ListenAndServe(":8080", nil))

}
