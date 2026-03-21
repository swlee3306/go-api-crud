# Go-CRUD

## 한 줄 소개
Go로 API와 DB CRUD 구조를 연습한 백엔드 실습 저장소입니다.

## 저장소 성격
- 분류: 백엔드 기본기
- 목적: API 설계, DB 연동, 인증과 검증 구조 연습
- 핵심 기술: Go, REST API, DB, Docker

기본적인 CRUD 기능을 넘어서 백엔드 구조를 연습하기 위한 저장소입니다.

Go 언어를 사용하여 데이터베이스의 기본적인 CRUD(Create, Read, Update, Delete) 작업을 수행하는 API 서버입니다. API 설계, 인증/인가, 검증, 로깅, 테스트 구조를 함께 연습하기 위해 정리했습니다.

## 🚀 주요 기능

### 🔐 인증 및 보안
- **JWT 기반 인증**: 안전한 토큰 기반 인증 시스템
- **역할 기반 접근 제어 (RBAC)**: Admin, User, Guest 역할 지원
- **패스워드 해싱**: bcrypt를 사용한 안전한 패스워드 저장
- **미들웨어 인증**: 요청별 인증 및 권한 검사

### 📊 데이터베이스 지원
- **다중 데이터베이스**: MySQL, PostgreSQL, SQLite 지원
- **GORM ORM**: 강력한 ORM을 통한 데이터베이스 추상화
- **자동 마이그레이션**: 스키마 자동 생성 및 업데이트
- **연결 풀링**: 효율적인 데이터베이스 연결 관리

### 🛡️ 데이터 검증
- **입력 검증**: 구조화된 데이터 검증 시스템
- **에러 처리**: 표준화된 에러 응답
- **타입 안전성**: Go의 타입 시스템을 활용한 안전한 데이터 처리

### 📝 로깅 및 모니터링
- **구조화된 로깅**: JSON 형식의 구조화된 로그
- **요청 추적**: 각 요청의 상세한 로깅
- **에러 로깅**: 상세한 에러 정보 및 스택 트레이스

### 🧪 테스트 및 품질
- **단위 테스트**: 포괄적인 단위 테스트 커버리지
- **통합 테스트**: 데이터베이스 통합 테스트
- **벤치마크 테스트**: 성능 측정 및 최적화

### 🐳 컨테이너화 및 배포
- **Docker 지원**: Dockerfile 및 docker-compose.yml 제공
- **환경별 설정**: YAML 및 환경 변수 기반 설정
- **헬스체크**: 시스템 상태 모니터링 및 진단
- **API 문서화**: Swagger 기반 자동 문서 생성

### 🔧 개발자 도구
- **속도 제한**: Rate limiting 미들웨어
- **페이지네이션**: 효율적인 데이터 페이징
- **내부 시스템**: 데이터베이스 링커 및 시스템 환경 관리
- **유틸리티**: 공통 기능 및 헬퍼 함수

## 📦 설치 및 실행

### 1. 저장소 클론
```bash
git clone https://github.com/swlee3306/go_crud.git
cd go_crud
```

### 2. 의존성 설치
```bash
go mod tidy
```

### 3. 환경 변수 설정
```bash
# .env 파일 생성
cp .env.example .env

# 데이터베이스 설정
export DB_CONNECTION=mysql
export DB_HOST=127.0.0.1
export DB_PORT=3306
export DB_DATABASE=go_crud_db
export DB_USERNAME=root
export DB_PASSWORD=password
```

### 4. 데이터베이스 설정
```bash
# MySQL 예시
mysql -u root -p
CREATE DATABASE go_crud_db;
```

### 5. 서버 실행
```bash
# 개발 모드
go run main_new.go

# 프로덕션 모드
go build -o go_crud main_new.go
./go_crud
```

## 🏗️ 프로젝트 구조

```
go_crud/
├── main.go                 # 원본 메인 파일 (레거시)
├── main_new.go            # 새로운 메인 파일 (포트폴리오용)
├── main_LoadEnv.go        # 환경 변수 로드
├── main_LoadYml.go        # YAML 설정 로드
├── config/                # 설정 관리
│   ├── database.go        # 데이터베이스 설정
│   ├── drivers.go         # 데이터베이스 드라이버
│   ├── connection.go      # 연결 관리
│   └── test.go           # 연결 테스트
├── models/               # 데이터 모델
│   ├── user.go           # 사용자 모델
│   ├── post.go           # 게시글 모델
│   ├── tag.go            # 태그 모델
│   └── migrate.go        # 마이그레이션
├── handlers/             # HTTP 핸들러
│   └── user.go           # 사용자 핸들러
├── routes/               # 라우팅
│   └── routes.go         # 라우트 설정
├── auth/                 # 인증 시스템
│   ├── jwt.go            # JWT 토큰 관리
│   └── rbac.go           # 역할 기반 접근 제어
├── middleware/           # 미들웨어
│   ├── auth.go           # 인증 미들웨어
│   └── ratelimit.go      # 속도 제한 미들웨어
├── validation/           # 데이터 검증
│   ├── validator.go      # 검증기
│   └── user_validation.go # 사용자 검증
├── logging/              # 로깅 시스템
│   └── logger.go         # 로거 설정
├── health/               # 헬스체크
│   └── health.go         # 헬스체크 시스템
├── docs/                 # API 문서
│   └── swagger.go        # Swagger 문서
├── internal/             # 내부 패키지
│   ├── dblinker/         # 데이터베이스 링커
│   ├── sysdef/           # 시스템 정의
│   └── sysenv/           # 시스템 환경
├── utils/                # 유틸리티
│   ├── pagination.go     # 페이지네이션
│   └── router/           # 라우터 유틸리티
├── docker-compose.yml    # Docker Compose 설정
├── Dockerfile            # Docker 이미지 설정
├── setting.yml           # YAML 설정 파일
├── .env.example          # 환경 변수 예제
├── .env                  # 환경 변수
├── go.mod               # Go 모듈 파일
├── go.sum               # 의존성 체크섬
└── README.md            # 프로젝트 문서
```

## 🔧 API 엔드포인트

### 인증 엔드포인트
```http
POST /auth/login
Content-Type: application/json

{
  "username": "testuser",
  "password": "password123"
}
```

### 사용자 관리
```http
# 사용자 생성
POST /users
Authorization: Bearer <token>
Content-Type: application/json

{
  "username": "newuser",
  "email": "user@example.com",
  "password": "password123"
}

# 사용자 목록 조회
GET /users
Authorization: Bearer <token>

# 특정 사용자 조회
GET /users/{id}
Authorization: Bearer <token>

# 사용자 정보 수정
PUT /users/{id}
Authorization: Bearer <token>
Content-Type: application/json

{
  "username": "updateduser",
  "email": "updated@example.com"
}

# 사용자 삭제
DELETE /users/{id}
Authorization: Bearer <token>
```

## 🛠️ 데이터베이스 모델

### User 모델
```go
type User struct {
    gorm.Model
    Username    string `gorm:"unique;not null"`
    Email       string `gorm:"unique;not null"`
    Password    string `gorm:"not null"`
    UserProfile UserProfile
    Posts       []Post
    Comments    []Comment
}
```

### Post 모델
```go
type Post struct {
    gorm.Model
    Title      string `gorm:"not null"`
    Content    string `gorm:"type:text"`
    UserID     uint   `gorm:"not null"`
    User       User
    CategoryID uint `gorm:"not null"`
    Category   Category
    Comments   []Comment
    Tags       []Tag `gorm:"many2many:post_tags;"`
}
```

## 🔐 보안 기능

### JWT 토큰 관리
```go
// 토큰 생성
token, err := jwtManager.GenerateToken(user)

// 토큰 검증
claims, err := jwtManager.VerifyToken(token)
```

### 역할 기반 접근 제어
```go
// 역할 확인
if claims.HasRole(auth.RoleAdmin) {
    // 관리자 권한 작업
}
```

### 패스워드 보안
```go
// 패스워드 해싱
hashedPassword, err := auth.HashPassword(password)

// 패스워드 검증
isValid := auth.CheckPasswordHash(password, hashedPassword)
```

### 헬스체크 시스템
```go
// 헬스체크 엔드포인트
GET /health

// 응답 예시
{
  "status": "healthy",
  "timestamp": "2023-01-01T00:00:00Z",
  "version": "1.0.0",
  "uptime": "2h30m15s",
  "checks": {
    "database": {
      "status": "healthy",
      "message": "Database connection is healthy"
    },
    "memory": {
      "status": "healthy",
      "message": "Memory usage is normal"
    }
  }
}
```

### Docker 사용법
```bash
# Docker 이미지 빌드
docker build -t go-crud .

# Docker Compose로 실행
docker-compose up -d

# 로그 확인
docker-compose logs -f

# 서비스 중지
docker-compose down
```

### Rate Limiting
```go
// Rate limiting 미들웨어 설정
r.Use(middleware.RateLimitMiddleware(100, time.Minute))

// IP별 요청 제한
r.Use(middleware.IPRateLimitMiddleware(10, time.Minute))
```

## 🧪 테스트

### 단위 테스트 실행
```bash
# 모든 테스트 실행
go test ./...

# 특정 패키지 테스트
go test ./models
go test ./handlers
go test ./auth

# 테스트 커버리지 확인
go test -cover ./...
```

### 통합 테스트
```bash
# 데이터베이스 통합 테스트
go test ./config -v

# 인증 시스템 테스트
go test ./auth -v
```

## 📊 성능 및 모니터링

### 데이터베이스 연결 풀 설정
```go
// 연결 풀 설정
sqlDB.SetMaxIdleConns(10)
sqlDB.SetMaxOpenConns(100)
sqlDB.SetConnMaxLifetime(time.Hour)
```

### 로깅 설정
```go
// 구조화된 로깅
logger.WithFields(logrus.Fields{
    "user_id": userID,
    "action":  "create_user",
}).Info("User created successfully")
```

## 🚀 배포 및 운영

### Docker 사용
```dockerfile
FROM golang:1.21-alpine AS builder
WORKDIR /app
COPY . .
RUN go mod download
RUN go build -o go_crud main_new.go

FROM alpine:latest
RUN apk --no-cache add ca-certificates
WORKDIR /root/
COPY --from=builder /app/go_crud .
COPY --from=builder /app/.env .
CMD ["./go_crud"]
```

### Docker Compose
```yaml
version: '3.8'
services:
  go_crud:
    build: .
    ports:
      - "8080:8080"
    environment:
      - DB_CONNECTION=mysql
      - DB_HOST=mysql
      - DB_PORT=3306
      - DB_DATABASE=go_crud_db
      - DB_USERNAME=root
      - DB_PASSWORD=REPLACE_WITH_PRIVATE_VALUE
    depends_on:
      - mysql
  
  mysql:
    image: mysql:8.0
    environment:
      MYSQL_ROOT_PASSWORD: REPLACE_WITH_PRIVATE_VALUE
      MYSQL_DATABASE: go_crud_db
    ports:
      - "3306:3306"
```

## 🔧 환경 변수

### 필수 환경 변수
```bash
# 데이터베이스 설정
DB_CONNECTION=mysql          # mysql, postgres, sqlite
DB_HOST=127.0.0.1
DB_PORT=3306
DB_DATABASE=go_crud_db
DB_USERNAME=root
DB_PASSWORD=REPLACE_WITH_PRIVATE_VALUE

# JWT 설정
JWT_SECRET=your-secret-key
JWT_EXPIRE_HOURS=24
```

### 선택적 환경 변수
```bash
# 서버 설정
SERVER_PORT=8080
SERVER_HOST=0.0.0.0

# 로깅 설정
LOG_LEVEL=info
LOG_FORMAT=json
```

## 🤝 기여하기

1. Fork the repository
2. Create your feature branch (`git checkout -b feature/AmazingFeature`)
3. Commit your changes (`git commit -m 'Add some AmazingFeature'`)
4. Push to the branch (`git push origin feature/AmazingFeature`)
5. Open a Pull Request

## 📝 라이센스

이 프로젝트는 MIT 라이센스 하에 배포됩니다. 자세한 내용은 [LICENSE](LICENSE) 파일을 참조하세요.

## 🙏 감사의 말

- [Go](https://golang.org/) - 프로그래밍 언어
- [Gin](https://gin-gonic.com/) - HTTP 웹 프레임워크
- [GORM](https://gorm.io/) - ORM 라이브러리
- [JWT-Go](https://github.com/golang-jwt/jwt) - JWT 토큰 라이브러리

## 📞 지원 및 문의

- 이슈 리포트: [GitHub Issues](https://github.com/swlee3306/go_crud/issues)
- 이메일: swlee3306@gmail.com
- 문서: [Wiki](https://github.com/swlee3306/go_crud/wiki)

---

**Go-CRUD** - 프로덕션 레디 RESTful API 서버 🚀
