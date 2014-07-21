package main

import (
  "net"
  "strings"
  "strconv"
  )

type GPIO struct {
  conn net.Conn
  Notify chan string
}

func NewGPIO(c net.Conn) *GPIO {
  gpio = &GPIO{
    conn: c,
    Notify: make(chan string),
  }

  go gpio.listen()

  return gpio
}

func (g *GPIO) Disconnect() {
  g.conn.Close()
}

func (g *GPIO) Open(id int64) {
  var event []string
  event = append(event, "action")
  event = append(event, strconv.FormatInt(id, 10))
  event = append(event, "open")

  g.conn.Write([]byte(strings.Join(event, ":")))
}

func (g *GPIO) Close(id int64) {
  var event []string
  event = append(event, "action")
  event = append(event, strconv.FormatInt(id, 10))
  event = append(event, "close")

  g.conn.Write([]byte(strings.Join(event, ":")))
}

func (g *GPIO) GetCurrentStatus() {
  var event []string
  event = append(event, "status")
  event = append(event, "list")

  g.conn.Write([]byte(strings.Join(event, ":")))
}

func (g *GPIO) listen() {
  defer g.Disconnect()
  for {
    buf := make([]byte, 1024 * 32) //32k for now... should be enough.
    n, err := g.conn.Read(buf)
    if err != nil {
      return
    }
    warden.Notify <- buf[0:n]
  }
}
