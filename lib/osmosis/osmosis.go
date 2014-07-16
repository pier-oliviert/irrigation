package main

import (
	"os"
	"os/signal"
	"log"
	"syscall"
	"fmt"
  "net"
	_ "github.com/lib/pq"
	"database/sql"
	"github.com/stianeikeland/go-rpio"
)

var db *sql.DB
var connections []net.Conn
var ln net.Listener
var warden *Warden

func main() {
	fmt.Printf("Osmosis starting up...\n")
	var err error
	db, err = sql.Open("postgres", "user=pothibo dbname=irrigation_dev sslmode=disable")

	if err != nil {
		log.Fatal(err)
	}

	activateGPIO()

	go exit()

	ln, err = net.Listen("unix", "../tmp/sockets/osmosis.sock")
	if err != nil {
		fmt.Printf("*** Fatal Error:  %s\n", err)
	}

	warden = NewWarden(db)

	for {
		conn, err := ln.Accept()
		connections = append(connections, conn)
		if err != nil {
			fmt.Printf("*** Connection Error:  %s\n", err)
			continue
		}
		go listen(conn)
	}
}

func exit() {
	sigc := make(chan os.Signal, 1)
	signal.Notify(sigc, os.Interrupt, os.Kill, syscall.SIGTERM)
	go func(c chan os.Signal) {
		sig := <-c
		log.Printf("Caught signal %s: shutting down.", sig)
		ln.Close()
		os.Exit(0)
	}(sigc)
}

func activateGPIO() error {
	zones, err := zones(db)

	if err != nil {
		return err
	}

	for zones.Next() {
		var gpio int
		if err := zones.Scan(&gpio); err != nil {
			return err
		}

		_ = rpio.Pin(gpio)
		//pin.Output()
	}
	return nil
}

func zones(db *sql.DB) (rows *sql.Rows, error error) {
	rows, err := db.Query("SELECT zones.gpio FROM zones;")
	if err != nil {
		return nil, err
	}
	return rows, nil
}
