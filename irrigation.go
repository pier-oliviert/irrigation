package main

import (
	"flag"
	"github.com/pothibo/irrigation/config"
	"github.com/gorilla/pat"
	"github.com/pothibo/irrigation/db"
	"github.com/pothibo/irrigation/gpio"
	"github.com/pothibo/irrigation/models"
	"github.com/pothibo/irrigation/scheduler"
	"log"
	"net/http"
	"os"
	"os/user"
  "fmt"
)

var path string
var cfg *config.Config

func init() {
  usr, err := user.Current()
  if err != nil {
    log.Fatalln(fmt.Sprintf(`There was a problem retrieving the current user.
      Error: %v`, err))
  }

	flag.Bool("server", true, "Start the server with port: 7777")
	flag.Bool("initialize", true, "Initialize the server. Doing this will reset everything to defaults, database included.")
	flag.Bool("activate", true, "Activate the relays.")
	flag.Bool("help", true, "Show this help menu.")
	flag.StringVar(&path, "path", fmt.Sprintf("%v/irrigation",usr.HomeDir), "Path where your config & HTML files will be located.")
}

func main() {
	flag.Parse()
	flag.Visit(actionFlag)

	flag.PrintDefaults()
}

func actionFlag(flag *flag.Flag) {
	switch {
	case flag.Name == "server":
    cfg = config.Init(path)
    initDB(cfg, false)
		err := launchServer()
		if err != nil {
			log.Panicln(err)
			os.Exit(1)
		}
	case flag.Name == "initialize":
    initialize()
	case flag.Name == "activate":
    cfg = config.Init(path)
    initDB(cfg, false)
		err := activateRelay()
		if err != nil {
			log.Fatalln(err)
		}
		os.Exit(1)
	}
}

func initialize() {
  var mysqlRootPassword string
  fetchRepository(path)
  cfg = config.Init(path)

  database := cfg.Database

  msg := fmt.Sprintf("User MySQL? (default: %v)", *database["user"])
  config.AskForValue(cfg.Database["user"], msg)

  msg = fmt.Sprintf("User password for MySQL? (default: %v)", *database["password"])
  config.AskForValue(cfg.Database["password"], msg)

  msg = fmt.Sprintf("MySQL Database name? (default: %v)", *database["name"])
  config.AskForValue(cfg.Database["name"], msg)

  fmt.Println("Your root password for MySQL:")
  fmt.Scanln(&mysqlRootPassword)

  cfg.Update()

  db.InitializeDatabase(cfg, mysqlRootPassword)

  initDB(cfg, true)

  fmt.Println(fmt.Sprintf(`Configuration of Irrigation is finished!
  You can now start the server
  $ irrigation -server
  If you want to modify your current configuration, you can do so by modifying
  %v/config.yml
  Enjoy!
  @pothibo`, path))
  os.Exit(1)

}

func initDB(cfg *config.Config, create bool) {
  db.ConfigureWith(cfg)
  db.Init()
  models.RegisterEntry()
  models.RegisterValve()
  models.RegisterSchedule()

  if create {
    err := db.Orm().CreateTablesIfNotExists()
    if err != nil {
      log.Fatalln(err)
    }
  }
}

func launchServer() error {
	scheduler.Run()

  stylesheets := fmt.Sprintf("%v/stylesheets/", path)
  javascripts := fmt.Sprintf("%v/javascripts/", path)
	http.Handle("/stylesheets/", http.StripPrefix("/stylesheets/", http.FileServer(http.Dir(stylesheets))))
	http.Handle("/javascripts/", http.StripPrefix("/javascripts/", http.FileServer(http.Dir(javascripts))))

	r := pat.New()

	r.Post("/schedules/{scheduleId}", updateSchedule)
	r.Get("/schedules/{scheduleId}/edit", editSchedule)
	r.Get("/schedules/new", newSchedule)
	r.Post("/schedules", createSchedule)

	r.Get("/valves/{valveId}/edit", editValve)
	r.Get("/valves/{valveId}/open", openValve)
	r.Get("/valves/{valveId}/close", closeValve)
	r.Post("/valves/{valveId}", updateValve)
	r.Get("/valves/{valveId}", showValve)

	r.Get("/manual", manual)
	r.Get("/", homepage)

	http.Handle("/", r)

	initializeTemplates(path)
  err := http.ListenAndServe(fmt.Sprintf(":%d", cfg.Port), nil)
	return err
}

func activateRelay() error {
	for _, valve := range Valves() {
		err := gpio.Activate(valve.RelayId)
		if err != nil {
			return err
		}
	}
	return nil
}

func Valves() map[uint8]*models.Valve {
	relays := make(map[uint8]*models.Valve)


	for _, valve := range cfg.Valves {
		relays[valve] = models.FirstValveOrCreate(valve)
	}

	return relays
}
