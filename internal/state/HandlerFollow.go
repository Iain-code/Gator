package state

import (
	"context"
	"fmt"
	"gator/internal/database"
	"log"
	"time"

	"github.com/google/uuid"
)

func HandlerFollow(s *State, cmd Command, user database.User) error {

	ctx := context.Background()
	if len(cmd.Arg) < 1 {
		log.Fatal("error args slice not long enough")
	}

	getFeed, err := s.Db.GetFeed(ctx, cmd.Arg[0])
	if err != nil {
		return err
	}

	params := database.CreateFeedFollowParams{}
	params.ID = uuid.New()
	params.UserID = user.ID
	params.FeedID = getFeed.ID
	params.CreatedAt = time.Now()
	params.UpdatedAt = time.Now()

	feedFollow, err := s.Db.CreateFeedFollow(ctx, params)
	if err != nil {
		return err
	}
	fmt.Printf("User now following %v RSS feed\n", getFeed.Url)
	fmt.Printf("Current user name: %v\n", s.Cfg.CurrentUserName)
	fmt.Printf("Current feed name: %v\n", feedFollow.FeedName)
	return nil
}
