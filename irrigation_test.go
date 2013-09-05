package main

import (
	"testing"
  "github.com/pothibo/irrigation/scheduler"
  "github.com/pothibo/irrigation/config"
	"os/user"
  "fmt"
)

func TestServerStarting(t *testing.T) {
  usr, err := user.Current()
  if err != nil {
    t.Fail()
    t.Log(fmt.Sprintf(`There was a problem retrieving the current user.
      Error: %v`, err))
  }

  cfg = config.Init(fmt.Sprintf("%v/irrigation", usr.HomeDir))
  cfg.Port = 7777
  dbTestName := fmt.Sprintf("%vtest", *cfg.Database["name"])
  cfg.Database["name"] = &dbTestName
  initDB(cfg, true)
  scheduler.Run()
  if !scheduler.IsRunning() {
    t.Fail()
    t.Log("Scheduler was not initialized while starting the server")
  }
}

