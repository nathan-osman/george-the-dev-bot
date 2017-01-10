FROM ubuntu:16.04
MAINTAINER Nathan Osman <nathan@quickmediasolutions.com>

# Unlike nearly every other Dockerfile, we actually want the APT data to stick
# around in the cache so that dlocate can use it
RUN \
    apt-get update && \
    apt-get install -y \
        ca-certificates \
        dlocate

# Add the binary to the container
ADD dist/george /usr/local/bin/

# Specify the command for running the application
CMD george
