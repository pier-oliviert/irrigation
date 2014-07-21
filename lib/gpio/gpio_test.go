package main

import(
  "testing"
  "github.com/stianeikeland/go-rpio"
  "encoding/json"
)

func TestPinInitializations(t *testing.T) {
  initializeRPIO(t)
  defer closeRPIO(t)

  pins = initializePins(gpios)
  if len(pins) != 16 {
    t.Log("Not all pins were initialized")
    t.Fail()
  }
}

func TestListGPIO(t *testing.T) {
  initializeRPIO(t)
  defer closeRPIO(t)

  pins = initializePins(gpios)
  type JSONPin struct {
    Id int `json:"id"`
    Status int `json:"status"`
  }

  var result []JSONPin

  connectPair(t)
  client := clients[0]
  ListGPIO()
  data := <- client.Notify 

  json.Unmarshal([]byte(data), &result)
  if len(result) != len(pins) {
    t.Log("Pin status sent was not the same as the amount of pins")
    t.Fail()
  }
}

func initializeRPIO(t *testing.T) {
  if err:= rpio.Open(); err != nil {
    t.Log("Couldn't open GPIO memory space")
    t.Fail()
  }
}

func closeRPIO(t *testing.T) {
  rpio.Close()
}
