# Build stage
FROM golang:1.24-alpine AS builder

WORKDIR /app

RUN apk add --no-cache git ca-certificates

COPY go.mod ./
# COPY go.sum ./
RUN go mod download || true

COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -ldflags="-w -s" -o /bin/restoreforge-api ./cmd/api
RUN CGO_ENABLED=0 GOOS=linux go build -ldflags="-w -s" -o /bin/restoreforge-worker ./cmd/worker
RUN CGO_ENABLED=0 GOOS=linux go build -ldflags="-w -s" -o /bin/restoreforge ./cmd/restoreforge

# Final runtime stage
FROM alpine:3.20

RUN apk add --no-cache ca-certificates tzdata

WORKDIR /app

COPY --from=builder /bin/restoreforge-api /app/restoreforge-api
COPY --from=builder /bin/restoreforge-worker /app/restoreforge-worker
COPY --from=builder /bin/restoreforge /app/restoreforge
COPY migrations/ /app/migrations/

EXPOSE 8080

ENTRYPOINT ["/app/restoreforge-api"]

