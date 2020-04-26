package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	addr := ":8080"
	http.HandleFunc("/", handler)
	log.Printf("Server started on port %s", addr)
	log.Fatal(http.ListenAndServe(addr, nil))
}

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hi there, I love %s!", r.URL.Path[1:])
	GetAlbum(r.URL.Path[1:])
}
