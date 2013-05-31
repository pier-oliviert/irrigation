package db

import (
	"database/sql"
	"github.com/coopernurse/gorp"
	_ "github.com/mattn/go-sqlite3"
)

var orm *gorp.DbMap

func Orm() *gorp.DbMap {
	return orm
}

func Init(path string) error {

	database, err := sql.Open("sqlite3", path)
	if err != nil {
      return err
	}

	orm = &gorp.DbMap{
		Db:      database,
		Dialect: gorp.SqliteDialect{},
	}
  return nil

}

func Create() error {
	err := Orm().CreateTables()
	return err
}
