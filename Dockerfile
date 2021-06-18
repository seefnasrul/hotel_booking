# Start from golang base image
FROM golang:alpine

# ENV GO111MODULE=on

# Add Maintainer info
# LABEL maintainer="Seef <seefnasrul@gmail.com>"

# Install git.
# Git is required for fetching the dependencies.
RUN apk update && apk add --no-cache git

# Set the current working directory inside the container 
WORKDIR /opt/hotel-booking

# Copy go mod and sum files 
# COPY go.mod go.sum ./
COPY . .
# Download all dependencies. Dependencies will be cached if the go.mod and the go.sum files are not changed 
RUN go mod download 


# Install Air for hot reload
RUN go get -u github.com/cosmtrek/air

# The ENTRYPOINT defines the command that will be ran when the 
# container starts up
# In this case air command for hot reload go apps on file changes
ENTRYPOINT air

# # Copy the source from the current directory to the working Directory inside the container 
# COPY . .

# # Build the Go app
# # RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o main .
# # RUN CGO_ENABLED=0 GOOS=linux go build .

# # Start a new stage from scratch
# FROM alpine:latest
# RUN apk --no-cache add ca-certificates

# WORKDIR /root/

# # Copy the Pre-built binary file from the previous stage. Observe we also copied the .env file
# COPY --from=builder /app/ .
# COPY --from=builder /app/.env .       

# # Expose port 8080 to the outside world
# EXPOSE 8080

# #Command to run the executable
# CMD ["./main"]
