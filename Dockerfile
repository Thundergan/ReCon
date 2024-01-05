FROM golang:latest

# Set the username, uid and gid the docker container is run with
ARG UNAME=recon
ARG UID=1000
ARG GID=1000

RUN groupadd -g $GID -o $UNAME
RUN useradd -m -u $UID -g $GID -o $UNAME
USER $UNAME

# Set destination for COPY
WORKDIR /app

# Copy the dev dir to /app
COPY . .
# Download Go modules
RUN go mod download

# Build
RUN make build
ENV USER=root

# Optional:
# To bind to a TCP port, runtime parameters must be supplied to the docker command.
# But we can document in the Dockerfile what ports
# the application is going to listen on by default.
# https://docs.docker.com/engine/reference/builder/#expose
EXPOSE 8080

# Run
CMD ["/bin/bash"]
#CMD ["/recon/build/recon"]
