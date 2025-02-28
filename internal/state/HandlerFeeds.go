package state

import (
	"context"
	"fmt"
)

func HandlerFeeds(s *State, cmd Command) error {

	ctx := context.Background()
	feedSlice, err := s.Db.GetFeeds(ctx)
	if err != nil {
		return err
	}

	for _, feed := range feedSlice {
		user, err := s.Db.GetUserName(ctx, feed.UserID)
		if err != nil {
			return err
		}
		fmt.Printf("Feed: %v\n", feed.Name)
		fmt.Printf("URL: %v\n", feed.Url)
		fmt.Printf("Created by: %v\n", user.Name)
		fmt.Println("")
	}
	if len(feedSlice) == 0 {
		fmt.Println("no feeds found")
	}
	return nil
}
