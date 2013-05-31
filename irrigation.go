package main

import (
    "errors"
	"flag"
	"github.com/globocom/config"
	"github.com/gorilla/pat"
	"github.com/pothibo/irrigation/gpio"
	"github.com/pothibo/irrigation/db"
	"github.com/pothibo/irrigation/models"
	"github.com/pothibo/irrigation/scheduler"
	"log"
	"net/http"
	"os"
	"strconv"
)
var configFile string

func init() {
	flag.Bool("server", true, "Start the server")
	flag.Bool("initdb", true, "Initialize the database.")
	flag.Bool("activate", true, "Activate the relays.")
	flag.Bool("help", true, "Show this help menu.")
  flag.StringVar(&configFile, "c", "/srv/http/irrigation/config.yml", "Configuration file")
}

func main() {
    config.ReadConfigFile(configFile)
    path, err := config.GetString("database")
    if err != nil {
        log.Panicln("No database file specified in config file.");
        os.Exit(1)
    }
	db.Init(path)
	models.RegisterEntry()
	models.RegisterValve()
	models.RegisterSchedule()

	flag.Parse()
	flag.Visit(actionFlag)

	flag.PrintDefaults()
}

func actionFlag(flag *flag.Flag) {
	switch {
	case flag.Name == "server":
      err := launchServer(flag)
      if err != nil {
          log.Panicln(err)
          os.Exit(1)
      }
	case flag.Name == "initdb":
		err := db.Create()
		if err != nil {
			log.Panicln(err)
		}
		os.Exit(1)

  case flag.Name == "activate":
      err := activateRelay()
      if err != nil {
          log.Fatalln(err)
      }
      os.Exit(1)
	}
}

func launchServer(flag *flag.Flag) error {
    path, err := config.GetString("assets")
    if err != nil {
        return errors.New("No assets folder specified in config file.")
    }

	scheduler.Run()

	http.Handle("/assets/", http.StripPrefix("/assets/", http.FileServer(http.Dir(path))))

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

  err = http.ListenAndServe(":7777", nil)
  
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

func Valves() map[int]*models.Valve {
	relays := make(map[int]*models.Valve)

	valves, err := config.GetList("valves")
	if err != nil {
		log.Fatalf("Could not load the valves id: %v", err)
	}

	for _, value := range valves {
		valve, err := strconv.Atoi(value)
		if err != nil {
			log.Printf("Valve %s could not be configured. Ignoring")
		}
		relays[valve] = models.FirstValveOrCreate(valve)
	}

	return relays
}
