FROM ubuntu:16.04
MAINTAINER Nathan Osman <nathan@quickmediasolutions.com>

RUN \
    apt-get update && \
    apt-get install -y \
        ca-certificates \
        dnsutils \
        iputils-ping

# Add the binary to the container
ADD dist/george /usr/local/bin/

# Specify the command for running the application
CMD george
