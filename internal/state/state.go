package state

import (
	"context"
	"errors"
	"fmt"
	"gator/internal/config"
	"gator/internal/database"
	"gator/internal/rss"
	"log"
	"os"
	"time"

	"github.com/google/uuid"
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

func HandlerLogin(s *State, cmd Command) error {

	ctx := context.Background()
	_, err := s.Db.GetUser(ctx, cmd.Arg[0])
	// s.Db is a struct that contains all the QUERIES made by "sqlc generate"

	if err != nil {
		log.Fatal(err)
		os.Exit(1)
	}

	if cmd.Name == "" {
		return fmt.Errorf("no command found")
	}
	if len(cmd.Arg) < 1 {
		log.Fatal("Args too short")
	}

	s.Cfg.CurrentUserName = cmd.Arg[0]
	config.Write(*s.Cfg)
	fmt.Printf("New Username %+v\n", s.Cfg.CurrentUserName)

	return nil
}

func HandlerRegister(s *State, cmd Command) error {

	if cmd.Name == "" {
		return fmt.Errorf("no command found")
	}
	if len(cmd.Arg) < 1 {
		log.Fatal("Args too short")
	}

	ctx := context.Background()
	uuid := uuid.New()

	fmt.Printf("command Arg[0]: %v\n", cmd.Arg[0])

	_, err := s.Db.GetUser(ctx, cmd.Arg[0])
	// s.Db is a struct that contains all the QUERIES made by "sqlc generate"

	if err == nil {
		fmt.Println("get user error lol")
		log.Fatal(err)
		os.Exit(1)
	}

	userParams := database.CreateUserParams{}
	userParams.ID = uuid
	userParams.CreatedAt = time.Now()
	userParams.UpdatedAt = time.Now()
	userParams.Name = cmd.Arg[0]

	newUser, err := s.Db.CreateUser(ctx, userParams)
	if err != nil {
		return err
	}

	s.Cfg.CurrentUserName = newUser.Name
	s.Cfg.SetUser(newUser.Name)

	fmt.Printf("Current User ++: %+v\n", s.Cfg.CurrentUserName)
	fmt.Println("New user created")

	fmt.Println(userParams.ID)
	fmt.Println(userParams.CreatedAt)
	fmt.Println(userParams.UpdatedAt)
	fmt.Println(userParams.Name)
	return nil

}

func HandlerReset(s *State, cmd Command) error {

	ctx := context.Background()
	err := s.Db.DeleteAllUsers(ctx)
	// s.Db is a struct that contains all the QUERIES made by "sqlc generate"

	if err != nil {
		return err
	}
	fmt.Println("All users were deleted")
	return nil
}

func HandlerUsers(s *State, cmd Command) error {

	// this func will show all the users within the database and print it to console
	ctx := context.Background()
	userSlice, err := s.Db.GetUsers(ctx)
	if err != nil {
		return err
	}

	for _, user := range userSlice {
		if user.Name == s.Cfg.CurrentUserName {
			fmt.Printf("%s (current)\n", user.Name)
			fmt.Printf("current user ID: %v\n", user.ID)
		} else {
			fmt.Println(user.Name)
		}
	}
	return nil
}

func HandlerAgg(s *State, cmd Command) error {

	ctx := context.Background()
	url := "https://www.wagslane.dev/index.xml"

	rss, err := rss.FetchFeed(ctx, url)
	if err != nil {
		return err
	}
	fmt.Println(rss.Channel.Title)
	fmt.Println(rss.Channel.Link)
	fmt.Println(rss.Channel.Description)
	fmt.Println(rss.Channel.Item)

	for _, item := range rss.Channel.Item {
		fmt.Println(item.Title)
		fmt.Println(item.Link)
		fmt.Println(item.Description)
		fmt.Println(item.PubDate)
	}

	return nil
}

func HandlerAddFeed(s *State, cmd Command) error {

	// this function adds an "RSS feed" instance into the feed table

	currentuser := s.Cfg.CurrentUserName
	ctx := context.Background()
	arg := database.CreateFeedParams{}

	user, err := s.Db.GetUser(ctx, currentuser)
	if err != nil {
		return err
	}

	if len(cmd.Arg) < 2 {
		log.Fatal("no name and url found")
	} else {

		arg.ID = uuid.New()
		arg.CreatedAt = time.Now()
		arg.UpdatedAt = time.Now()
		arg.Name = cmd.Arg[0]
		arg.Url = cmd.Arg[1]
		arg.UserID = user.ID

		feed, err := s.Db.CreateFeed(ctx, arg)
		if err != nil {
			return err
		}

		fmt.Println(feed.ID)
		fmt.Println(feed.CreatedAt)
		fmt.Println(feed.UpdatedAt)
		fmt.Println(feed.Name)
		fmt.Println(feed.Url)
		fmt.Println(feed.UserID)
	}
	return nil
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
