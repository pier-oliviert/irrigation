package models

import (
	"github.com/coopernurse/gorp"
	"github.com/pothibo/irrigation/db"
	"log"
	"time"
)

type Entry struct {
	Id      int32
	Length  int64
	ValveId int32
}

func (e *Entry) PreInsert(s gorp.SqlExecutor) error {
	valve, err := e.Valve()
	if err != nil {
		return err
	}
	valve.Open()
	go func(e *Entry) {
		valve, err := e.Valve()
		if err != nil {
			log.Fatalln(err)
			return
		}
		time.Sleep(time.Duration(e.Length) * time.Second)
		valve.Close()
	}(e)
	return nil
}

func (e *Entry) Valve() (valve *Valve, err error) {
	return GetValveById(e.ValveId)
}

func RegisterEntry() {
	db.Orm().AddTableWithName(Entry{}, "entries").SetKeys(true, "Id")
}
