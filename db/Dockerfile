FROM mongo:latest

# copy initialization script
COPY initdata.js /docker-entrypoint-initdb.d/

# Install updates (slowly)
# RUN apt-get update && apt-get upgrade -y
