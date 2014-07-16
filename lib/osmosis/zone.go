package main

type Zone struct {
  id int
  gpio int
  active bool
}

func NewZone(id int, gpio int) *Zone {
  z := new(Zone)
  z.id = id
  z.gpio = gpio
  return z
}
