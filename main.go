package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
)

func main() {
	var port = flag.Int("p", 7000, "port to listen on (default: 7000)")
	flag.Parse()
	http.Handle("/", http.StripPrefix("/", http.FileServer(http.Dir("."))))
	log.Fatal("ListenAndServe: ", http.ListenAndServe(fmt.Sprintf(":%d", *port), nil))
}
