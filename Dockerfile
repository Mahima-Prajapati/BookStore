# Multi-stage Dockerfile for BookStore Go app

# 1) Build stage
FROM golang:1.20-alpine AS builder
WORKDIR /src

# Cache and download dependencies
COPY go.mod go.sum ./
RUN go mod download

# Copy the rest and build
COPY . .
# Build a small static binary
RUN CGO_ENABLED=0 GOOS=linux go build -ldflags="-s -w" -o /bookstore ./cmd

# 2) Production image
FROM alpine:3.18
RUN apk add --no-cache ca-certificates

# Copy the binary
COPY --from=builder /bookstore /usr/local/bin/bookstore

# Expose the port the app will run on
EXPOSE 9010
ENV PORT=9010

# Use a non-root user (optional)
USER 1000

ENTRYPOINT ["/usr/local/bin/bookstore"]
