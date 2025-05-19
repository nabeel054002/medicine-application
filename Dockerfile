# Stage 1: Build
FROM golang:1.24.3 AS builder

WORKDIR /app

# Copy go.mod and go.sum and download dependencies
COPY go.mod go.sum ./
RUN go mod download

# Copy entire project
COPY . .

# Build with CGO enabled targeting Linux
RUN CGO_ENABLED=1 GOOS=linux go build -a -installsuffix cgo -o server ./cmd/api

# Stage 2: Runtime
FROM debian:bookworm-slim

# Install ca-certificates and sqlite3 runtime library
RUN apt-get update && \
    apt-get install -y ca-certificates libsqlite3-0 && \
    rm -rf /var/lib/apt/lists/*

WORKDIR /app

# Copy built binary from builder stage
COPY --from=builder /app/server .

EXPOSE 8080

CMD ["./server"]
