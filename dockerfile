FROM golang:1.21.7-alpine3.18 as builder

WORKDIR /app

# Copy the go.mod and go.sum files to the workspace
COPY go.mod go.sum ./

# Download dependencies
RUN go mod download

# Copy the entire application
COPY . /app

# Build the Go application
RUN go build -o app

# Stage 2: Final Stage
FROM alpine:latest

WORKDIR /app

# Copy the binary from the builder stage to the final stage
COPY --from=builder /app/app .

# Expose the port the application runs on
EXPOSE 1323

# Command to run the application
CMD ["./app"]