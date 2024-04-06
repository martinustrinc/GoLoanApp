FROM golang:1.22.1 AS builder

# Set working directory
WORKDIR /app

# Copy source code
COPY go.mod go.sum ./
RUN go mod download

COPY *.go ./

# Build the application
RUN go build -o main

# Tahap 2: Runtime
FROM alpine:latest

# Set working directory
WORKDIR /app

# Copy executable from the builder stage
COPY --from=builder /app/main .

# Define command to execute when container starts
ENTRYPOINT ["./main"]
