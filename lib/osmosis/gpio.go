package main

import (
  "net"
  "log"
  "encoding/json"
  )

type GPIO struct {
  conn net.Conn
  decoder *json.Decoder
}

type Pin struct {
  Id int64 `json:"id"`
  State int64 `json:"state"`
}

func NewGPIO(c net.Conn) *GPIO {
  gpio = &GPIO{
    conn: c,
    decoder: json.NewDecoder(c),
  }

  go gpio.listen()

  return gpio
}

func (g *GPIO) Disconnect() {
  g.conn.Close()
}

func (g *GPIO) Send(cmd *Command) {
  log.Print("Sending command")
  g.conn.Write([]byte(cmd.Bytes()))
}

func (g *GPIO) listen() {
  defer g.Disconnect()
  for {
    var pins []Pin
    if err := g.decoder.Decode(&pins); err != nil {
      return
    }

    warden.Pins <- pins
  }
}
