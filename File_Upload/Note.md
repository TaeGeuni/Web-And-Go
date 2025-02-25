# HTTP 서버 기본 원리

## 기본 구조
```go
http.ListenAndServe(":3000", mux)
```
- **동작 방식**: TCP 소켓 생성 → 요청 대기 → 요청별로 고루틴 생성
- **멀티플렉서(mux)**: URL 경로와 핸들러 매핑 관리
- **고루틴**: 각 연결은 별도 고루틴에서 처리 (동시성 지원)

---

## Multipart Form 처리
```go
r.FormFile("upload_file")
```
- **MIME 타입**: `multipart/form-data` 형식으로 전송
- **파싱 과정**:
  1. 요청 본문을 파트 단위로 분할
  2. 각 파트의 헤더 분석 (파일명, 컨텐츠 타입 등)
  3. 임시 저장소에 파일 데이터 저장

---

## 파일 시스템 핸들링
```go
os.Create(filepath)
io.Copy(dst, src)
```
- **스트림 복사**: 메모리 효율적 처리 (대용량 파일 핸들링 가능)
- **퍼미션**:
  - `0777`: 모든 권한 허용 (개발용)
  - `0755`: 소유자 RWX, 그룹/기타 RX (실서비스 권장)
- **defer**: 파일 핸들 누수 방지

---

## 에러 처리 패턴
```go
if err != nil {
    w.WriteHeader(http.StatusInternalServerError)
    fmt.Fprint(w, err)
    return
}
```
- **HTTP 상태 코드**:
  - `400`: 클라이언트 오류 (잘못된 요청)
  - `500`: 서버 내부 오류
- **에러 노출**: 실제 서비스에선 상세 에러 클라이언트에 노출 금지

---

## 보안 고려사항 (실서비스 필수!)
- **파일 크기 제한**:
  ```go
  r.ParseMultipartForm(32 << 20) // 32MB
  ```
- **파일명 검증**:
  ```go
  cleanName := filepath.Base(header.Filename) // 악성 문자 필터링
  ```
- **MIME 타입 체크**:
  ```go
  buff := make([]byte, 512)
  if _, err = file.Read(buff); err != nil {...}
  detectedType := http.DetectContentType(buff)
  ```

---

## 성능 최적화
- **버퍼링**: `bufio.NewReader` 사용
- **병렬 처리**: 고루틴 + 채널 활용
- **메모리 관리**: `sync.Pool`로 버퍼 재사용

---

## Go 네트워크 스택
- **TCP 소켓**: `net.Listener` 기반
- **HTTP 프로토콜**: 요청/응답 라이프사이클
- **Keep-Alive**: 커넥션 재사용 메커니즘