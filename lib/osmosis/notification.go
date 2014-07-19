package main

import (
  "strings"
  "strconv"
  "log"
  )

func notify(clients []*Client, cmd *Command) {
  for i := 0; i < len(clients) - 1; i++ {
    client := clients[i]
    log.Print(cmd)
    var event []string
    event = append(event, cmd.name)
    event = append(event, strconv.FormatInt(cmd.id, 10))
    client.Conn.Write([]byte(strings.Join(event, ":")))
  }
}
