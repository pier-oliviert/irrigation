package main

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
	"log"
	"net"
	"os"
	"os/signal"
	"syscall"
)

var db *sql.DB
var ln net.Listener
var warden *Warden

func main() {
	fmt.Printf("Osmosis starting up...\n")

	log.SetFlags(log.LstdFlags | log.Lshortfile)

	sigc := make(chan os.Signal, 1)
	signal.Notify(sigc, os.Interrupt, os.Kill, syscall.SIGTERM)

	var err error
	db, err = sql.Open("postgres", "user=pothibo dbname=irrigation_dev sslmode=disable")
	handleFatalErr(err)

	go exit(sigc)

	warden = NewWarden(db, &GPIO{})

	ln, err = net.Listen("unix", "../tmp/sockets/osmosis.sock")
	handleFatalErr(err)

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
	ln.Close()
	warden.GPIO.Disconnect()
	os.Exit(0)
}

func handleFatalErr(err error) {
	if err != nil {
		log.Fatalf("*** Fatal Error: %s", err)
	}
}
