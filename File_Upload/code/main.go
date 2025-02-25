package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

func uploadsHandler(w http.ResponseWriter, r *http.Request) {
	uploadFile, header, err := r.FormFile("upload_file")

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprint(w, err)
		return
	}
	defer uploadFile.Close()

	dirname := "./uploads"
	os.MkdirAll(dirname, 0777)
	filepath := fmt.Sprintf("%s/%s", dirname, header.Filename)
	file, err := os.Create(filepath)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprint(w, err)
		return
	}

	defer file.Close()

	io.Copy(file, uploadFile)
	w.WriteHeader(http.StatusOK)
	fmt.Fprint(w, filepath)
}

func main() {
	port := ":3000"
	mux := http.NewServeMux()

	mux.HandleFunc("/uploads", uploadsHandler)
	mux.Handle("/", http.FileServer(http.Dir("public")))

	http.ListenAndServe(port, mux)
}
