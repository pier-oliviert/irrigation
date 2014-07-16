package main

import (
  "net"
  "log"
  "fmt"
  "strings"
  )

func listen(c net.Conn) {
  defer close(c)
  for {
    buf := make([]byte, 32)
    n, err := c.Read(buf)
    if err != nil {
      log.Print(err)
      return
    }
    command, err := NewCommand(strings.Split(string(buf[0:n]), ":"))

    if err != nil {
      log.Print(err)
      return
    }

    err = command.Execute()
    if err != nil {
      log.Print(err)
      return
    }

    notify(connections, command)
  }
}

func close(c net.Conn) {
  c.Close()
  s := connections
  i := indexOf(s, c)
  if i < 0 {
    return
  }
	copy(s[i:], s[i+1:])
	s[len(s)-1] = nil
	connections = s[:len(s)-1]
  fmt.Printf("*** Connection to a client is now closed.\n", c)
}

func indexOf(slice []net.Conn, c net.Conn) int {
  for p, v := range slice {
    if v == c {
      return p
    }
  }
  return -1
}
