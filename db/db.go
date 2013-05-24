package db

import (
	"database/sql"
	"github.com/coopernurse/gorp"
	_ "github.com/mattn/go-sqlite3"
	"log"
	"os"
	"strings"
)

var orm *gorp.DbMap

func Orm() *gorp.DbMap {
	return orm
}

func Init(env string) {
	var dbPath []string
	gopath := os.Getenv("GOPATH")

	if env == "production" {
		dbPath = []string{gopath, "/assets/irrigation/", "irrigation", ".db"}
	} else {
		dbPath = []string{"irrigation", "_", env, ".db"}
	}

	database, err := sql.Open("sqlite3", strings.Join(dbPath, ""))
	if err != nil {
		log.Fatalln(err)
		return
	}

	orm = &gorp.DbMap{
		Db:      database,
		Dialect: gorp.SqliteDialect{},
	}

}

func Create() error {
	err := Orm().CreateTables()
	return err
}
