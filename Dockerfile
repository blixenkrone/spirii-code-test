# Stage 1: Build
FROM golang:1.23-alpine AS builder

# Set environment
ARG APP_ENV=local
ENV APP_ENV=${APP_ENV}

# Set the working directory
WORKDIR /src

# Copy go.mod and go.sum
COPY go.mod go.sum ./

# Download dependencies
RUN go mod download

# Copy the rest of the application code
COPY . .

# Build the Go app for ARM64 architecture
RUN CGO_ENABLED=0 GOOS=linux GOARCH=arm64 go build -o /app/main ./cmd/main.go

# Stage 2: Run
FROM alpine:latest

# Set workdir in final image
WORKDIR /app

# Copy the built Go binary from builder
COPY --from=builder /app/main /app/main

# Run the binary
ENTRYPOINT ["/app/main"]
