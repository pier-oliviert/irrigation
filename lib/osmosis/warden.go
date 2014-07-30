package main

import (
	"database/sql"
	_ "github.com/lib/pq"
	"log"
	"time"
	"encoding/json"
)

type Warden struct {
	GPIO  *GPIO
	Zones *Zones
	Update chan *Zone
}

func NewWarden(db *sql.DB, g *GPIO) *Warden {
	warden = &Warden{
		Update: make(chan *Zone),
		Zones: NewZones(),
		GPIO: g,
	}

	warden.GPIO.StartListening(warden.Zones.Update)
	go warden.updateClients()
	go warden.Zones.States(warden.Update)
	go warden.monitor(db)
	return warden
}

func (w *Warden) updateClients() {
	for {
		zone := <- w.Update
		payload, err := json.Marshal(zone)
		if err != nil {
			log.Print(err)
		}

		mutex.Lock()
		for _, client := range clients {
			go func(c *Client, msg []byte) {
				ch := make(chan int, 1)
				select {
				case ch <- c.Write(msg):
				case <-time.After(5 * time.Second):
					RemoveClient(c)
				}
			}(client, payload)
		}
		mutex.Unlock()
	}
}

func (w *Warden) monitor(db *sql.DB) {
	for _ = range time.Tick(time.Second) {
		for _, zone := range w.Zones.All() {
			active := false
			query := zone.ActiveSchedules(db)
			if query == nil {
				continue
			}

			for query.Next() {
				active = true
				break
			}

			query.Close()

			if active == false && zone.Opened() {
				w.GPIO.Send(&Command{Name: "close", Id: zone.Gpio})
			}
		}
	}
}
