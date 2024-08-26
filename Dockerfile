# Build stage
FROM golang:1.23-alpine AS build

# Install certificates
RUN apk --no-cache add ca-certificates

# Set the working directory
WORKDIR /app

# Copy the source code
COPY . .

# Build the Go application
RUN go build -ldflags="-s -w" -trimpath -o dist/go-fr cmd/go-fr/main.go
RUN go build -ldflags="-s -w" -trimpath -o dist/go-echo cmd/go-echo/main.go
RUN go build -ldflags="-s -w" -trimpath -o dist/go-bun cmd/go-bun/main.go
RUN go build -ldflags="-s -w" -trimpath -o dist/go-fuego cmd/go-fuego/main.go
RUN go build -ldflags="-s -w" -trimpath -o dist/go-fiber cmd/go-fiber/main.go
RUN go build -ldflags="-s -w" -trimpath -o dist/go-std cmd/go-std/main.go
RUN go build -ldflags="-s -w" -trimpath -o dist/go-gin cmd/go-gin/main.go
RUN go build -ldflags="-s -w" -trimpath -o dist/go-gorilla-mux cmd/go-gorilla-mux/main.go

# Production stage
FROM scratch as prod

# Set the working directory
WORKDIR /app

# Copy the certificates from the build stage
COPY --from=build /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/

# Copy the built binary from the build stage
COPY --from=build /app/dist /app


# Expose the port your application listens on
# EXPOSE 8080

# Run the application
# CMD ["/app/main"]

