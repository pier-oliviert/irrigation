package db

import (
	"github.com/pothibo/irrigation/config"
	"database/sql"
	"github.com/coopernurse/gorp"
	_ "github.com/go-sql-driver/mysql"

  "fmt"
  "log"
  "os/exec"
)

var orm *gorp.DbMap
var cfg *config.Config

func Orm() *gorp.DbMap {
	return orm
}

func ConfigureWith(c *config.Config) {
  cfg = c
}

func Init() error {
  path := fmt.Sprintf("%v:%v@/%v", *cfg.Database["user"], *cfg.Database["password"], *cfg.Database["name"])
	database, err := sql.Open("mysql", path)
	if err != nil {
		return err
	}

	orm = &gorp.DbMap{
		Db:      database,
		Dialect: gorp.MySQLDialect{"InnoDB", "UTF8"},
	}

	return nil
}

func InitializeDatabase(config *config.Config, rootPassword string) {
  name := *config.Database["name"]
  user := *config.Database["user"]
  password := *config.Database["password"]

  createDatabase(rootPassword, name)
  createUser(rootPassword, user, password, name)
}

func createDatabase(rootPassword string, name string) {
  var access string
  if len(rootPassword) > 0 {
    access = fmt.Sprintf("-p%v", rootPassword)
  }
  sql := fmt.Sprintf(`-e 
  drop database if exists %v; create database %v;
  drop database if exists %vtest; create database %vtest;
  `, name, name, name, name)

  cmd := exec.Command("mysql", `-uroot`, access, sql)
  out, err := cmd.CombinedOutput()

  if err != nil {
    log.Printf("%s",out)
    log.Fatalln(fmt.Sprintf(`An error occurred trying to create the database %v
    %v`, name, err))
  }
}

func createUser(rootPassword string, user string, password string, name string) {
  var access string
  if len(rootPassword) > 0 {
    access = fmt.Sprintf("-p%v", rootPassword)
  }
  sql := fmt.Sprintf(`-e
  grant all privileges on %v.* to '%v'@'localhost' identified by '%v';
  grant all privileges on %vtest.* to '%v'@'localhost' identified by '%v';
  `,name, user, password, name, user, password)

  cmd := exec.Command("mysql", "-uroot", access, sql)
  out, err := cmd.CombinedOutput()

  if err != nil {
    log.Printf("%s", out)
    log.Fatalln(fmt.Sprintf(`An error occurred trying to create the user %v
    %v`, user, err))
  }
}
