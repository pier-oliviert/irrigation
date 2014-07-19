package main

import (
  "time"
  _ "github.com/lib/pq"
  "database/sql"
  "strconv"
  "log"
  )

type Warden struct {
  gpios []int
}

var ticker = time.NewTicker(time.Second)

func NewWarden(db *sql.DB) *Warden {
  w := new(Warden)
  go func() {
    for _ = range ticker.C {
      err := w.list()

      if err != nil {
        log.Print(err)
      }
    }
  }()
  return w
}

func (w *Warden) list() error {
  gpios := make(map[int]*Zone)
  zones, err := db.Query("select zones.id, zones.gpio from zones;")
  if err != nil {
    return err
  }
  defer zones.Close()
  for zones.Next() {
    var id int
    var gpio int
    zones.Scan(&id, &gpio)
    gpios[gpio] = NewZone(id, gpio)
    sprinkles, err := db.Query("select zones.gpio from zones inner join sprinkles on (sprinkles.zone_id = zones.id) where sprinkles.ends_at > CAST(NOW() at time zone 'utc' as timestamp);")
    if err != nil {
      return err
    }

    defer sprinkles.Close()

    for sprinkles.Next() {
      sprinkles.Scan(&gpio)
      gpios[gpio].active = true
    }
  }

  for key, zone := range gpios {
    if zone.active == true {
      // Check if Pin is up, if it is. Close it.
      cmd, err := NewCommand([]string{"close", strconv.FormatInt(int64(key), 10)})
      if err != nil {
        return err
      }

      notify(clients, cmd)
    }
  }

  return nil
}
