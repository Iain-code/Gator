package state

import (
	"errors"
	"fmt"
	"gator/internal/config"
	"log"
)

type State struct {
	Cfg *config.Config
}

type Command struct {
	Name string
	Arg  []string
}

type Commands struct {
	RegisteredCommands map[string]func(*State, Command) error
}

func HandlerLogin(s *State, cmd Command) error {

	if cmd.Name == "" {
		return fmt.Errorf("no username found")
	}
	if len(cmd.Arg) < 1 {
		log.Fatal("Args too short")
	}

	s.Cfg.CurrentUserName = cmd.Arg[0]
	config.Write(*s.Cfg)
	fmt.Printf("New Username %+v\n", s.Cfg.CurrentUserName)
	fmt.Printf("ConfigStruct: %+v\n", s.Cfg)

	return nil
}

func (c *Commands) Register(name string, f func(*State, Command) error) {

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
