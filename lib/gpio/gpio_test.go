package main

import(
  "testing"
)

type JSONPin struct {
  Id int `json:"id"`
  Status int `json:"status"`
}

func init() {
  connectPair()
  pins = initializePins(gpios)
}

func TestPinInitializations(t *testing.T) {

  pins = initializePins(gpios)
  if len(pins) != 16 {
    t.Log("Not all pins were initialized")
    t.Fail()
  }
}

func TestOpenGPIO(t *testing.T) {
  connectPair()
  go func() {
    _ = <-clients[0].Notify
  }()

  id := gpios[3]
  pin := pins[id]
  OpenGPIO(int64(id))

  if pin.State() != 1 {
    t.Log("Pin should be opened")
    t.Fail()
  }
}

func TestListGPIO(t *testing.T) {
  ListGPIO()
}

func TestCloseGPIO(t *testing.T) {
  id := gpios[3]
  pin := pins[id]
  CloseGPIO(int64(id))

  if pin.State() != 0 {
    t.Log("Pin should be closed")
    t.Fail()
  }
}

