package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

type User struct {
	FirstName string    `json:"first_name"`
	LastName  string    `json:"last_name"`
	Email     string    `json:"email"`
	CreatedAt time.Time `json:"created_at"`
}

type testHandler struct{}

func (t *testHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {

	user := new(User)

	err := json.NewDecoder(r.Body).Decode(user)

	if err != nil {
		fmt.Fprintln(w, "Bad Request:", err)
		return
	}
	user.CreatedAt = time.Now()
	data, _ := json.Marshal(user)

	w.Header().Add("content-type", "application:json")
	w.WriteHeader(http.StatusCreated)
	fmt.Fprintln(w, string(data))

}

func main() {
	port := ":3000"
	mux := http.NewServeMux()

	mux.Handle("/", &testHandler{})

	http.ListenAndServe(port, mux)
}
