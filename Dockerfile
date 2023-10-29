# Use an official Golang runtime as a parent image
FROM golang:1.21-alpine

# Set the working directory in the container
WORKDIR /app

# Install a file watcher for live reloading (e.g., air)
RUN go install github.com/cosmtrek/air@latest

# Copy go.mod and go.sum files and download dependencies
COPY go.mod go.sum ./
RUN go mod download

# Copy the rest of your application code
COPY . .

# Start your application with the live reload tool (e.g., air)
CMD ["air", "-c", ".air.toml"]
    