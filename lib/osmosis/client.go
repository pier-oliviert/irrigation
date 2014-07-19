package main

import (
  "log"
  "strings"
  "net"
  )

type Client struct {
  Conn net.Conn
  SendCmd chan string
}

var clients []*Client

func AddClient(conn net.Conn) (*Client) {
  log.Print("Client connected")
  c := &Client{
    Conn: conn,
    SendCmd: make(chan string),
  }

  clients = append(clients, c)

  return c
}

func RemoveClient(c *Client) {
  log.Print("Client Deconnected")
  idx := -1
  for i := 0; i < len(clients); i++ {
    obj := clients[i]
    if obj == c {
      idx = i
      break
    }
  }

  if idx != len(clients) - 1 {
    clients[idx] = clients[len(clients) - 1]
  }

  clients = clients[:len(clients) -1]

  c.Conn.Close()
}

func (c *Client) Read(buffer []byte) (int, error) {
  bytesRead, err := c.Conn.Read(buffer)
  if err != nil {
    return 0, err
  }

  return bytesRead, nil
}

func (c *Client) Listen() {
  defer RemoveClient(c)

  for {
    buf := make([]byte, 32)
    n, err := c.Read(buf)
    if err != nil {
      log.Print(err)
      return
    }

    command, err := NewCommand(strings.Split(string(buf[0:n]), ":"))

    if err != nil {
      log.Print(err)
      return
    }

    err = command.Execute()

    if err != nil {
      log.Print(err)
      return
    }

    notify(clients, command)
  }
}
