package main

import (
	"fmt"
	"log"
	"net/http"
)

type String string

type Hello string

type Struct struct {
	Greeting string
	Punct    string
	Who      string
}

func (h Hello) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, h)
}

func (s String) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, s)
}

func (ss *Struct) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, fmt.Sprintf("%v", ss))
}

func main() {
	http.Handle("/string", String("I'm a frayed knot."))
	http.Handle("/", Hello("Hello World"))
	http.Handle("/struct", &Struct{"Hello", ":", "Gophers!"})
	// http.Handle("/struct", &Struct{"Hello", ":", "Gophers!"})
	r := http.ListenAndServe("localhost:4000", nil)
	if r != nil {
		log.Fatal(r)
	}
}
