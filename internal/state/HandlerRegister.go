package state

import (
	"context"
	"fmt"
	"gator/internal/database"
	"log"
	"os"
	"time"

	"github.com/google/uuid"
)

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

	fmt.Println("New user created")

	fmt.Println(userParams.ID)
	fmt.Println(userParams.CreatedAt)
	fmt.Println(userParams.UpdatedAt)
	fmt.Println(userParams.Name)
	return nil

}
