# Use the official golang image as the base
FROM golang:alpine AS builder

# Set working directory for build context
WORKDIR /app

# Copy go.mod and go.sum files (if using modules)
COPY go.mod go.sum ./

# Download dependencies (if using modules)
RUN go mod download

# Copy your application source code
COPY . .

# Build the Go application (replace main.go with your actual entrypoint)
RUN go build -o main cmd/server/main.go

# Use a slimmer alpine image for the final image
FROM alpine:edge

# Copy the built binary from the builder stage
COPY --from=builder /app/main /sbin/app

# Set the working directory for the application
WORKDIR /app

# Expose port 8080
EXPOSE 8080

# Command to run the application (replace app with your actual binary name)
CMD [ "/sbin/app" ]