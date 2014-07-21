package main

import (
  "syscall"
  "log"
  "net"
  "os"
  "os/signal"
  "encoding/json"
  "github.com/stianeikeland/go-rpio"
  )

var gpios []int = []int{3,5,8,10,11,12,13,15,16,18,19,21,22,23,24,26}
var pins map[int]*Pin

func main() {
  path := "/tmp/gpio.sock"
  if syscall.Getuid() != 0 {
    log.Fatal("Root privilege required to handle GPIO")
  }

  if err := rpio.Open(); err != nil {
    log.Fatal("Couldn't initialize GPIO pins")
  }

  defer rpio.Close()

  pins = initializePins(gpios)

  addr, err := net.ResolveUnixAddr("unix", path)
  handleFatalErr(err)

  ln, err := net.ListenUnix("unix", addr)
  handleFatalErr(err)

  file, err := ln.File()
  handleFatalErr(err)
  defer closeSocket(ln)
  go exit(ln)

  info, err := file.Stat()
  handleFatalErr(err)

  err = os.Chmod(path, info.Mode()|0777)
  handleFatalErr(err)

  for {
    conn, err := ln.Accept()
    if err != nil {
      log.Println("Connection Error: ", err)
    }
    AddClient(conn)
  }
}

func handleFatalErr(err error) {
  if err != nil {
    log.Fatalf("*** Fatal Error: %s", err)
  }
}

func closeSocket(ln *net.UnixListener) {
  ln.Close()
}

func exit(ln *net.UnixListener) {
  sigc := make(chan os.Signal, 1)
  signal.Notify(sigc, os.Interrupt, os.Kill, syscall.SIGTERM)
  go func(c chan os.Signal) {
    sig := <-c
    log.Printf("Caught signal %s: shutting down.", sig)
    ln.Close()
    os.Exit(0)
  }(sigc)
}

func initializePins(gpios []int) map[int]*Pin {
  pins := make(map[int]*Pin)
  for _, id := range gpios {
    pins[id] = NewPin(int64(id))
  }
  return pins
}

func OpenGPIO(id int64) {
  defer ListGPIO()
}

func CloseGPIO(id int64) {
  defer ListGPIO()
}

func ListGPIO() {
  var list []Pin
  for _, value := range pins {
    list = append(list, *value)
  }

  data, err := json.Marshal(list)
  if err != nil {
    log.Fatal(err)
  }
  updateClients(string(data))
}

func updateClients(msg string) {
  for i := 0; i < len(clients); i++ {
    client := clients[i]
    client.Notify <- msg
  }
}
