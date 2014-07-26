package main

import (
	"encoding/json"
	"io"
	"log"
	"net"
	"sync"
)

type Client struct {
	conn    net.Conn
	decoder *json.Decoder
}

var clients []*Client
var mutex sync.Mutex

func AddClient(conn net.Conn) *Client {
	c := &Client{
		conn:    conn,
		decoder: json.NewDecoder(conn),
	}

	mutex.Lock()
	clients = append(clients, c)
	mutex.Unlock()
	return c
}

func RemoveClient(c *Client) {
	mutex.Lock()
	defer mutex.Unlock()

	idx := -1
	for i := 0; i < len(clients); i++ {
		obj := clients[i]
		if obj == c {
			idx = i
			break
		}
	}

	if idx < 0 {
		return
	}

	if idx != len(clients)-1 {
		clients[idx] = clients[len(clients)-1]
	}

	clients = clients[:len(clients)-1]

	c.conn.Close()
}

func (c *Client) Write(msg []byte) int {
	size, err := c.conn.Write(msg)
	if err != nil {
		log.Print(err)
	}
	return size
}

func (c *Client) Listen() {
	defer RemoveClient(c)

	for {
		var data map[string]Command
		if err := c.decoder.Decode(&data); err != nil {
			// break the loop if it's a socket error
			_, netErr := err.(net.Error)
			if netErr == true || err == io.EOF {
				break
			} else {
				continue
			}
		}

		cmd, ok := data["action"]
		if ok {
			go warden.GPIO.Send(&cmd)
		}
	}
}
