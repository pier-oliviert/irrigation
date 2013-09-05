package models

import (
	"testing"
  "github.com/pothibo/irrigation/db"
  "github.com/pothibo/irrigation/config"
  "time"
  "os/user"
  "fmt"
)

func connect(t *testing.T) {
  usr, err := user.Current()
  if err != nil {
    t.Fail()
    t.Log(fmt.Sprintf(`There was a problem retrieving the current user.
      Error: %v`, err))
  }

  cfg := config.Init(fmt.Sprintf("%v/irrigation", usr.HomeDir))
  cfg.Port = 7777
  dbTestName := fmt.Sprintf("%vtest", *cfg.Database["name"])
  cfg.Database["name"] = &dbTestName
  initDB(cfg)
}

func initDB(cfg *config.Config) {
  db.ConfigureWith(cfg)
  db.Init()
  RegisterEntry()
  RegisterValve()
  RegisterSchedule()
}

func TestSetIntervalWithSchedule(t *testing.T) {
  connect(t)
	schedule := &Schedule{}
	err := schedule.SetInterval("5", "second")
	if err != nil {
		t.Fail()
		t.Log(err)
	}
	if schedule.Interval != 5 {
		t.Fail()
		t.Log("Schedule Interval should be set to 5.")
	}

	err = schedule.SetInterval("5", "minute")
	if err != nil {
		t.Fail()
		t.Log(err)
	}

	if schedule.Interval != 300 {
		t.Fail()
		t.Log("Schedule Interval should be set to 300.")
	}
}

func TestValveCreation(t *testing.T) {
  connect(t)
	valve := FirstValveOrCreate(17)
	if valve.RelayId != 17 {
		t.Fail()
	}
}

func TestValvesAreUnique(t *testing.T) {
  connect(t)
	query := "select v.RelayId, v.Name " +
		"from valves v " +
		"where v.RelayId = 17"
	valves, err := db.Orm().Select(Valve{}, query)

	if err != nil || len(valves) != 1 {
		t.Fail()
		t.Log(valves)
		t.Log(err)
	}

}

func TestSchedulesCreation(t *testing.T) {
  connect(t)
	valve, _ := GetValveByRelayId(17)
	schedule := &Schedule{
		ValveId:  valve.Id,
		Length:   3,
		Interval: 2,
		Active:   true,
	}

	schedule.SetStart(time.Now().Add(time.Second * 5).String())

	db.Orm().Insert(schedule)
}

