# go-api-crud

## 한 줄 소개
JWT 인증, 사용자 CRUD, health check를 포함한 Go API 서버 예제입니다.

## 저장소 성격
- 분류: 백엔드 / API 서버
- 목적: 인증, 검증, DB 연동, 사용자 CRUD 구조 정리
- 핵심 기술: Go, Gin, GORM, JWT, MySQL/PostgreSQL/SQLite

## 현재 구현 범위
- 사용자 회원가입과 로그인
- JWT 기반 인증 미들웨어
- 사용자 목록, 조회, 수정, 삭제
- DB 자동 마이그레이션
- health, live, ready 엔드포인트
- IP 기준 rate limit

`Post`, `Comment`, `Tag`, `RBAC` 모델과 코드도 포함되어 있지만, 현재 기본 실행 경로에서 공개 API로 연결된 범위는 사용자 인증과 사용자 CRUD입니다.

## 실행 방법

### 1. 저장소 클론
```bash
git clone https://github.com/swlee3306/go-api-crud.git
cd go-api-crud
```

### 2. 의존성 설치
```bash
go mod tidy
```

### 3. 환경 변수 설정
가장 간단한 실행 방식은 SQLite입니다.

```bash
export DB_DRIVER=sqlite
export DB_NAME=go_crud_db
export JWT_SECRET=change-me
export SERVER_PORT=8080
```

MySQL 예시는 아래와 같습니다.

```bash
export DB_DRIVER=mysql
export DB_HOST=localhost
export DB_PORT=3306
export DB_USER=root
export DB_PASSWORD=password
export DB_NAME=go_crud_db
export JWT_SECRET=change-me
```

`.env.example` 파일도 함께 제공됩니다.

### 4. 실행
```bash
go run .
```

### 5. 빌드
```bash
go build -o go-api-crud .
./go-api-crud
```

## 주요 엔드포인트

### 인증
```http
POST /auth/register
POST /auth/login
```

로그인은 `email` 또는 `username` 중 하나와 `password`로 요청할 수 있습니다.

예시:
```http
POST /auth/register
Content-Type: application/json

{
  "username": "testuser",
  "email": "test@example.com",
  "password": "Password123!",
  "first_name": "Test",
  "last_name": "User"
}
```

```http
POST /auth/login
Content-Type: application/json

{
  "email": "test@example.com",
  "password": "Password123!"
}
```

### 사용자 CRUD
아래 엔드포인트는 `Authorization: Bearer <token>` 헤더가 필요합니다.

```http
GET    /users
POST   /users
GET    /users/{id}
PUT    /users/{id}
DELETE /users/{id}
```

같은 핸들러가 `/api/v1/users` 아래에도 연결되어 있습니다.

### 상태 확인
```http
GET /health
GET /health/live
GET /health/ready
```

## 프로젝트 구조
```text
go-api-crud/
├── main.go
├── config/
├── auth/
├── middleware/
├── handlers/
├── routes/
├── models/
├── validation/
├── health/
├── docs/
├── logging/
├── utils/
├── .env.example
└── README.md
```

## 레거시 코드
이 저장소에는 예전 `datastore` 실험 코드가 남아 있습니다.
기본 빌드에서는 `legacy` 빌드 태그로 분리해 제외했습니다.

```bash
go build -tags legacy ./...
```

## 검증
```bash
go test ./...
```

## 핀 저장소용 설명 문구
JWT 인증과 사용자 CRUD, health check를 포함한 Go API 서버 예제
