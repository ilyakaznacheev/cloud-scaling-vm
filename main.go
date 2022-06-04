package main

import (
	"fmt"
	"net/http"

	name "github.com/dustinkirkland/golang-petname"
)

var nodeName string

func greet(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hi, this is %s node", nodeName)
}

func main() {
	name.NonDeterministicMode()
	nodeName = name.Generate(2, "-")

	http.HandleFunc("/", greet)
	http.ListenAndServe(":8080", nil)
}
