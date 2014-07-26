package main

import (
	"database/sql"
	"encoding/json"
	"log"
	"sync"
)

type Zone struct {
	Id    int64 `json:"id"`
	Gpio  int64 `json:"gpio"`
	State int64 `json:"state"`
	mutex sync.Mutex
}

type Pin struct {
	Id    int64 `json:"id"`
	State int64 `json:"state"`
}

func AllZones() map[int64]*Zone {
	results := make(map[int64]*Zone)
	query, err := db.Query("select zones.id, zones.gpio from zones;")
	if err != nil {
		log.Fatal(err)
	}

	defer query.Close()

	for query.Next() {
		var id int64
		var gpio int64
		query.Scan(&id, &gpio)
		results[id] = &Zone{Id: id, Gpio: gpio}
	}

	return results
}

func (z *Zone) SetState(s int64) bool {
	z.mutex.Lock()
	defer z.mutex.Unlock()
	if z.State != s {
		z.State = s
		return true
	}
	return false
}

func (z *Zone) SetOpen() {
}

func (z *Zone) SetClose() {
}

func ZonesAsJSON(zones []Zone) ([]byte, error) {
	return json.Marshal(zones)
}

func (z *Zone) ActiveSchedules(db *sql.DB) *sql.Rows {
	rows, err := db.Query(`select zones.gpio
		from zones
		inner join sprinkles
			on (sprinkles.zone_id = zones.id)
		where sprinkles.ends_at > CAST(NOW() at time zone 'utc' as timestamp);`)

	if err != nil {
		log.Print(err)
	}

	return rows
}
