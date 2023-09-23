# Use the official Golang image with Debian as a base to build the app
FROM golang:1.21.1-bullseye AS build-env

# Set the working directory inside the container
WORKDIR /app

# Copy all the files from the host to the container
COPY . .

# Build the Go app
RUN go mod download
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o main .

# Start from a scratch image
FROM scratch

# Copy the binary from the build-env
COPY --from=build-env /app/main /main

# Copy static files and templates
COPY --from=build-env /app/templates /templates
COPY --from=build-env /app/static /static

# Copy CA certificates for HTTPS requests from the Debian base image
COPY --from=build-env /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/

EXPOSE 9123

# Run the app
CMD ["/main"]
