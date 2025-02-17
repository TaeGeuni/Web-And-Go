# Golang HTTP Mux & Data Transmission

## 1. HTTP Mux란?

`http.NewServeMux()`는 Go에서 제공하는 HTTP 요청을 라우팅하는 **멀티플렉서(Mux)** 입니다.
Mux를 사용하면 다양한 URL 패턴에 대한 핸들러를 등록하고 관리할 수 있습니다.

### ✅ Mux 사용 예시
```go
mux := http.NewServeMux()
mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintln(w, "Hello World")
})

mux.HandleFunc("/about", func(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintln(w, "About Page")
})

http.ListenAndServe(":3000", mux)
```

## 2. URL을 통한 데이터 전송 방법

### 🔹 1) **Query Parameters (쿼리 매개변수)**
URL에 `?key=value` 형식으로 데이터를 전달하는 방식입니다.

#### 📌 예제
```go
func handler(w http.ResponseWriter, r *http.Request) {
    name := r.URL.Query().Get("name")
    if name == "" {
        name = "Guest"
    }
    fmt.Fprintf(w, "Hello %s!", name)
}
```

#### 🌍 요청 예시
```sh
curl "http://localhost:3000/?name=Alice"
```

---

### 🔹 2) **Path Parameters (경로 매개변수)**
URL의 일부로 데이터를 포함하는 방식입니다.

#### 📌 예제 (고급 라우팅 필요)
```go
mux.HandleFunc("/users/{id}", func(w http.ResponseWriter, r *http.Request) {
    id := r.PathValue("id")
    fmt.Fprintf(w, "User ID: %s", id)
})
```

#### 🌍 요청 예시
```sh
curl "http://localhost:3000/users/123"
```

---

### 🔹 3) **POST Body를 통한 데이터 전송 (JSON 사용)**
POST 요청의 본문(Body)으로 데이터를 전달하는 방식입니다.

#### 📌 예제
```go
type User struct {
    Name  string `json:"name"`
    Email string `json:"email"`
}

func handler(w http.ResponseWriter, r *http.Request) {
    var user User
    json.NewDecoder(r.Body).Decode(&user)
    fmt.Fprintf(w, "Received: %s, %s", user.Name, user.Email)
}
```

#### 🌍 요청 예시
```sh
curl -X POST http://localhost:3000 -d '{"name":"Alice","email":"alice@example.com"}' -H "Content-Type: application/json"
```

---

## 3. JSON (JavaScript Object Notation) 이란?

JSON은 데이터를 저장하고 교환하기 위한 가볍고 읽기 쉬운 데이터 형식입니다.

### ✅ JSON의 특징
- **키-값 쌍(Key-Value Pair)** 형태
- **경량 & 언어 독립적**
- **사람이 읽기 쉽고, 기계가 처리하기 쉬움**

### 📌 JSON 예제
```json
{
  "name": "Alice",
  "age": 25,
  "email": "alice@example.com"
}
```

### 📌 Go에서 JSON 변환

#### **구조체 → JSON 변환**
```go
user := User{Name: "Alice", Email: "alice@example.com"}
jsonData, _ := json.Marshal(user)
fmt.Println(string(jsonData))
```

#### **JSON → 구조체 변환**
```go
var user User
json.Unmarshal([]byte('{"name":"Alice","email":"alice@example.com"}'), &user)
fmt.Println(user.Name, user.Email)
```

---

## 🎯 정리
✅ `http.NewServeMux()`를 사용하면 **라우팅을 효율적으로 관리**할 수 있음  
✅ **쿼리 매개변수, 경로 매개변수, POST Body(JSON)** 등으로 데이터를 전송할 수 있음  
✅ JSON은 **API 응답 및 데이터 교환에 필수적인 형식**으로 Go에서 `encoding/json` 패키지를 사용하여 쉽게 처리 가능

