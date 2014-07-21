package main

import (
  "log"
  "net"
  "strings"
  "strconv"
  )

type Client struct {
  Conn net.Conn
  Notify chan string
}

var clients []*Client

func AddClient(conn net.Conn) (*Client) {
  log.Print("Client connected")
  c := &Client{
    Conn: conn,
    Notify: make(chan string),
  }

  clients = append(clients, c)

  go c.Listen()
  go c.listenForUpdate()
  return c
}

func RemoveClient(c *Client) {
  log.Print("Client disconnected")
  c.Conn.Close()
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

  if idx != len(clients) - 1 {
    clients[idx] = clients[len(clients) - 1]
  }

  clients = clients[:len(clients) -1]

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
      return
    }

    c.execute(strings.Split(string(buf[0:n]), ":"))
  }
}

func (c *Client) execute(instructions []string) {
  if len(instructions) < 3 {
    ListGPIO()
  } else {
    id, err := strconv.ParseInt(instructions[1], 10, 16)
    if err != nil {
      ListGPIO()
    }
    switch instructions[2] {
      case "open": OpenGPIO(id)
      case "close": CloseGPIO(id)
      default: ListGPIO()
    }
  }
}

func (c *Client) listenForUpdate() {
  for {
    msg := <-c.Notify
    c.Conn.Write([]byte(msg))
  }
}
