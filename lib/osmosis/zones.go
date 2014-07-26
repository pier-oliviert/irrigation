package main

import (
  "sync"
  "time"
  )

type Zones struct {
  zones map[int64]*Zone
  Update chan []Pin
  mutex sync.Mutex
}

func NewZones() *Zones {
  z := &Zones{
    zones: AllZones(),
    Update: make(chan []Pin),
    }

  go z.update()
  return z
}

func (z *Zones) All() []*Zone {
  var list []*Zone
  for _, zone := range z.zones {
    list = append(list, zone)
  }
  return list
}

func (z *Zones) update() {
  z.mutex.Lock()
  for _ = range time.Tick(time.Second) {
    for id, zone := range AllZones() {
      if z.zones[id] == nil {
        z.zones[id] = zone
      }
    }
  }
  z.mutex.Unlock()
}

func (z *Zones) getZoneByPinId(id int64) *Zone {
  for _, value := range z.zones {
    if value.Gpio == id {
      return value
    }
  }
  return nil
}

func (z *Zones) States(ch chan<- *Zone) {
  for {
    for _, pin := range <- z.Update {
      zone := z.getZoneByPinId(pin.Id)
      if zone != nil && zone.SetState(pin.State) {
        ch <- zone
      }
    }
  }
}
