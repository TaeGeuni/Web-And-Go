package main

import (
	"net/http"
	"test1/myapp"
)

func main() {
	port := ":3000"

	http.ListenAndServe(port, myapp.NewHttpHandler()) // 서버 실행
}
