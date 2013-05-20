package models

import (
  "strconv"
  "errors"
  "irrigation/db"
  "time"
)

type Schedule struct {
  Id int32
  ValveId int32
  Active bool
  Start int64
  Interval int64
  Length int64
}

func (s *Schedule) Valve() (valve *Valve, err error) {
  return GetValveById(s.ValveId)
}

func (s *Schedule) DateTimeInputValue() string {
  const layout = "2006-01-02T15:04:05"
  return time.Unix(s.Start, 0).Format(layout)
}

func (s *Schedule) SetInterval(multiplicator string, dur string) error {
  value, err := strconv.ParseInt(multiplicator, 10, 32)
  if err != nil {
    return err
  }

  duration, err := durationFromString(dur)
  if err != nil {
    return err
  }

  s.Interval = value * duration
  return nil
}

func (s *Schedule) SetLength(multiplicator string, dur string) error {
  value, err := strconv.ParseInt(multiplicator, 10, 32)
  if err != nil {
    return err
  }

  duration, err := durationFromString(dur)
  if err != nil {
    return err
  }

  s.Length = value * duration
  return nil
}

func (s *Schedule) SetStart(date string) error {
  const form = "2006-01-02T15:04"
  loc, err := time.LoadLocation("America/Montreal")

  if err != nil {
    return err
  }
  
  parsed, err := time.ParseInLocation(form, date, loc)

  if err != nil {
    return err
  }

  s.Start = parsed.UTC().Truncate(time.Second).Unix()
  return nil
  
}

func RegisterSchedule() {
  db.Orm().AddTableWithName(Schedule{}, "schedules").SetKeys(true, "Id")
}

func GetScheduleById(id int32) (schedule *Schedule, err error) {
  query := "select * " +
    "from schedules s " +
    "where s.Id = ?"

  instances, err := db.Orm().Select(Schedule{}, query, id)

  if err != nil {
    return nil, err
  }

  schedules := make([]*Schedule, len(instances))
  for i := range instances { schedules[i] = instances[i].(*Schedule) }

  if len(schedules) > 0 {
    return schedules[0], err
  } else {
    return nil, err
  }
}

func GetSchedulesForValve(valve *Valve) (schedules []*Schedule, err error) {
  query := "select * " +
    "from schedules s " +
    "where s.ValveId = ?"
  instances, err := db.Orm().Select(Schedule{}, query, valve.Id)

  if err != nil {
    return nil, err
  }

  schedules = make([]*Schedule, len(instances))
  for i := range instances { schedules[i] = instances[i].(*Schedule) }

  return schedules, err
}

func durationFromString(duration string) (seconds int64, err error) {
  switch duration {
  case "second":
    return 1, nil
  case "minute":
    return 60, nil
  case "hour":
    return 60 * 60, nil
  case "day":
    return 60 * 60 * 60, nil
  case "week":
    return 60 * 60 * 60 * 60, nil
  case "month":
    return 60 * 60 * 60 * 60  * 60, nil
  }

  return 0, errors.New("models: Duration needs to be one of the following: second, minute, hour, day, week, month")
}
