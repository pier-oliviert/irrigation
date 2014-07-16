package main

import (
  "net"
  "strings"
  "strconv"
  "log"
  )

func notify(connections []net.Conn, cmd *Command) {
  for i := 0; i < len(connections); i++ {
    conn := connections[i]
    var event []string
    event = append(event, cmd.name)
    event = append(event, strconv.FormatInt(cmd.id, 10))
    conn.Write([]byte(strings.Join(event, ":")))
  }
}
