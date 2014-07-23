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
)

var db *sql.DB
var ln net.Listener
var warden *Warden
var gpio *GPIO

func main() {
	fmt.Printf("Osmosis starting up...\n")
	log.SetFlags(log.LstdFlags|log.Lshortfile)

	sigc := make(chan os.Signal, 1)
	signal.Notify(sigc, os.Interrupt, os.Kill, syscall.SIGTERM)

	var err error
	db, err = sql.Open("postgres", "user=pothibo dbname=irrigation_dev sslmode=disable")
	handleFatalErr(err)

	go exit(sigc)
	conn, err := net.Dial("unix", "/tmp/gobble.sock")
	handleFatalErr(err)

	gpio = NewGPIO(conn)

	ln, err = net.Listen("unix", "../tmp/sockets/osmosis.sock")
	handleFatalErr(err)

	go StartWarden(db)

	for {
		conn, err := ln.Accept()
		if err != nil {
			fmt.Printf("*** Connection Error:  %s\n", err)
			continue
		}
		client := AddClient(conn)
		go client.Listen()
	}

	sigc <- syscall.SIGINT
}

func exit(c chan os.Signal) {
	sig := <-c
	log.Printf("Caught signal %s: shutting down.", sig)
	gpio.Disconnect()
	ln.Close()
	os.Exit(0)
}

func zones(db *sql.DB) (rows *sql.Rows, error error) {
	rows, err := db.Query("SELECT zones.gpio FROM zones;")
	if err != nil {
		return nil, err
	}
	return rows, nil
}

func handleFatalErr(err error) {
	if err != nil {
		log.Fatalf("*** Fatal Error: %s", err)
	}
}
