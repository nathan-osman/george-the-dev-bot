FROM ubuntu:16.04
MAINTAINER Nathan Osman <nathan@quickmediasolutions.com>

# Add the binary to the container
ADD dist/george /usr/local/bin/

# Specify the command for running the application
CMD george
