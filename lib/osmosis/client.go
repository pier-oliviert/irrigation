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
var mutex sync.RWMutex

func AddClient(conn net.Conn) (*Client) {
  log.Print("Client connected")
  c := &Client{
    Conn: conn,
    decoder: json.NewDecoder(conn),
  }

  mutex.Lock()
  clients = append(clients, c)
  mutex.Unlock()
  return c
}

func RemoveClient(c *Client) {
  log.Print("Client disconnected")
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

  if idx != len(clients) - 1 {
    clients[idx] = clients[len(clients) - 1]
  }

  clients = clients[:len(clients) -1]

  c.Conn.Close()
}

func ExecuteOnClients(fn func([]*Client)) {
  mutex.Lock()
  fn(clients)
  mutex.Unlock()
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

    log.Println("Received command: ", cmd)
    if ok {
      if err := cmd.Execute(); err != nil {
        log.Print(err)
      }
    }
  }
}
