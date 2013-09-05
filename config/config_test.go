package config

import (
	"github.com/globocom/config"
  "testing"
  "strings"
)

func TestSetADefaultPort(t *testing.T) {
  cfg := &Config{}
  cfg.port()

  if cfg.Port != 7777 {
    t.Fail()
    t.Log("Should set a default port to 7777")
  }
}

func TestSetADifferentPort(t *testing.T) {
  cfg := &Config{}
  config.Set("port", 8888)

  cfg.port()

  if cfg.Port != 8888 {
    t.Fail()
    t.Log("Port should be set to 8888")
  }
}

func TestDatabaseParsing(t *testing.T) {
  cfg := &Config{}
  config.Set("database:name", "test")
  config.Set("database:user", "testuser")
  config.Set("database:password", "password12345")

  cfg.database()

  if !strings.EqualFold(*cfg.Database["name"], "test") {
    t.Fail()
    t.Log("Database name should be 'test'")
  }

  if !strings.EqualFold(*cfg.Database["user"], "testuser") {
    t.Fail()
    t.Log("Database name should be 'testuser'")
  }

  if !strings.EqualFold(*cfg.Database["password"], "password12345") {
    t.Fail()
    t.Log("Database name should be 'password12345'")
  }
}
