package db

import (
	"database/sql"
	"github.com/coopernurse/gorp"
	_ "github.com/go-sql-driver/mysql"
)

var orm *gorp.DbMap

func Orm() *gorp.DbMap {
	return orm
}

func Init(path string) error {

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

func Create() error {
	err := Orm().CreateTablesIfNotExists()
	return err
}
