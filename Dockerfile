FROM golang:latest

WORKDIR /app

# Copy go mod and source files
COPY go.mod ./
COPY main.go ./

# Download dependencies
RUN go mod download

# Build the application
RUN go build -o go-server main.go

# Expose port 8080
EXPOSE 8080

# Run the application
CMD ["./go-server"]
