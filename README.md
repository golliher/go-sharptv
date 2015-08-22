

# Overview

go-sharptv is your command line interface to your television set.  That is assuming you have a compatible TV from Sharp Electronics Corporation.

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

# Incompatible equipment
* Model LC-43LE653U  (2015 43") has been confired by Sharp representative to not support IP CONTROL.
Not all Smart TVs support IP CONTROL.   Some 2015 models do; check the detailed specs before purchasing if IP CONTROL is important to you.

# Installing

For now do something like this:
(assuming you have go(lang) installed and configured)

    go build -o go-sharptv main.go
    go install #  Or you could just... cp ./go-sharptv ~/bin

I will consider distributing binaries or via homebrew if there is interest.  File an issue if you would like to see that happen.

# Configuration

First, refer to your TV manual for how to enable network access and determine your TV's IP address.

You will want to configure go-sharptv for the IP address of your TV.  This can be done either view config file in ~/.sharptv/config. Or via environment variables.

The config file can be either config.yaml or config.json.

Example:

    ip: 192.168.1.2
    port: 10002

Environment variables are prefixed with GOSHARPTV_

Example:

    export GOSHARPTV_IP=192.168.1.2
    export GOSHARPTV_PORT=10002 #The default on my TV

# Example usage or Why do you want this?

You're in a shell window anyway and you want to control
your TV without moving your hands from the keyboard.

Or you want to automate something such as turning on or off your information radiator, or switching to the Weather Channel on the 8's or the news and then
back to your radiator.

If you are like me, you might use it to build up an Alfred workflow and maybe even and Alfred Remote panel.  

I use this in conjunction a separate system that changes the channels on my Tivo.
so I can do something like turn on the TV, Switch to the weather channel and set the
volume to a quiet level.

# Project Principals

1. Consistent CLI for great UX
2. Well documented
3. Solidly engineered - handles error cases, has tests
4. Automation for build and deployment
5. Configurable

Current status is that 1 & 5 are done.  I am not yet fulfilling my aspirations
on the others just yet. Pull requests welcome.

# Technologies used
* Vendoring https://github.com/kardianos/vendor
* Configuration https://github.com/spf13/viper
* CLI framework https://github.com/spf13/cobra

# License

Made available under a [MIT license ](LICENSE.md).

# Contributing

Please open issues at github with problems or feature requests.  Pull requests are welcome, though I beg your patience as I am a rookie open source maintainer.  

# Author

Darrell Golliher - http://golliher.net/

# Contributors
*  your_name_could_be_here
