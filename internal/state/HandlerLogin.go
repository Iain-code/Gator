package state

import (
	"context"
	"fmt"
	"gator/internal/config"
	"log"
	"os"
)

func HandlerLogin(s *State, cmd Command) error {

	ctx := context.Background()
	user, err := s.Db.GetUser(ctx, cmd.Arg[0])
	// s.Db is a struct that contains all the QUERIES made by "sqlc generate"

	if user.Name == "" {
		fmt.Printf("no user named %v found\n", cmd.Arg[0])
	}

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
