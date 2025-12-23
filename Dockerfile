# =========================
# Build stage
# =========================
FROM golang:1.22-alpine AS builder

WORKDIR /app

# Copy go mod files first (better caching)
COPY go.mod go.sum ./
RUN go mod download

# Copy source code
COPY . .

# Build static binary
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o arena-dash-server

# =========================
# Runtime stage
# =========================
FROM alpine:latest

WORKDIR /app

# Copy only the binary
COPY --from=builder /app/arena-dash-server .

# Server listens on 8080
EXPOSE 8080

# Run the server
CMD ["./arena-dash-server"]
