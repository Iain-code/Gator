package state

import (
	"errors"
	"gator/internal/config"
	"gator/internal/database"
)

type State struct {
	Db  *database.Queries // is a struct that contains all the QUERIES made by "sqlc generate"
	Cfg *config.Config
}

type Command struct {
	Name string
	Arg  []string
}

type Commands struct {
	RegisteredCommands map[string]func(*State, Command) error
}

func (c *Commands) Register(name string, f func(*State, Command) error) {

	// this function maps the string given by call in main (key) with the value of the function also given in main
	c.RegisteredCommands[name] = f
}

func (c *Commands) Run(s *State, cmd Command) error {

	handlerFunction, ok := c.RegisteredCommands[cmd.Name]

	if !ok {
		return errors.New("command not found")
	}

	return handlerFunction(s, cmd)
	// handlerFunction here will be whatever function is in the MAP for cmd.Name as
	// we named handlerFunction the value of the map above
}
