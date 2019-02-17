# golang image where workspace (GOPATH) configured at /go.
FROM golang

# Copy the local package files to the container’s workspace.
ADD . $GOPATH/src/rhodopsin-micros/ips

# Setting up working directory
WORKDIR $GOPATH/src/rhodopsin-micros/ips

# Get dep for managing and restoring dependencies
RUN curl https://raw.githubusercontent.com/golang/dep/master/install.sh | sh

# Check dependencies
RUN dep ensure

# Build ips inside the container.
RUN go install $GOPATH/src/rhodopsin-micros/ips

# Run ips when the container starts.
# ENTRYPOINT $GOPATH/src/rhodopsin-micros/ips
CMD $GOPATH/src/rhodopsin-micros/ips


# Service listens on port 8080.
EXPOSE 8080

