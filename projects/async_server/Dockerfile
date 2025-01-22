# Stage 1: Build the Go binary
FROM golang:1.23-alpine AS builder

# Set the working directory inside the container
WORKDIR /app

COPY go.* ./
COPY models models
COPY examples/long_running/* ./

# get dependencies
RUN go mod tidy

# Build the Go application
RUN CGO_ENABLED=0 GOOS=linux GOARCH=arm64 go build -o server .

# Stage 2: Create a minimal image to run the Go binary
FROM alpine:latest

# Set the working directory inside the container
WORKDIR /root/

# Copy the binary from the builder stage
COPY --from=builder /app/server .

# Expose the port that your Go server listens on
EXPOSE 8080

# Command to run the executable
CMD ["./server"]
