# Use the official Golang alpine image from the Docker Hub
FROM golang:1.22-alpine

# Arguments for username, UID, GID and port
ARG USERNAME
ARG UID
ARG GID
ARG PORT

# Install ssh
RUN apk add --no-cache openssh-client

# Create a group and user
RUN addgroup -g $GID $USERNAME && \
    adduser -D -u $UID -G $USERNAME $USERNAME

# Set the working directory to the home directory of the user
WORKDIR /home/$USERNAME

# Copy the local package files to the container's workspace.
COPY . /home/$USERNAME

# Set Gin to release mode
ENV GIN_MODE=release

# Download dependencies
RUN go mod tidy

# Build the Go app
RUN go build -o main .

# Run the command as the specified user
USER $USERNAME

# Expose the application on the specified port
EXPOSE $PORT

# Run the Go app
CMD ["./main"]