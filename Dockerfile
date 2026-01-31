# Stage 1: Build
FROM golang:1.25.5-alpine AS builder

# Set working directory
WORKDIR /app

# Install dependencies
RUN apk add --no-cache git

# Copy go mod files
COPY go.mod go.sum ./

# Download dependencies
RUN go mod download

# Copy source code
COPY . .

# Build the application
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o backend-app .

# Stage 2: Run
FROM alpine:latest

# Install ca-certificates for HTTPS
RUN apk --no-cache add ca-certificates

# Set working directory
WORKDIR /root/

# Copy binary from builder
COPY --from=builder /app/main .

# Copy .env file (optional, can use environment variables instead)
COPY --from=builder /app/.env .

# Expose port
EXPOSE 8282

# Run the application
CMD ["./main"]