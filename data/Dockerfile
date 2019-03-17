# golang image where workspace (GOPATH) configured at /go.
FROM golang:latest

RUN apt-key adv --keyserver hkp://keyserver.ubuntu.com:80 --recv 9DA31620334BD75D9DCB49F368818C72E52529D4
RUN echo "deb http://repo.mongodb.org/apt/debian stretch/mongodb-org/4.0 main" | tee /etc/apt/sources.list.d/mongodb-org-4.0.list
RUN apt-get update
RUN apt-get install -y mongodb-org-tools cron

# Get dep for managing and restoring dependencies
RUN curl https://raw.githubusercontent.com/golang/dep/master/install.sh | sh

# Copy the local package files to the container’s workspace.
COPY . $GOPATH/src/github.com/alexiasa/rhodopsin-micros/data

# Take ownership of copied files
RUN chown -R ${USER:=$(/usr/bin/id -run)}:$USER $GOPATH/src/

# Setting up working directory
WORKDIR $GOPATH/src/github.com/alexiasa/rhodopsin-micros/data

# Check dependencies and install ips inside the container.
RUN dep ensure && go build .

# Make executables executable.
RUN chmod +x $GOPATH/src/github.com/alexiasa/rhodopsin-micros/data/data

# Create hourly cron script for data
#ADD datacron /etc/cron.d/datacron
#RUN chmod 0644 /etc/cron.d/datacron

# Run data when the container starts and then run cron to keep data running hourly...eventually maybe
ENTRYPOINT $GOPATH/src/github.com/alexiasa/rhodopsin-micros/data/data