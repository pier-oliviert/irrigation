package main

import (
  "net/http"
    "strconv"
    "html/template"
    "log"
    "irrigation/db"
    "irrigation/models"
)

var indexTmpl = template.Must(template.ParseFiles(
    "views/_base.html",
    "views/home.html",
))

var valvesTmpl = template.Must(template.ParseFiles(
    "views/_base.html",
    "views/valves/show.html",
))

var editValveTmpl = template.Must(template.ParseFiles(
    "views/_base.html",
    "views/valves/edit.html",
))

var editScheduleTmpl = template.Must(template.ParseFiles(
    "views/_base.html",
    "views/schedules/edit.html",
))

func homepage(w http.ResponseWriter, r *http.Request)  {
    err := indexTmpl.Execute(w, map[string]interface{} {
        "Relays": Valves(),
    })

    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
    }
}

func createSchedule(w http.ResponseWriter, r *http.Request) {
    if r.Method != "POST" {
        //Set a flash and redirect
        http.Redirect(w, r, "/", 302)
    }
    r.ParseForm()
    valveId, err := strconv.ParseInt(r.PostFormValue("valve"), 10, 32)
    if err != nil {
        log.Fatalln(err)
    }
    valve, err := models.GetValveById(int32(valveId))
    
    if err != nil {
      log.Println(err)
      http.Redirect(w, r, "/", 302)
    }

    schedule := &models.Schedule{
        ValveId: valve.Id,
        Active: true,
    }

    err = schedule.SetInterval(
        r.PostFormValue("interval[multiplicator]"),
        r.PostFormValue("interval[denominator]"))

    if err != nil {
        log.Fatal(err)
    }

    err = schedule.SetLength(
        r.PostFormValue("length[multiplicator]"),
        r.PostFormValue("length[denominator]"))

    if err != nil {
        log.Fatal(err)
    }

    err = schedule.SetStart(r.PostFormValue("date"))
    if err != nil {
        log.Println(err)
    }

    err = db.Orm().Insert(schedule)
    if err != nil {
        log.Println(err)
    }
    http.Redirect(w, r, "/", 302)
}

func updateSchedule(w http.ResponseWriter, r *http.Request) {
  id, err := strconv.Atoi(r.URL.Query().Get(":scheduleId"))
  if err != nil {
    log.Println(err)
  }

  schedule, err := models.GetScheduleById(int32(id))

  err = schedule.SetInterval(
      r.PostFormValue("interval[multiplicator]"),
      r.PostFormValue("interval[denominator]"))

  if err != nil {
      log.Println(err)
      return
  }

  err = schedule.SetLength(
      r.PostFormValue("length[multiplicator]"),
      r.PostFormValue("length[denominator]"))

  if err != nil {
      log.Println(err)
      return
  }

  err = schedule.SetStart(r.PostFormValue("date"))
  if err != nil {
      log.Println(err)
  }

  if err != nil {
    log.Println(err)
  } else {
    db.Orm().Update(schedule)
  }
}

func editSchedule(w http.ResponseWriter, r *http.Request) {
  id, err := strconv.Atoi(r.URL.Query().Get(":scheduleId"))
  if err != nil {
    log.Println(err)
  }

  schedule, err := models.GetScheduleById(int32(id))

  err = editScheduleTmpl.Execute(w, schedule)

  if err != nil {
    log.Println(err)
  }
}

func showValve(w http.ResponseWriter, r *http.Request) {
  valveId, _ := strconv.Atoi(r.URL.Path[len("/valves/"):])
  valve, err := models.GetValveById(int32(valveId))
  
  if err != nil {
    log.Println(err)
    http.Redirect(w, r, "/", 302)
  }

  schedules, err := models.GetSchedulesForValve(valve)

  if err != nil {
    log.Println(err)
    http.Redirect(w, r, "/", 302)
  }

  err = valvesTmpl.Execute(w, map [string]interface{} {
    "Schedules": schedules,
    "Valve": valve,
  })
}

func editValve(w http.ResponseWriter, r *http.Request) {
  id, err := strconv.Atoi(r.URL.Query().Get(":valveId"))
  if err != nil {
    log.Println(err)
    http.Redirect(w, r, "/", 302)
  }

  valve, err := models.GetValveById(int32(id))
  err = editValveTmpl.Execute(w, valve)
  if err != nil {
    log.Println(err)
  }
}

func updateValve(w http.ResponseWriter, r *http.Request) {
  id, err := strconv.Atoi(r.URL.Query().Get(":valveId"))
  if err != nil {
    log.Println(err)
    http.Redirect(w, r, "/", 302)
  }

  valve, err := models.GetValveById(int32(id))

  if err == nil {
    valve.Name = r.PostFormValue("name")
    db.Orm().Update(valve)
  }

  http.Redirect(w, r, "/", 302)
}

func openValve(w http.ResponseWriter, r *http.Request) {
    relay, _ := strconv.Atoi(r.URL.Path[len("/open/"):])

    valve, err := models.GetValveByRelayId(relay)

    if err == nil && valve != nil {
      valve.Open()

    }
    http.Redirect(w, r, "/", 302)
}

func closeValve(w http.ResponseWriter, r *http.Request) {
    relay, _ := strconv.Atoi(r.URL.Path[len("/close/"):])

    valve, err := models.GetValveByRelayId(relay)

    if err == nil && valve != nil {
      valve.Close()

    }
    http.Redirect(w, r, "/", 302)
}
