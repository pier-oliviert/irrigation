package main

import (
  "time"
  _ "github.com/lib/pq"
  "database/sql"
  )

type Warden struct {
  gpios []int
}

var ticker = time.NewTicker(time.Second)

func NewWarden(db *sql.DB) *Warden {
  w := new(Warden)
  go func() {
    for _ = range ticker.C {
      w.list()
    }
  }()
  return w
}

func (w *Warden) list() {
}
