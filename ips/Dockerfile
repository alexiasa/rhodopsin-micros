# golang image where workspace (GOPATH) configured at /go.
FROM golang

# Copy the local package files to the container’s workspace.
ADD . rhodopsin-micros/ips

# Setting up working directory
WORKDIR rhodopsin-micros/ips

# Get godeps for managing and restoring dependencies
RUN go get github.com/tools/godep

# Restore godep dependencies
RUN godep restore

# Build the taskmanager command inside the container.
RUN go install rhodopsin-micros/ips

# Run the taskmanager command when the container starts.
ENTRYPOINT rhodopsin-micros/ips

# Service listens on port 8080.
EXPOSE 8080