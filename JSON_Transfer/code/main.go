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

type fooHandler struct{}

func (f *fooHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	user := new(User)                           // 빈 User 구조체 생성
	err := json.NewDecoder(r.Body).Decode(user) // 요청 본문(JSON) → 구조체 변환

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintln(w, "Bad Request:", err)
		return
	}

	user.CreatedAt = time.Now() // 현재 시간 설정

	data, _ := json.Marshal(user) // 구조체 → JSON 변환
	w.Header().Add("content-type", "application/json")
	w.WriteHeader(http.StatusCreated) // 201 Created 응답 코드
	fmt.Fprintln(w, string(data))     // JSON 응답 전송
}

func barHandler(w http.ResponseWriter, r *http.Request) {
	name := r.URL.Query().Get("name") // URL 쿼리 파라미터 가져오기

	if name == "" {
		name = "Bar" // 기본값 설정
	}

	fmt.Fprintf(w, "Hello %s!", name) // 응답 전송
}

func main() {
	port := ":3000"
	mux := http.NewServeMux()

	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "Hello World")
	})

	mux.HandleFunc("/bar", barHandler) // 함수 기반 핸들러
	mux.Handle("/foo", &fooHandler{})  // 구조체 기반 핸들러

	http.ListenAndServe(port, mux) // 서버 실행
}
