# Irrigation

Irrigation is a Go web application that allows you to control your landscaping irrigation valve via a raspberry Pi.

## Screenshots
![Irrigation](http://f.cl.ly/items/302d441S2P2a2R0F3Y1s/Screen%20Shot%202013-05-24%20at%201.45.59%20PM.png)

## Installation

This installation assumes you know a bit about linux and you are running Arch Linux ARM. To get started you need to install Go 1.1. Install instruction can be found (http://golang.org/doc/install)[here].

Once installed, make sure you have added a $GOPATH environment variable. Then go to your $GOPATH and fetch this project and its dependencies.
```bash
$ cd $GOPATH/
$ go get github.com/pothibo/irrigation
$ cd src/github.com/pothibo/irrigation
$ go get ./..
```

Installation for this project is done in 2 steps. First, we want to copy all the assets so they will be available to the executable and then we want to build & install the server's executable.

```bash
$ ./assets.sh && go install
```

I assume you don't have $GOPATH in your path. Let's go to $GOPATH/bin folder and initialize the server.

```bash
$ cd $GOPATH/bin
$ ./irrigation -initdb
```

Now your server should be configured and ready to go. Start the webserver and access it through raspberryPi_IP:7777:
```bash
$ ./irrigation -server
```

