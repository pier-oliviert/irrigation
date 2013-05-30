package main

import (
	"flag"
	"github.com/globocom/config"
	"github.com/gorilla/pat"
	"github.com/pothibo/irrigation/db"
	"github.com/pothibo/irrigation/models"
	"github.com/pothibo/irrigation/scheduler"
	"log"
	"net/http"
	"os"
	"strconv"
	"strings"
)

var Path string

func main() {
	Path = os.Getenv("GOPATH")
	db.Init("production")
	models.RegisterEntry()
	models.RegisterValve()
	models.RegisterSchedule()
	flag.Bool("server", true, "Start the server")
	flag.Bool("initdb", true, "Initialize the database.")
	flag.Bool("activate", true, "Activate the relays.")

	flag.Parse()
	flag.Visit(actionFlag)

	flag.PrintDefaults()
}

func actionFlag(flag *flag.Flag) {
	switch {
	case flag.Name == "server":
      err := launchServer()
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
	}
}

func launchServer() error {
	configPath := []string{Path, "/assets/irrigation/config", ".yml"}
	config.ReadConfigFile(strings.Join(configPath, ""))

	scheduler.Run()

	http.Handle("/assets/", http.StripPrefix("/assets/", http.FileServer(http.Dir(Path+"/assets/irrigation/assets"))))

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

  err := http.ListenAndServe(":7777", nil)
  
  return err

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
			log.Panicf("Valve %s could not be configured. Ignoring")
		}
		relays[valve] = models.FirstValveOrCreate(valve)
	}

	return relays
}
