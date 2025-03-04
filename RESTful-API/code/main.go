package main

import (
	"RESTful-API/myapp"
	"net/http"
)

func main() {
	port := ":3000"

	http.ListenAndServe(port, myapp.NewHandler())
}
