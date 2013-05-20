package main

import (
  "os"
  "testing"
  "time"
  _ "github.com/astaxie/beedb"
  _ "github.com/mattn/go-sqlite3"

  "irrigation/db"
  "irrigation/models"
  "irrigation/scheduler"
)

func TestStart(t *testing.T) {
  db.Init("test")
  models.RegisterEntry()
  models.RegisterValve()
  models.RegisterSchedule()

  err := db.Create()
  if err != nil {
    t.Fail()
    t.Log(err)
  }

}

func TestValveCreation(t *testing.T) {
  valve := models.FirstValveOrCreate(17)
  if valve.RelayId != 17 {
    t.Fail()
  }
}

func TestValvesAreUnique(t *testing.T) {
  query := "select v.RelayId, v.Name " +
    "from valves v " +
    "where v.RelayId = 17"
  valves, err := db.Orm().Select(models.Valve{}, query)

  if err != nil || len(valves) != 1 {
    t.Fail()
    t.Log(valves)
    t.Log(err)
  }

}

func TestSchedulesCreation(t *testing.T) {
  valve, _ := models.GetValveByRelayId(17)
  schedule := &models.Schedule{
    ValveId: valve.Id,
    Length: 3,
    Interval: 2,
    Active: true,
  }

  schedule.SetStart(time.Now().Add(time.Second * 5).String())

  db.Orm().Insert(schedule)
}

func TestSchedulerIsRunning(t *testing.T) {
  scheduler.Run()
  if !scheduler.IsRunning() {
    t.Fail()
    t.Log("Scheduler is not running")
  }
}

func TestSchedulerManageSchedules(t *testing.T) {
  
    query := "select s.Id, s.ValveId, s.Active, s.Start, s.Interval, s.Length " +
      "from schedules s " +
      "where s.Active = 1 "

    instances, err := db.Orm().Select(models.Schedule{}, query)

    if err != nil {
      t.Fail()
      t.Log(err)
    }

    schedules := make([]*models.Schedule, len(instances))

    for i, instance := range instances { schedules[i] = instance.(*models.Schedule) }

    time.Sleep(20 * time.Second)
}

func TestEnded(t *testing.T) {
  os.Remove("irrigation-test.db")
}
