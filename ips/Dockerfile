# golang image where workspace (GOPATH) configured at /go.
FROM golang

# Copy the local package files to the container’s workspace.
ADD . rhodopsin-micros/ips

# Setting up working directory
WORKDIR rhodopsin-micros/ips

# Get dep for managing and restoring dependencies
RUN curl https://raw.githubusercontent.com/golang/dep/master/install.sh | sh

# Check dependencies
RUN dep ensure

# Build ips inside the container.
RUN go install rhodopsin-micros/ips

# Run ips when the container starts.
ENTRYPOINT rhodopsin-micros/ips

# Service listens on port 8080.
EXPOSE 8080