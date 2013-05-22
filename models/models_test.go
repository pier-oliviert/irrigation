package models

import (
	"testing"
)

func TestSetIntervalWithSchedule(t *testing.T) {
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
