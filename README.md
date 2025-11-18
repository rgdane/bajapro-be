# JalanKerja Backend

## Project Structures

```bash
.
└── src/
    ├── api/                                 # Layer presentasi (HTTP/API layer)
    │   ├── controllers/
    │   │   └── v1/                          # Versi API v1
    │   │       ├── dto/                     # DTO (Data Transfer Object)
    │   │       │   └── squad_dto.go         # Struct DTO untuk operasi Squad
    │   │       ├── handler/                 # Handler yang menangani logic
    │   │       │   └── squad_handler.go     # Handler fungsi CRUD untuk Squad
    │   │       ├── mapper                   # (Opsional) Mapping antara model, DTO
    │   │       └── squad_controller.go      # Menghubungkan route ke handler Squad
    │   ├── presenters/                      # Output formatter
    │   ├── middleware/                      # Middleware seperti Auth, Logging, dll.
    │   └── routes/                          # Routing layer
    │       └── v1/                          # Versi API v1
    │          ├── routes.go                 # Entry point untuk semua route
    │          └── squad_routes.go           # Daftar route terkait Squad
    ├── cmd/                                 # Entry point aplikasi
    │   └── main.go                          # Fungsi utama untuk menjalankan server
    ├── internal/
    │   └── database/
    │       ├── config/                      # Konfigurasi
    │       │   └── postgre.go               # Setup PostgreSQL (DSN, koneksi, migrasi)
    │       ├── models/                      # Model ORM (representasi tabel database)
    │       │   └── squads.go                # Struct model Squad
    │       └── seeder/                      # Seeder data awal ke database
    └── pkg/
        ├── repository/                      # Layer repository (akses data ke DB)
        │   ├── adapter/
        │   │   └── sql/                     # Implementasi repository berbasis SQL
        │   │       └── squad.go             # Repository interface Squad
        │   └── query/
        │       └── sql/
        │           └── squad.go             # Implement repo Squad
        └── services/                        # Logika bisnis
            └── v1/                          # Versi API v1
                └── squad_service.go         # Business logic squad
```

# Flows

![Flow Diagram](https://jam.dev/cdn-cgi/image/width=1000,quality=100,dpr=1/https://cdn-jam-screenshots.jam.dev/4bee731580457b0e55664da35511cc00/screenshot/da869424-f1e0-4435-a921-c790171b9c9d.png)

# Installations

1. Clone Repo

   ```bash
   git clone <repo>
   cd <repo>
   ```

2. Copy environtment

   ```bash
   cp src/.env.example src/.env
   ```

3. Run App

   - Development

   ```javascript
   cd src
   go mod tidy
   go run cmd/main.go
   ```

   - Production

   ```bash
   make build
   make run

   # for log
   make logs
   ```

App run on <http://localhost:5000>

1. Masuk `src` directory
2. Pastikan `air` sudah terinstall

   ```console
   go install github.com/cosmtrek/air@latest
   ```

   📌 Perintah ini akan mengunduh dan meng-install Air ke direktori $GOPATH/bin atau $HOME/go/bin.
   Sehingga ketika ingin membuat project **Go** baru kita tidak perlu mengnstall lagi

3. Jika sudah terinstall jalankan

   ```console
   air init
   ```

   Ini akan menghasilkan file konfigurasi `.air.toml` di dalam `src`

4. Ubah konfigurasi di file `.air.toml`
   Ubah konfigurasi `cmd` di file `.air.toml` sesuai dengan direktori main.go

   ```console
   cmd = "go build -o ./tmp/main <direktori file main.go>"
   // dalam kasus project ini uban menjadi ./cmd/main.go

   cmd = "go build -o ./tmp/main ./cmd/main.go"
   ```

5. Jalankan `air` di console untuk menjalankan `main.go` dengan auto reload
   ```console
   air
   ```
   Ini akan membuat folder `tmp` di dalam `src` yang berisi `main.exe` hasil build dan `builds-errors.log` untuk menyimpan error log ketika build `main.exe`
6. Jika ingin informasi lebih soal konfigurasi `.air.toml` misal merubah warna, mendifinisikan nama log dan direktorinya, memilih file dan folder yang akan di `watch` atau dengan kata lain diawasi perubahannya. Kunjungi

   [Hot Reload for Golang with Air -- byteSizeGo](https://www.bytesizego.com/blog/golang-air)

Reference :

- [Github Air](https://github.com/air-verse/air)
- [Medium - Berkenalan dengan Air Live Reload untuk Aplikasi Go](https://alitindrawan24.medium.com/berkenalan-dengan-air-live-reload-untuk-aplikasi-go-4cd1b4c16b6d)
