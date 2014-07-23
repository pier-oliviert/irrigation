package main

import (
  "log"
  )

type Zone struct {
  Id int64 `json:"id"`
  Gpio int64 `json:"gpio"`
  State int64 `json:"state"`
}

func NewZone(id int64, gpio int64, pins []Pin) *Zone {
  z := &Zone{
    Id: id,
    Gpio: gpio,
  }

  z.extractPinInfo(pins)
  return z
}

func (z *Zone) extractPinInfo(pins []Pin) {
  for i := 0; i < len(pins); i++ {
    pin := pins[i]
    if pin.Id == z.Gpio {
      z.State = pin.State
    }
  }
}

func (z *Zone) HasActiveSchedule() bool {
  query, err := db.Query("select zones.gpio from zones inner join sprinkles on (sprinkles.zone_id = zones.id) where sprinkles.ends_at > CAST(NOW() at time zone 'utc' as timestamp);")
  defer query.Close()
  if err != nil {
    log.Print(err)
    return false
  }

  for query.Next() {
    return true
  }
  return false
}
