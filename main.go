package main

import (
	"database/sql"
	"fmt"
	"gator/internal/config"
	"gator/internal/database"
	"gator/internal/state"
	"log"
	"os"

	_ "github.com/lib/pq"
)

const dbURL = "postgres://postgres:disturbed9@localhost:5432/gator"

func main() {
	fmt.Println("Attempting to connect to database...")
	db, err := sql.Open("postgres", dbURL)
	if err != nil {
		fmt.Printf("Database connection error: %v\n", err)
		log.Fatal(err)
	}
	fmt.Println("Database connection successful!")
	dbQueries := database.New(db)

	rcfg, err := config.Read()
	if err != nil {
		fmt.Println(err)
	}

	newstate := &state.State{}
	newstate.Cfg = &rcfg

	newstate.Db = dbQueries

	cmds := state.Commands{
		RegisteredCommands: make(map[string]func(*state.State, state.Command) error),
	}
	cmds.Register("login", state.HandlerLogin)
	cmds.Register("register", state.HandlerRegister)
	cmds.Register("reset", state.HandlerReset)
	cmds.Register("users", state.HandlerUsers)
	cmds.Register("agg", state.HandlerAgg)         //
	cmds.Register("addfeed", state.HandlerAddFeed) // adds a feed instance when given NAME + URL after addfeed

	if len(os.Args) < 2 {
		log.Fatal("not enough args")
	}

	input := os.Args // takes the CLI input as a slice

	newcommand := state.Command{} // make a new Command struct with Name and Arg
	newcommand.Name = input[1]
	newcommand.Arg = input[2:]
	fmt.Println(input[2:])

	err = cmds.Run(newstate, newcommand) // passes the config file (newstate) and the Command (with .Name and .Args into RUN)
	if err != nil {
		log.Fatal(err)
	}

}
