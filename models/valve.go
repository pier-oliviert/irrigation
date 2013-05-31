package models

import (
	"github.com/pothibo/irrigation/db"
	"github.com/pothibo/irrigation/gpio"
	"log"
	"strconv"
)

type Valve struct {
	Id      int32
	RelayId int
	Name    string
}

func (v *Valve) IsOpened() bool {
	opened, err := gpio.IsOpened(v.RelayId)
	if err != nil {
		log.Fatalln(err)
	}
	return opened
}

func (v *Valve) Title() string {
	if v.Name == "" {
		return "Relay #" + strconv.Itoa(int(v.Id))
	} else {
		return v.Name
	}
}

func (v *Valve) ActiveSchedules() int {
	count := 0
	for _, schedule := range v.Schedules() {
		if schedule.Active {
			count++
		}
	}
	return count
}

func (v *Valve) Schedules() []*Schedule {
	schedules, err := GetSchedulesForValve(v)
	if err != nil {
		log.Println(err)
	}
	return schedules
}

func (v *Valve) Open() {
	gpio.Open(v.RelayId)
}

func (v *Valve) Close() {
	gpio.Close(v.RelayId)
}

func FirstValveOrCreate(relay int) *Valve {
	var valve *Valve
	query := "select v.Id, v.RelayId, v.Name " +
		"from valves v " +
		"where v.RelayId = ?"

	instances, err := db.Orm().Select(Valve{}, query, relay)

	if err != nil {
		log.Fatalln(err)
	}

	if len(instances) == 0 {
		valve = &Valve{
			RelayId: relay,
		}
		db.Orm().Insert(valve)
	} else {
		valve = instances[0].(*Valve)
	}

	return valve
}

func GetValveByRelayId(id int) (valve *Valve, err error) {
	query := "select v.Id, v.RelayId, v.Name " +
		"from valves v " +
		"where v.RelayId = ?"
	instances, err := db.Orm().Select(Valve{}, query, id)

	if err != nil {
		return nil, err
	}

	if len(instances) == 0 {
		return nil, nil
	} else {
		return instances[0].(*Valve), nil
	}
}

func GetValveById(id int32) (valve *Valve, err error) {
	query := "select v.Id, v.RelayId, v.Name " +
		"from valves v " +
		"where v.Id = ?"

	instances, err := db.Orm().Select(Valve{}, query, id)

	if err != nil {
		return nil, err
	}

	if len(instances) == 0 {
		return nil, err
	} else {
		return instances[0].(*Valve), nil
	}
}

func CloseAllValves() (valves []Valve, err error) {
	return nil, nil
}

func RegisterValve() {
	db.Orm().AddTableWithName(Valve{}, "valves").SetKeys(true, "Id")
}
