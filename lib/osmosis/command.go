package main

import (
  "errors"
  "strconv"
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
  } else {
    c.close()
  }
  return nil
}

func (c *Command) open() error {
  rows, err := db.Query("SELECT zones.gpio, zones.name FROM zones inner join sprinkles on (sprinkles.zone_id = zones.id) WHERE sprinkles.id = $1", c.id)
  if err != nil {
    return err
  }

  for rows.Next() {
    var name string
    var gpio int
    if err := rows.Scan(&gpio, &name); err != nil {
      return err
    }
    // Open the GPIO.
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
    var gpio int
    if err := rows.Scan(&gpio, &name); err != nil {
      return err
    }
    // Close the GPIO.
  }
  return nil
}
