package scheduler

import (
	"github.com/pothibo/irrigation/db"
	"github.com/pothibo/irrigation/models"
	"log"
	"time"
)

var working = false

func Run() {
	working = true
	go func() {
		for working == true {
			time.Sleep(time.Second)
			LaunchScheduledEntries(scheduledToLaunch())
		}
	}()
	return
}

func Stop() {
	working = false
	models.CloseAllValves()
}

func IsRunning() bool {
	return !!working
}

func LaunchScheduledEntries(list []*models.Schedule) {
	for i := 0; i < len(list); i++ {
		schedule := list[i]
		valve, err := schedule.Valve()

		if err != nil {
			log.Fatalln(err)
		}

		entry := &models.Entry{
			Length:  schedule.Length,
			ValveId: valve.Id,
		}

		err = db.Orm().Insert(entry)

		if err != nil {
			log.Fatalln(err)
		}
	}
}

func scheduledToLaunch() []*models.Schedule {
	active_at := time.Now().Truncate(time.Second).Unix()
	query := "select *" +
		"from schedules s " +
		"where s.Active = 1 and " +
		"( s.Start = ? OR ( (unix_timestamp() - from_unixtime(s.Start)) % s.Interval) = 0)"
	instances, err := db.Orm().Select(models.Schedule{}, query, active_at)

	if err != nil {
		log.Fatalln(err)
	}

	schedules := make([]*models.Schedule, len(instances))

	for i, instance := range instances {
		schedules[i] = instance.(*models.Schedule)
	}

	return schedules
}
