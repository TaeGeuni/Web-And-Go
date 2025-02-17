package main

import (
	"fmt"
	"net/http"
)

type fooHandler struct{}

func (f *fooHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Hello Foo!")
}

func barHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Hello Bar!")
}

func main() {
	port := ":3000"

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) { //함수기반
		fmt.Fprintln(w, "Hello World")
	})

	http.HandleFunc("/bar", barHandler) //함수기반

	http.Handle("/foo", &fooHandler{}) //구조체 기반

	http.ListenAndServe(port, nil)
}
