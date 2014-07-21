package main

import (
  "testing"
  "net"
  )

func TestAddingClient(t *testing.T) {
  if len(clients) > 0 {
    t.Log("No client should be connected at this point")
    t.Fail()
  }

  conn, _ := net.Pipe()

  client := AddClient(conn)

  if client == nil {
    t.Log("AddClient should return client")
    t.Fail()
  }

  if len(clients) == 0 {
    t.Log("The client was not added to the slice")
    t.Fail()
  }
}

func TestDeconnectingClient(t *testing.T) {
  if len(clients) == 0 {
    conn, _ := net.Pipe()
    AddClient(conn)
  }

  for i := 0; i < len(clients); i++ {
    c := clients[i]
    RemoveClient(c)
  }

  if len(clients) > 0 {
    t.Log("There should not be any connected client at this point")
    t.Fail()
  }
}
