package main

import (
  "os"
  "os/exec"
  "log"
  "fmt"
)

func fetchRepository(path string) {

  err := os.RemoveAll(path)
  if err != nil {
    log.Fatalln(fmt.Sprintf(`Couldn't reset %v
    %v`, path, err))
  }

  _, err = os.Open(path)
  if os.IsNotExist(err) {
    err = os.MkdirAll(path, os.ModeDir | 0744)
    if err != nil {
      log.Fatalln(fmt.Sprintf(`Couldn't create the directory %v.
      Make sure the path is a valid directory path and that you can create a directory
      in the parent path with your current user.

      Error: %v`, path, err))
    }
  }

  cmd := exec.Command("git", "clone", "https://github.com/pothibo/irrigation-assets", path)
  out, err := cmd.CombinedOutput()

  if err != nil {
    log.Fatalln(fmt.Sprintf(`An error happened when trying to clone irrigation's asset repository.
    Output: %s
    Error: %v`, out, err))
  }
}

func setValueFor(key string, options map[string]string) {
  value := options[key]
  fmt.Scanln(&value)

  if len(value) != 0 {
    options[key] = value
  }
}
