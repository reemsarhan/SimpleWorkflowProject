
# Specify Base image we need for our Go app
FROM golang:1.20-alpine AS builder

# Create /app directory within the image to hold our app source code 
WORKDIR /app

# Copy everything in the root directory into our /app directory
COPY . .

# Install dependencies
RUN go mod download

# Build app with optional configuration
RUN go build -o /dashboard

# Tells Docker that the container listens to specified network ports at runtime
EXPOSE 8081

# Command to be used to execute when the image is used to start a container
CMD ["/dashboard"]
