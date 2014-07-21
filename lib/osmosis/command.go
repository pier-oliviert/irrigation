package main

import (
  "errors"
  "strconv"
  "strings"
  )

type Command struct {
  name string
  id int64
}

func NewCommand(infos []string) (*Command, error) {
  if len(infos) < 2 {
    return nil, errors.New("Not enough information to determine a command")
  }
  cmd := new(Command)
  cmd.name = infos[0]
  id, err := strconv.ParseInt(infos[1], 10, 64)
  if err != nil {
    return nil, err
  }

  cmd.id = id

  return cmd, nil
}

func (c *Command) Execute() error {
  if c.name == "open" {
    c.open()
  } else if c.name == "close" {
    c.close()
  }
  return nil
}

func (c *Command) String() string {
  var event []string
  event = append(event, c.name)
  event = append(event, strconv.FormatInt(c.id, 10))
  return strings.Join(event, ":")
}

func (c *Command) open() error {
  rows, err := db.Query("SELECT zones.gpio, zones.name FROM zones inner join sprinkles on (sprinkles.zone_id = zones.id) WHERE sprinkles.id = $1", c.id)
  if err != nil {
    return err
  }

  for rows.Next() {
    var name string
    var id int
    if err := rows.Scan(&id, &name); err != nil {
      return err
    }
    gpio.Open(int64(id))
  }
  return nil
}

func (c *Command) close() error {
  rows, err := db.Query("SELECT zones.gpio, zones.name FROM zones inner join sprinkles on (sprinkles.zone_id = zones.id) WHERE sprinkles.id = $1", c.id)
  if err != nil {
    return err
  }

  for rows.Next() {
    var name string
    var id int
    if err := rows.Scan(&id, &name); err != nil {
      return err
    }
    gpio.Close(int64(id))
  }
  return nil
}
