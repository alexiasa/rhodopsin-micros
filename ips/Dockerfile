# golang image where workspace (GOPATH) configured at /go.
FROM golang

# Get dep for managing and restoring dependencies
RUN curl https://raw.githubusercontent.com/golang/dep/master/install.sh | sh

# Copy the local package files to the container’s workspace.
COPY . $GOPATH/src/rhodopsin-micros/ips

# Take ownership of copied files
RUN chown -R ${USER:=$(/usr/bin/id -run)}:$USER $GOPATH/src/

# Setting up working directory
WORKDIR $GOPATH/src/rhodopsin-micros/ips

# Check dependencies and install ips inside the container.
RUN dep ensure && go build .

# Make executables executable.
RUN chmod +x $GOPATH/src/rhodopsin-micros/ips/ips && chmod +x $GOPATH/src/rhodopsin-micros/ips/wait-for.sh

# Run ips when the container starts.
ENTRYPOINT $GOPATH/src/rhodopsin-micros/ips/ips

# Service listens on port 80.
EXPOSE 80

