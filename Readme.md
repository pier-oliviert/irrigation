# Irrigation

Irrigation is a Go web application that allows you to control your landscaping irrigation valve via a raspberry Pi.

## Usage

First build the project:
```bash
$ go build
```

Then create the sqlite database that will handle the schedules, events and valves:
```bash
$ ./irrigation -initdb
```

Now you can start the webserver and access it through raspberryPi_IP:7777:
```bash
$ ./irrigation -server
```

## Dependencies

Gorilla/Pat
Gopi
Gorp
