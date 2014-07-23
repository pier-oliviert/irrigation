package main

import (
  "errors"
  "strconv"
  "encoding/json"
  "log"
  )

type Command struct {
  Name string `json:"name"`
  Id int64 `json:"id"`
}

func NewCommand(infos []string) (*Command, error) {
  if len(infos) < 2 {
    return nil, errors.New("Not enough information to determine a command")
  }
  cmd := new(Command)
  cmd.Name = infos[0]
  id, err := strconv.ParseInt(infos[1], 10, 64)
  if err != nil {
    return nil, err
  }

  cmd.Id = id

  return cmd, nil
}

func (c *Command) Execute() error {
  if c.Name == "open" || c.Name == "close" {
    rows, err := db.Query("SELECT zones.gpio FROM zones inner join sprinkles on (sprinkles.zone_id = zones.id) WHERE sprinkles.id = $1", c.Id)
    if err != nil {
      return err
    }

    for rows.Next() {
      var id int64
      if err := rows.Scan(&id); err != nil {
        return err
      }
      cmd := &Command{
        Name: c.Name,
        Id: id,
      }
      gpio.Send(cmd)
    }
    return nil
  }
  return nil
}

func (c *Command) Bytes() []byte {
  cmd := make(map[string]Command)
  cmd["action"] = *c
  d, err := json.Marshal(cmd)
  if err != nil {
    log.Print(err)
  }

  return d
}
