# Use Go base image
FROM golang:1.22.6

# Set working directory
WORKDIR /app

# Copy all files to the container
COPY . .

# Install dependencies
RUN go mod download

# List files in /app to debug
RUN ls -l /app

# Compile the application from the cmd directory
RUN go build -o /app/main ./cmd

# Run the application
CMD ["./main"]
