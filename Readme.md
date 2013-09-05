# Irrigation

Irrigation is a Go web application that allows you to control your landscaping irrigation valve via a raspberry Pi.

## Screenshots
![Irrigation](http://f.cl.ly/items/302d441S2P2a2R0F3Y1s/Screen%20Shot%202013-05-24%20at%201.45.59%20PM.png)

## Dependencies

Change pacman -Sy by your package manager

```bash
sudo pacman -Sy bzr mariadb go git pkg-config gcc
```

### MySQL (Mariadb)
Because SQLite3 is bugged in Go until at least 1.3, I had no choice to move away from it to a MySQL solution. 

When you install mariadb, it will ask you to run ```bash mysql_secure_installation```, do it. Once this is done, you can initialize your irrigation server.


## Installation

This installation assumes you know a bit about linux and you are running Arch Linux ARM. To get started you need to install Go 1.1. Install instruction can be found [here](http://golang.org/doc/install).

```bash
$ export GOPATH=~/go
$ cd $GOPATH
$ go get github.com/pothibo/irrigation
$ cd $GOPATH/src/github.com/pothibo/irrigation
$ go build && go install
$ $GOPATH/bin/irrigation -initialize -path="some/path"
```

The path is where irrigation will install the html, css, js, and config.yml files. These are the files you can modify if you want to
customize the look & feel of your app. 

By default, it will use ```~/irrigation```

The initialization process requires root access to your database. It will create a user and a two database (one for testing the other one for production).

## Activate your relays (_root_ privileges needed)

```bash
$ $GOPATH/bin/irrigation -activate -path="path/you/specified/above"
```

## Launch the server
If everything was successful, you can launch your server

```bash
$ $GOPATH/bin/irrigation -server &
```

By default, it will be available at ```raspberryPi_ip:7777```

## License
[MIT](http://pothibo.mit-license.org)
