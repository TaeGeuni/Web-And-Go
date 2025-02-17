# Golang HTTP 핸들러: 함수 기반 vs 구조체 기반

## 1. 함수 기반 (`http.HandleFunc`)

```go
http.HandleFunc("/bar", func(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Hello Bar!")
})
```

### ✅ 장점
- 간결하고 빠르게 구현 가능
- 간단한 라우팅에 적합
- 유지보수가 쉬움

### ❌ 단점
- 상태(예: DB 연결, 설정 값)를 유지하기 어려움
- 여러 개의 핸들러가 있을 때 코드가 복잡해질 수 있음

### 📌 언제 사용?
- 간단한 API 서버나 정적 페이지 제공할 때
- 상태를 가질 필요가 없는 단순 요청 처리

---

## 2. 구조체 기반 (`http.Handler` 인터페이스 구현)

```go
type fooHandler struct{}

func (f *fooHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Hello Foo!")
}

http.Handle("/foo", &fooHandler{})
```

### ✅ 장점
- 핸들러가 상태(예: DB 연결, 설정 값, 인증 정보)를 가질 수 있음
- 재사용성과 확장성이 높음
- 중간 미들웨어 적용이 쉬움

### ❌ 단점
- 초기 설정이 다소 복잡
- 간단한 API 서버에는 불필요하게 복잡할 수 있음

### 📌 언제 사용?
- 상태를 유지해야 하는 핸들러(예: DB 연결, 사용자 인증 등)
- 복잡한 API 서버에서 여러 개의 엔드포인트를 관리할 때
- 중간에 미들웨어 적용이 필요할 때

---

## 🎯 결론: 일반적으로 어떤 것을 사용할까?
✅ **간단한 API** → **함수 기반 (`http.HandleFunc`)**
✅ **상태가 필요한 API** (DB 연결, 인증) → **구조체 기반 (`http.Handler`)**
✅ **대규모 프로젝트** → **구조체 기반 + 미들웨어 활용**

실제로 초기 개발에서는 `http.HandleFunc()`을 많이 사용하고, 프로젝트가 커질수록 구조체 기반 또는 `mux` 라우터(Gorilla Mux, chi, gin)를 활용하는 방식으로 발전합니다. 🚀

