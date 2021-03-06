FROM ubuntu:16.04
MAINTAINER Nathan Osman <nathan@quickmediasolutions.com>

RUN \
    apt-get update && \
    apt-get install -y \
        apt-file \
        ca-certificates \
        dnsutils \
        iputils-ping && \
    apt-file update

# Add the binary to the container
ADD dist/george /usr/local/bin/

# Expose the status port
EXPOSE 8000

# Specify the command for running the application
CMD george
