

# Overview

go-sharptv is your command line interface to your Sharp or Aquos television set.  That is assuming you have a compatible TV from Sharp Inc.

go-sharptv is a hobbyist project by an owner of a Sharp brand TV for other owners of Sharp brand TVs.    This project is not affiliated with Sharp Inc. in any way.

# Features

* Power: on, off, toggle, status
* Mute: on, off, toggle, status
* Input switching
* Volume up, down, set to value
* Consistent CLI for great UX
* Configurable
  - config file
  - environment
  - flags

# Compatible equipment

Originally developed and tested for the one SharpTV I own:

* Model LC-70C6400U
* Protocol Version 0100
* Firmware Version 222U1302091

# Installing

I may later provide binaries and a homebrew recipe.  For now
do something like this:

    go build -o sharptv main.go
    cp ./sharptv ~/bin

# Configuration

You will want to configure the IP address of your TV.  This can be done either view config file in ~/.sharptv/config. Or via environment variables.

The config file can be either config.yaml or config.json.

Example:

    ip: 192.168.1.2
    port: 10002

Environment variables are prevised with GOSHARPTV_

Example:

    export GOSHARPTV_IP=192.168.1.2
    export GOSHARPTV_PORT=10002 #The default on my TV

<!-- # Example usage

## Advanced usage
### Alfred integration
### Alfred remote integration -->

# Project Principals

1. Consistent CLI for great UX
2. Well documented
3. Solidly engineered - handles error cases, has tests
4. Automation for build and deployment
5. Configurable

Current status is that 1 & 5 are done, but 2-4 need work.

# License

Copyright 2015 Darrell Golliher. All rights reserved until I pick a suitable open source license.

# Contributing

Please open issues at github with problems or feature requests.  Pull requests are welcome, though I beg your patience as I am a rookie open source maintainer.

# Author(s)

Darrell Golliher - http://golliher.net/
