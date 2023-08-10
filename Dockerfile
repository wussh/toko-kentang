# Use an official Golang runtime as a parent image
FROM golang:1.20-alpine

# Set the working directory inside the container
WORKDIR /go/src/app

# Copy the local package files to the container's workspace
COPY . .

# Build the Go application
RUN go build -o main .

# Set the command to run the executable
CMD ["./main"]
