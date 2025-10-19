# Build stage
FROM golang:1.25-alpine AS builder

WORKDIR /app

RUN apk add --no-cache gcc g++ make git

# Copy go mod files
COPY go.mod go.sum ./

# Download dependencies
RUN go mod download

# Copy source code
COPY . .

# Set environment variables for Go build
ENV CGO_ENABLED=0 GOOS=linux GOARCH=amd64

# Build the Go binary
RUN GOEXPERIMENT=greenteagc go build -ldflags='-s -w -extldflags "-static"' -o sample-release-app ./cmd/sample-release-app

############################
# STEP 2 build a small image
############################
FROM alpine:3.20

# Set the working directory in the final container
WORKDIR /root/

# Copy the binary from builder
COPY --from=builder /app/sample-release-app .

# Expose the port
EXPOSE 8080

# Command to run the binary
CMD ["./sample-release-app"]
