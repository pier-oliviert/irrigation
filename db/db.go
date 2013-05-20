package db

import (
  "database/sql"
  "github.com/coopernurse/gorp"
  _ "github.com/mattn/go-sqlite3"
  "log"
  "strings"

)

var orm *gorp.DbMap

func Orm() *gorp.DbMap {
  return orm;
}

func Init(env string) {
  var dbPath []string

  if env == "production" {
    dbPath = []string{"db/irrigation", ".db"}
  } else {
    dbPath = []string{"db/irrigation", "_", env, ".db"}
  }

  database, err := sql.Open("sqlite3", strings.Join(dbPath, ""))
  if err != nil {
    log.Fatalln(err)
    return
  }

  orm = &gorp.DbMap{
    Db: database,
    Dialect: gorp.SqliteDialect{},
  }

}

func Create() error {
  err := Orm().CreateTables()
  return err
}

