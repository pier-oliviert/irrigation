package main

import (
	"encoding/json"
	"errors"
	"log"
	"net"
	"time"
	"io"
)

type GPIO struct {
	conn    net.Conn
	decoder *json.Decoder
}

func (g *GPIO) StartListening(ch chan []Pin) {
	c, err := net.Dial("unix", "/tmp/gobble.sock")
	handleFatalErr(err)
	g.decoder = json.NewDecoder(c)
	g.conn = c
	go g.update()
	go g.listen(ch)
}

func (g *GPIO) Disconnect() {
	g.conn.Close()
}

func (g *GPIO) Send(c *Command) error {
	ch := make(chan int, 1)

	select {
	case ch <- g.write(c.Bytes()):
	case <-time.After(5 * time.Second):
		return errors.New("Couldn't write command on socket")
	}

	return nil
}

func (g *GPIO) listen(ch chan []Pin) {
	defer g.Disconnect()
	for {
		var pins []Pin
		if err := g.decoder.Decode(&pins); err != nil {
			// break the loop if it's a socket error
			_, netErr := err.(net.Error)
			if netErr == true || err == io.EOF {
				break
			} else {
				continue
			}
		}
		ch <- pins
	}
}

func (g *GPIO) update() {
	for _ = range time.Tick(time.Second) {
		cmd := &Command{Name: "List"}
		g.write(cmd.Bytes())
	}
}

func (g *GPIO) write(msg []byte) int {
	size, err := g.conn.Write(msg)
	if err != nil {
		log.Fatal(err)
	}
	return size
}
