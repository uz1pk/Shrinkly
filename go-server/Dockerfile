# Use the official Golang image as a base
FROM golang:1.17

# Set the working directory in the container
WORKDIR /app

# Copy the source code into the container
COPY . .

# Build the Golang application
RUN go build -o main .

# Expose the server's port
EXPOSE 8080

# Start the server
CMD ["./main"]
