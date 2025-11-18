# Stage 1: Build
FROM golang:1.24-alpine AS builder

WORKDIR /app

ARG ENV=dev

COPY src/go.mod ./
COPY src/ ./
COPY src/internal/creds/gcp_firebase.${ENV}.json internal/creds/gcp_firebase.json
COPY src/internal/creds/gcp_bucket.${ENV}.json internal/creds/gcp_bucket.json

COPY src/.env.${ENV} .env

RUN go mod tidy && go mod download
RUN go build -o server cmd/main.go

# Stage 2: Run
FROM alpine:latest

WORKDIR /app

COPY --from=builder /app/server .
COPY --from=builder /app/.env .env
COPY --from=builder /app/internal/creds/gcp_firebase.json ./internal/creds/gcp_firebase.json
COPY --from=builder /app/internal/creds/gcp_bucket.json ./internal/creds/gcp_bucket.json

RUN apk --no-cache add ca-certificates

EXPOSE 8080

CMD ["./server"]
