# golang image where workspace (GOPATH) configured at /go.
FROM golang:latest

RUN apt-get update && apt-get install mongo

# Get dep for managing and restoring dependencies
RUN curl https://raw.githubusercontent.com/golang/dep/master/install.sh | sh

# Copy the local package files to the container’s workspace.
COPY . $GOPATH/src/rhodopsin-micros/data

# Take ownership of copied files
RUN chown -R ${USER:=$(/usr/bin/id -run)}:$USER $GOPATH/src/

# Setting up working directory
WORKDIR $GOPATH/src/rhodopsin-micros/data

# Check dependencies and install ips inside the container.
RUN dep ensure && go build .

# Make executables executable.
RUN chmod +x $GOPATH/src/rhodopsin-micros/data/data && chmod +x $GOPATH/src/rhodopsin-micros/data/datacheck.sh

# Run data when the container starts.
ENTRYPOINT $GOPATH/src/rhodopsin-micros/data/data

