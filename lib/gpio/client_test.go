package main

import (
  "testing"
  "log"
  "encoding/json"
  "net"
  )

func TestAddingClient(t *testing.T) {
  length := len(clients)

  connectPair()

  if length + 2 != len(clients) {
    t.Logf("The client was not added to the slice. # of clients connected: %d", len(clients))
    t.Fail()
  }
}

func TestSendOpenActionToClient(t *testing.T) {
  client := clients[0]

  payload := make(map[string]Action)

  payload["action"] = Action{
    Name: "open",
    Id: 10,
  }

  d, err := json.Marshal(payload)
  if err != nil {
    t.Log(err)
    t.Fail()
  }

  client.Conn.Write(d)

  pin := pins[gpios[3]]
  if pin.State() != 1 {
    t.Logf("The GPIO is not opened: %d", pin.Id())
    t.Fail()
  }
}

func TestSendCloseActionToClient(t *testing.T) {
  client := clients[1]

  payload := make(map[string]Action)

  payload["action"] = Action{
    Name: "close",
    Id: 10,
  }

  d, err := json.Marshal(payload)
  if err != nil {
    t.Log(err)
    t.Fail()
  }

  client.Conn.Write(d)

  pin := pins[gpios[3]]
  if pin.State() != 0 {
    t.Logf("The GPIO is still opened: %d", pin.Id())
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

func connectPair() {
  conn1, conn2 := net.Pipe()
  socks := []net.Conn{conn1, conn2}
  for i := 0; i < 2; i++ {
    c := AddClient(socks[i])
    if c == nil {
      log.Fatal("Couldn't add a client")
    }
  }
}
