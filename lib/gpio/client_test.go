package main

import (
  "testing"
  "net"
  )

func TestAddingClient(t *testing.T) {
  length := len(clients)

  connectPair(t)

  if length + 2 != len(clients) {
    t.Log("The client was not added to the slice")
    t.Fail()
  }
}

func TestReadingByte(t *testing.T) {
  t.Log("Helo")
  t.Log(clients)
  client := clients[0]
  data := make([]byte, 32)
  client.Conn.Write([]byte("Hello"))
  length, err := client.Read(data)
  if err != nil || length == 0 {
    t.Log("Couldn't read data")
    t.Fail()
  }
}

func TestDeconnectingClient(t *testing.T) {
  for len(clients) > 0 {
    RemoveClient(clients[0])
  }

  if len(clients) > 0 {
    t.Log("There should not be any connected client at this point")
    t.Fail()
  }
}

func connectPair(t *testing.T) {
  conn1, conn2 := net.Pipe()
  socks := []net.Conn{conn1, conn2}
  for i := 0; i < 2; i++ {
    if len(clients) - 1 < i {
      c := AddClient(socks[i])
      if c == nil {
        t.Log("Couldn't add a client")
        t.Fail()
      }
    }
  }
}
