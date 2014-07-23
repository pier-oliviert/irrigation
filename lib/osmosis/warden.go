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
  Pins chan []Pin
  Updates chan []byte
}

func StartWarden(db *sql.DB) {
  warden = &Warden{
    Pins: make(chan []Pin),
    Updates: make(chan []byte),
  }

  go warden.notify()
  go warden.makeTheRound()

  for _ = range time.Tick(time.Second) {
    log.Print("Ticking")
    gpio.Send(&Command{Name: "list"})
  }
}

func (w *Warden) notify() {
  for {
    update := <- w.Updates
    fn := func(cs []*Client) {
      for i := 0; i < len(cs); i++ {
        client := cs[i]
        client.Conn.Write(update)
      }
    }
    ExecuteOnClients(fn)
  }
}

func (w *Warden) makeTheRound() {
  for {
    pins := <- w.Pins
    zones, err := w.getZones(pins)

    if err != nil {
      log.Fatal(err)
    }

    outdated := false
    for i := 0; i < len(zones); i++ {
      z := zones[i]
      if !z.HasActiveSchedule() && z.State > 0 {
        cmd := &Command{
          Name: "close",
          Id: z.Gpio,
        }
        gpio.Send(cmd)
        outdated = true
        break
      }
    }

    if outdated == false {
      payload, err := json.Marshal(zones)
      if err == nil {
        w.Updates <- payload
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
