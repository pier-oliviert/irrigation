package main

import (
	"database/sql"
	"encoding/json"
	"log"
	"sync"
	"time"
)

// All those fields should be private and only
// be accessible via Get/Set to make use of mutexes.
type Zone struct {
	Id    int64
	Gpio  int64
	State int64

	json.Marshaler
	mutex sync.RWMutex
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

func (z *Zone) GetState() int64 {
	z.mutex.Lock()
	defer z.mutex.Unlock()
	return z.State
}

func (z *Zone) Opened() bool {
	return z.GetState() == 0
}

func ZonesAsJSON(zones []Zone) ([]byte, error) {
	return json.Marshal(zones)
}

func (z *Zone) ActiveSchedules() *sql.Rows {
	rows, err := db.Query(`select to_char(sprinkles.ends_at, 'DD Mon IY HH24:MI:SS')
		from sprinkles
		inner join zones
			on (sprinkles.zone_id = zones.id)
		where sprinkles.ends_at > CAST(NOW() at time zone 'utc' as timestamp);`)

	if err != nil {
		log.Print(err)
	}

	return rows
}

func (z *Zone) ClosingTime() time.Time {
	var closingTime time.Time
	if rows := z.ActiveSchedules(); rows != nil {
		defer rows.Close()

		for rows.Next() {
			var timeStr string
			rows.Scan(&timeStr)
			date, _ := time.Parse("02 Jan 06 15:04:05", timeStr)
			if closingTime.Before(date) {
				closingTime = date
			}
		}
	}

	return closingTime

}

func (z *Zone) MarshalJSON() ([]byte, error) {
	obj := struct {
		Id    		int `json:"id"`
		CloseAt   time.Time `json:"close_at"`
		Status		string `json:"status"`
	}{
		Id: int(z.Id),
		Status: "close",
	}

	if date := z.ClosingTime(); !date.IsZero() {
		obj.CloseAt = date
	}

	if z.Opened() {
		obj.Status = "open"
	}

	return json.Marshal(obj)
}
