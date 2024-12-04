# Start from the official Go image
FROM golang:alpine
# Set the Current Working Directory inside the container
WORKDIR /src
# Copy go.mod and go.sum files
COPY go.mod go.sum ./
# Download all dependencies. Dependencies will be cached if the go.mod and go.sum files are not changed
RUN go mod download
# Copy the source from the current directory to the Working Directory inside the container
COPY . .

RUN CGO_ENABLED=0 go install -ldflags "-s -w -extldflags '-static'" github.com/go-delve/delve/cmd/dlv@latest

# Build the Go app in debug
RUN CGO_ENABLED=0 go build -gcflags "all=-N -l" -o ./main ./src/main.go
# Expose port 8080 to the outside world
EXPOSE 8080 4000
# Set environment variable for Gin mode
ENV GIN_MODE=release
# Run the executable
CMD [ "/go/bin/dlv", "--listen=:4000", "--headless=true", "--log=true", "--accept-multiclient", "--api-version=2", "exec", "./main" ]





