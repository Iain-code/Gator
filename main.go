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

	cmds.Register("login", state.HandlerLogin)                                   // user login
	cmds.Register("register", state.HandlerRegister)                             // registers a user
	cmds.Register("reset", state.HandlerReset)                                   // resets all users (deletes)
	cmds.Register("users", state.HandlerUsers)                                   // lists all users
	cmds.Register("scrape", state.HandlerAgg)                                    // collects posts from feeds every arg[0] seconds
	cmds.Register("addfeed", state.MiddlewareLoggedIn(state.HandlerAddFeed))     // adds a feed instance when given NAME + URL after addfeed
	cmds.Register("feeds", state.HandlerFeeds)                                   // gets ALL FEEDS and prints the information to console
	cmds.Register("follow", state.MiddlewareLoggedIn(state.HandlerFollow))       // Takes a URL and new feed follow record linking the current user to RSS
	cmds.Register("following", state.MiddlewareLoggedIn(state.HandlerFollowing)) // prints all the names of the feeds the current user is following.
	cmds.Register("unfollow", state.MiddlewareLoggedIn(state.HandlerUnfollow))   // unfollows or deletes an RSS follow for a given user
	cmds.Register("browse", state.MiddlewareLoggedIn(state.HandlerBrowse))       // browses the posts for the logged in user
	cmds.Register("remove", state.HandlerRemove)                                 // removes a feed from the database

	input := os.Args // takes the CLI input as a slice

	newcommand := state.Command{} // make a new Command struct with Name and Arg
	newcommand.Name = input[1]
	newcommand.Arg = input[2:]

	arguement := input[2:]
	if len(arguement) < 1 {
		fmt.Println("")
	} else {
		fmt.Printf("Arguements given: %v\n", arguement)
	}
	err = cmds.Run(newstate, newcommand) // passes the config file (newstate) and the Command (with .Name and .Args into RUN)
	if err != nil {
		log.Fatal(err)
	}

}
