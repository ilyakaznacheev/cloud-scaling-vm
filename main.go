package main

import (
	"flag"
	"fmt"
	"log"
	"net"
	"net/http"

	name "github.com/dustinkirkland/golang-petname"
)

const port = "80"

var nodeName string

func greet(w http.ResponseWriter, r *http.Request) {
	log.Printf("request from %s\n", r.RemoteAddr)
	fmt.Fprintf(w, "Hi, this is %s node", nodeName)
}

func main() {
	port := flag.String("p", "8080", "port")
	flag.Parse()

	name.NonDeterministicMode()
	nodeName = name.Generate(2, "-")

	log.Printf("starting app on port %s\n", *port)
	defer log.Println("exiting app")

	http.HandleFunc("/", greet)
	http.ListenAndServe(net.JoinHostPort("", *port), nil)
}
