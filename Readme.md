# Irrigation

Irrigation is a Go web application that allows you to control your landscaping irrigation valve via a raspberry Pi.

## Screenshots
![Irrigation](http://f.cl.ly/items/302d441S2P2a2R0F3Y1s/Screen%20Shot%202013-05-24%20at%201.45.59%20PM.png)

## Dependencies

Change pacman -Sy by your package manager

```bash
sudo pacman -Sy bzr sqlites3 go git pkg-config gcc
```

## Installation

This installation assumes you know a bit about linux and you are running Arch Linux ARM. To get started you need to install Go 1.1. Install instruction can be found (http://golang.org/doc/install)[here].

```bash
$ export GOPATH=~/go
$ cd $GOPATH
$ go get github.com/pothibo/irrigation
```

Create a folder ```/srv/http/irrigation``` and symlink assets/ folder and config.yml:

```bash
sudo mkdir /srv/http/irrigation && sudo chown your_user /srv/http/irrigation && sudo chgrp http /srv/http/irrigation
ln -s $GOPATH/src/github.com/pothibo/irrigation/assets/ /srv/http/irrigation/assets
ln -s $GOPATH/src/github.com/pothibo/irrigation/config.yml /srv/http/irrigation/config.yml
```

Activate the relays first (_root_ privileges needed)

```bash
$ cd $GOPATH/bin
# Activate the relay based on config.yml
$ sudo ./irrigation -activate

# Initialize Sqlite database
./irrigation -initdb

#Run the server
./irrigation -server
```

