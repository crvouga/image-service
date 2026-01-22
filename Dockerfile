# Build stage
FROM golang:1.24.2-alpine AS builder

# Set working directory
WORKDIR /app

# Install git (needed for some Go dependencies)
RUN apk add --no-cache git

# Copy go mod files
COPY go.mod ./

# Download dependencies
RUN go mod download

# Copy source code
COPY . .

# Build the application
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o image-service ./main.go

# Runtime stage
FROM alpine:latest

# Install ca-certificates for HTTPS requests and wget for healthcheck
RUN apk --no-cache add ca-certificates wget

WORKDIR /app

# Copy the binary from builder
COPY --from=builder /app/image-service .

# Copy the app directory (contains HTML templates and other assets)
COPY --from=builder /app/app ./app

# Copy other required files
COPY --from=builder /app/robots.txt ./robots.txt
COPY --from=builder /app/keyValueDB.json ./keyValueDB.json

# Expose port 80
EXPOSE 80

# Set environment variable for port
ENV PORT=80

# Run the application
CMD ["./image-service"]
