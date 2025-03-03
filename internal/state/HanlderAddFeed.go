package state

import (
	"context"
	"fmt"
	"gator/internal/database"
	"log"
	"time"

	"github.com/google/uuid"
)

func HandlerAddFeed(s *State, cmd Command, user database.User) error {

	// this function adds an "RSS feed" instance into the feed table

	ctx := context.Background()
	arg := database.CreateFeedParams{}

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
		fmt.Printf("added %v RSS feed into %v profile\n", feed.Name, user.Name)
		input := cmd.Arg[1:]
		cmd.Arg = input
		err = HandlerFollow(s, cmd, user)
		if err != nil {
			return err
		}

		fmt.Printf("feed.ID: %v\n", feed.ID)
		fmt.Printf("feed.CreatedAt: %v\n", feed.CreatedAt)
		fmt.Printf("feed.UpdatedAt: %v\n", feed.UpdatedAt)
		fmt.Printf("feed.Name: %v\n", feed.Name)
		fmt.Printf("feed.Url: %v\n", feed.Url)
		fmt.Printf("feed.UserID %v\n", feed.UserID)
	}
	return nil
}
