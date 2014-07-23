package main

import (
  "log"
  "net"
  "encoding/json"
  "io"
  "sync"
  )

type Client struct {
  Conn net.Conn
  decoder *json.Decoder
}

var clients []*Client

func AddClient(conn net.Conn) (*Client) {
  log.Print("Client connected")
  c := &Client{
    Conn: conn,
    decoder: json.NewDecoder(conn),
  }

  clients = append(clients, c)

  return c
}

func RemoveClient(c *Client) {
  log.Print("Client Disconnected")
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

  c.Conn.Close()
}

func Clients() {
  mutex := sync.RWMutex
  mutex.Lock()
  c = clients
  mutex.Unlock()
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
      if err := cmd.Execute(); err != nil {
        log.Print(err)
      }
    }
  }
}
