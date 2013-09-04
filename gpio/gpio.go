package gpio

import (
	"errors"
	"fmt"
	"os"
	"strconv"
)

func Open(relay uint8) error {
	path := fmt.Sprintf("/sys/class/gpio/gpio%d/value", relay)
	file, err := os.OpenFile(path, os.O_WRONLY|os.O_APPEND, os.ModeDevice)
	if err != nil {
		return err
	}

	defer file.Close()

	_, err = file.WriteString("0")

	return err
}

func Close(relay uint8) error {
	path := fmt.Sprintf("/sys/class/gpio/gpio%d/value", relay)
	file, err := os.OpenFile(path, os.O_WRONLY|os.O_APPEND, os.ModeDevice)
	if err != nil {
		return err
	}

	defer file.Close()

	_, err = file.WriteString("1")

	return err
}

func IsOpened(relay uint8) (opened bool, err error) {
	path := fmt.Sprintf("/sys/class/gpio/gpio%d/value", relay)
	file, err := os.OpenFile(path, os.O_RDONLY, os.ModeDevice)
	if err != nil {
		return false, err
	}

	defer file.Close()

	data := make([]byte, 1)
	_, err = file.Read(data)

	if err != nil {
		return false, err
	}

	return string(data) == "0", nil
}

func IsClosed(relay uint8) (opened bool, err error) {
	opened, err = IsOpened(relay)
	return !opened, err
}

func Activate(relay uint8) error {
	if os.Getuid() != 0 {
		return errors.New("You have to be root to activate relays")
	}

	err := export(relay)
	if err != nil {
		return err
	}

	err = setRelayAsOutput(relay)
	if err != nil {
		return err
	}

	return nil
}

func export(relay uint8) error {
	file, err := os.OpenFile("/sys/class/gpio/export", os.O_WRONLY|os.O_APPEND, os.ModeDevice)
	if err != nil {
		return err
	}

	defer file.Close()

	_, err = file.WriteString(strconv.Itoa(int(relay)))


  return err
}

func setRelayAsOutput(relay uint8) error {
	path := fmt.Sprintf("/sys/class/gpio/gpio%d", relay)

	file, err := os.OpenFile(path+"/direction", os.O_WRONLY|os.O_APPEND, os.ModeDevice)
	if err != nil {
		return err
	}

	defer file.Close()

	_, err = file.WriteString("out")

	err = makeValueAvailable(path + "/value")
  if err != nil {
      return err
  }

  return Close(relay)

}

func makeValueAvailable(path string) error {

	file, err := os.OpenFile(path, os.O_WRONLY, os.ModeDevice)
	if err != nil {
		return err
	}

	defer file.Close()

	return file.Chmod(0666)
}
