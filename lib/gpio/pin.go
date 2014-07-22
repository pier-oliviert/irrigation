package main

import (
  "github.com/stianeikeland/go-rpio"
  "encoding/json"
)

type Pin struct {
  json.Marshaler
  gpio rpio.Pin
}

func NewPin(id int64) *Pin {
  p := &Pin{}
  
  p.gpio = rpio.Pin(id)
  p.gpio.Output()

  return p
}

func (p *Pin) Id() int8 {
  return int8(p.gpio)
}

func (p *Pin) Open() {
  p.gpio.High()
}

func (p *Pin) Close() {
  p.gpio.Low()
}

func (p *Pin) State() int {
  return int(p.gpio.Read())
}

func (p *Pin) MarshalJSON() ([]byte, error) {
  return json.Marshal(struct{
    Id int `json:"id"`
    State int `json:"state"`
  }{
    State: p.State(),
    Id: int(p.gpio),
  })
}
