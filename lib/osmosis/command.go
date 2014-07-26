package main

import (
	"encoding/json"
	"errors"
	"log"
	"strconv"
)

type Command struct {
	Name string `json:"name"`
	Id   int64  `json:"id"`
}

func NewCommand(infos []string) (*Command, error) {
	if len(infos) < 2 {
		return nil, errors.New("Not enough information to determine a command")
	}
	cmd := new(Command)
	cmd.Name = infos[0]
	id, err := strconv.ParseInt(infos[1], 10, 64)
	if err != nil {
		return nil, err
	}

	cmd.Id = id

	return cmd, nil
}

func (c *Command) Bytes() []byte {
	cmd := make(map[string]Command)
	cmd["action"] = *c
	d, err := json.Marshal(cmd)
	if err != nil {
		log.Print(err)
	}

	return d
}
