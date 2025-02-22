package main

import (
	"fmt"
	"gator/internal/config"
	"gator/internal/state"
	"log"
	"os"
)

func main() {

	rcfg, err := config.Read()
	if err != nil {
		fmt.Println(err)
	}

	newstate := &state.State{}
	newstate.Cfg = &rcfg

	cmds := state.Commands{
		RegisteredCommands: make(map[string]func(*state.State, state.Command) error),
	}
	cmds.Register("login", state.HandlerLogin)

	if len(os.Args) < 2 {
		log.Fatal("not enough args")
	}

	input := os.Args

	newcommand := state.Command{}
	newcommand.Name = input[1]

	newcommand.Arg = input[2:]

	err = cmds.Run(newstate, newcommand) // passes the config file (newstate) and the Command (with .Name and .Args into RUN)
	if err != nil {
		log.Fatal(err)
	}

}
