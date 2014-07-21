package main

import (
  "time"
  _ "github.com/lib/pq"
  "database/sql"
  "encoding/json"
  "log"
  )

type Warden struct {
  gpios []int
  Notify chan []byte
}

type Pin struct {
  Id int64
  Status int64
}

var ticker = time.NewTicker(time.Second)

func StartWarden(db *sql.DB) {
  warden = &Warden{
    Notify: make(chan []byte),
  }

  go warden.makeTheRound()

  func() {
    for _ = range ticker.C {
      gpio.GetCurrentStatus()
    }
  }()
}

func (w *Warden) makeTheRound() {
  for {
    data := <- w.Notify
    var pins []Pin
    json.Unmarshal(data, &pins)

    zones, err := w.getZones(pins)

    if err != nil {
      log.Fatal(err)
    }

    outdated := false
    for i := 0; i < len(zones); i++ {
      z := zones[i]
      if !z.HasActiveSchedule() && z.Status > 0 {
        gpio.Close(z.Gpio)
        outdated := true
      }
    }

    if !outdated {
      payload, err := json.Marshal(zones)
      if err == nil {
        notify(clients, string(payload))
      }
    }
  }
}

func (w *Warden) getZones(pins []Pin) ([]Zone, error) {
  var results []Zone
  query, err := db.Query("select zones.id, zones.gpio from zones;")
  if err != nil {
    return nil, err
  }

  defer query.Close()

  for query.Next() {
    var id int64
    var gpio int64
    query.Scan(&id, &gpio)
    zone := NewZone(id, gpio, pins)
    results = append(results, *zone)
  }

  return results, nil
}
