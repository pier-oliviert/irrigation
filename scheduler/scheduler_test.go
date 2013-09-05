package scheduler

import (
  "testing"
  "github.com/pothibo/irrigation/db"
  "github.com/pothibo/irrigation/models"
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
  models.RegisterEntry()
  models.RegisterValve()
  models.RegisterSchedule()
}


func TestSchedulerManageSchedules(t *testing.T) {
  connect(t)
	query := "select s.Id, s.ValveId, s.Active, s.Start, s.Interval, s.Length " +
		"from schedules s " +
		"where s.Active = 1 "

	instances, err := db.Orm().Select(models.Schedule{}, query)

	if err != nil {
		t.Fail()
		t.Log(err)
	}

	schedules := make([]*models.Schedule, len(instances))

	for i, instance := range instances {
		schedules[i] = instance.(*models.Schedule)
	}

	time.Sleep(20 * time.Second)
}
