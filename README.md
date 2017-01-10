## George the Dev Bot

George is a bot for the Stack Exchange chat network. It is written in Go and uses the excellent [go-sechat](https://github.com/nathan-osman/go-sechat) package for interacting with the chat server.

This application is based on earlier work and is still a WIP.

### Building

A makefile is included to simplify the process of building the application. Assuming you have Docker installed, the command is simply:

    make

After this completes, you will find the binary in `dist/`.
