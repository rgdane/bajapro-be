# BAJAPRO Backend

Backend aplikasi E-Learning Java **BAJAPRO**. Go + Fiber, GORM, PostgreSQL.

## Project Structures

```bash
.
├── src/
│   ├── api/                                 # Layer presentasi (HTTP/API)
│   │   ├── controllers/v1/                  # Controller per resource
│   │   │   ├── dto/                         # DTO (Create/Update/Filter/Response)
│   │   │   ├── handlers/                    # Logic handler (transaksi)
│   │   │   ├── mapper/                      # Konversi DTO <-> model
│   │   │   └── m_course_controller.go       # Contoh controller master
│   │   ├── middleware/                      # JWT, RBAC (RequireRole)
│   │   ├── presenters/                      # SuccessResponse / ErrorResponse
│   │   └── routes/v1/                       # routes.go + <resource>_routes.go
│   ├── cmd/                                 # Entry point (main.go + bootstrap/)
│   ├── internal/
│   │   ├── config/                          # fiber, logger, postgres, neo4j
│   │   ├── container/                       # DI wiring (app_container.go)
│   │   ├── database/
│   │   │   ├── models/                      # Model GORM
│   │   │   ├── migrations/                  # AutoMigrate
│   │   │   └── seeders/                     # Data awal
│   │   └── constant/
│   └── pkg/
│       ├── repository/
│       │   ├── adapter/sql/                 # Interface repository
│       │   └── query/sql/                   # Implementasi query
│       └── services/v1/                     # Business logic
├── docker-compose.yml                       # Stack produksi (db, backend, frontend, piston)
├── .env.example                             # Variabel untuk docker compose
├── Dockerfile                               # Image backend (multi-stage)
└── makefile                                 # Perintah Docker (build/run/logs)
```

Alur request: `routes → controller → handler → service → repository → DB`.

## Kebutuhan

- Go 1.24+
- PostgreSQL 16
- Docker + Docker Compose (untuk produksi & sandbox)

## Development

1. Clone repo

   ```bash
   git clone <repo>
   cd bajapro-be
   ```

2. Siapkan environment

   ```bash
   cp src/.env.example src/.env
   ```

   Sesuaikan `DB_HOST`, `DB_PORT`, `DB_USERNAME`, `DB_PASSWORD`, `DB_DATABASE`,
   `JWT_SECRET`. Lihat daftar variabel di `src/.env.example`.

3. Jalankan

   ```bash
   cd src
   go mod tidy
   go run cmd/main.go
   ```

   Server berjalan di <http://localhost:5000> (default `PORT=8080` bila tidak diisi).

   Migrasi dan seeder dijalankan otomatis saat startup
   (`internal/database/migrations`, `internal/database/seeders`).

4. Cek kompilasi

   ```bash
   cd src
   go build ./...
   ```

## Produksi (Docker Compose)

Struktur clone di server — dua repo bersebelahan:

```
<dir>/bajapro-be/
<dir>/bajapro-fe-revamp/
```

Jalankan dari dalam `bajapro-be`:

```bash
cp .env.example .env      # isi POSTGRES_* dan JWT_SECRET
docker compose up -d --build
```

Service yang berjalan:

| Service    | Port | Keterangan |
|------------|------|------------|
| `db`       | -    | PostgreSQL 16 (volume `db_data`) |
| `backend`  | 8080 | API Go Fiber |
| `frontend` | 3000 | Next.js (repo `bajapro-fe-revamp`) |
| `piston`   | 2000 | Sandbox eksekusi kode (Java) |

Aktivasi runtime Java di Piston (sekali saja):

```bash
docker compose exec piston cli install java
```

Log: `docker compose logs -f backend`.

> **Catatan:** Piston membutuhkan **cgroup v2** (default di Linux modern). Di
> Docker Desktop Windows dengan cgroup v1, service `piston` tidak bisa start.

### Tanpa Piston

Sandbox tidak dibutuhkan untuk fitur admin. Jalankan hanya service inti:

```bash
docker compose up -d db backend frontend
```

## Environment

| Variabel | Keterangan |
|----------|------------|
| `PORT` | Port HTTP server |
| `DB_CONNECTION` | Driver DSN (`postgres`) |
| `DB_HOST`, `DB_PORT` | Alamat PostgreSQL |
| `DB_USERNAME`, `DB_PASSWORD`, `DB_DATABASE` | Kredensial database |
| `DB_SSLMODE` | Mode SSL (default `disable`) |
| `JWT_SECRET` | Secret penandatangan JWT |
| `PISTON_URL` | Base URL Piston untuk sandbox (default `http://localhost:2000`) |
| `NEO4J_*`, `MONGODB_*`, `ELASTIC_*`, `GCP_BUCKET_NAME` | Disiapkan untuk kebutuhan mendatang; **belum dipakai** saat ini |

## Sandbox eksekusi kode

`POST /api/v1/code/run` (butuh JWT) meneruskan kode Java ke Piston dan
mengembalikan `{ stdout, stderr, code }`. Endpoint ini dipakai halaman
`/sandbox` di frontend.

## Role & akses

Role user bersifat statis: `admin`, `teacher`, `student` (tidak ada tabel role).
Route terlindungi `middleware.JWTMiddleware()`; operasi tulis admin memakai
`middleware.RequireRole("admin")`.
