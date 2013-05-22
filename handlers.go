package main

import (
  "net/http"
  "strconv"
  "html/template"
  "log"
  "irrigation/db"
  "irrigation/models"
  "irrigation/helpers"
)

var templates = map[string]*template.Template{}

func init() {
  templates["base"] = template.Must(template.ParseFiles(
    "views/_base.html",
  ))

  templates["error"] = template.Must(
    template.Must(templates["base"].Clone()).ParseFiles(
      "views/error.html",
    ))

  templates["home"] = template.Must(
    template.Must(templates["base"].Clone()).ParseFiles(
      "views/home.html",
    ))

  templates["showValve"] = template.Must(
    template.Must(templates["base"].Clone()).ParseFiles(
      "views/valves/show.html",
    ))

  templates["editValve"] = template.Must(
    template.Must(templates["base"].Clone()).ParseFiles(
      "views/valves/edit.html",
    ))

  templates["editSchedule"] = template.Must(
    template.Must(templates["base"].Clone()).ParseFiles(
      "views/schedules/edit.html",
    ))
}

func homepage(w http.ResponseWriter, r *http.Request)  {
    err := templates["home"].Execute(w, map[string]interface{} {
        "Relays": Valves(),
    })

    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
    }
    return
}

func createSchedule(w http.ResponseWriter, r *http.Request) {
  var schedule *models.Schedule
  var valve *models.Valve
  var err error

  valve, err = models.GetValveById(helpers.Int32ValueFrom(r.PostFormValue("valve"), -1))

  if err != nil {
    goto Error
  }

  schedule = &models.Schedule{
    ValveId: valve.Id,
    Active: true,
  }

  err = schedule.SetInterval(
    r.PostFormValue("interval[multiplicator]"),
    r.PostFormValue("interval[denominator]"))

  if err != nil {
    goto Error
  }

  err = schedule.SetLength(
    r.PostFormValue("length[multiplicator]"),
    r.PostFormValue("length[denominator]"))

  if err != nil {
    goto Error
  }

  err = schedule.SetStart(r.PostFormValue("date"))
  if err != nil {
    goto Error
  }

  err = db.Orm().Insert(schedule)
  if err != nil {
    goto Error
  }

  http.Redirect(w, r, "/", 302)
  return

  Error:
  tmpl_err := templates["error"].Execute(w, err)
  if tmpl_err != nil {
    log.Panicln(err)
  }

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
  } 

  db.Orm().Update(schedule)
  return
}

func editSchedule(w http.ResponseWriter, r *http.Request) {
  id, err := strconv.Atoi(r.URL.Query().Get(":scheduleId"))
  if err != nil {
    log.Println(err)
  }

  schedule, err := models.GetScheduleById(int32(id))

  err = templates["editSchedule"].Execute(w, schedule)

  if err != nil {
    log.Println(err)
  }
  return
}

func showValve(w http.ResponseWriter, r *http.Request) {
  valve, err := models.GetValveById(helpers.Int32ValueFrom(r.URL.Query().Get(":valveId"), -1))
  
  if err != nil {
    log.Println(err)
    http.Redirect(w, r, "/", 302)
    return
  }

  if valve == nil {
    http.Redirect(w, r, "/", 302)
    return
  }

  schedules, err := models.GetSchedulesForValve(valve)

  if err != nil {
    log.Println(err)
    http.Redirect(w, r, "/", 302)
  }

  err = templates["showValve"].Execute(w, map [string]interface{} {
    "Schedules": schedules,
    "Valve": valve,
  })
  return
}

func editValve(w http.ResponseWriter, r *http.Request) {
  id, err := strconv.Atoi(r.URL.Query().Get(":valveId"))
  if err != nil {
    log.Println(err)
    http.Redirect(w, r, "/", 302)
  }

  valve, err := models.GetValveById(int32(id))
  err = templates["editValve"].Execute(w, valve)
  if err != nil {
    log.Println(err)
  }
  return
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
  return
}

func openValve(w http.ResponseWriter, r *http.Request) {
    relay, _ := strconv.Atoi(r.URL.Path[len("/open/"):])

    valve, err := models.GetValveByRelayId(relay)

    if err == nil && valve != nil {
      valve.Open()

    }
    http.Redirect(w, r, "/", 302)
  return
}

func closeValve(w http.ResponseWriter, r *http.Request) {
    relay, _ := strconv.Atoi(r.URL.Path[len("/close/"):])

    valve, err := models.GetValveByRelayId(relay)

    if err == nil && valve != nil {
      valve.Close()

    }
    http.Redirect(w, r, "/", 302)
  return
}

