# Build stage
FROM golang:1.21-alpine AS builder

WORKDIR /app

# Copy go mod and sum files
COPY go.mod go.sum ./

# Download dependencies
RUN go mod download

# Copy source code
COPY . .

# Build the application
RUN CGO_ENABLED=0 GOOS=linux go build -o /app/weather-api ./cmd/api

# Final stage
FROM alpine:3.18

WORKDIR /app

# Copy the binary from builder
COPY --from=builder /app/weather-api .
COPY .env .

# Install CA certificates for HTTPS requests
RUN apk --no-cache add ca-certificates

# Expose the application port
EXPOSE 8080

# Run the application
CMD ["./weather-api"]
