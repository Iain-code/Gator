package state

import (
	"context"
	"fmt"
	"gator/internal/database"
)

func HandlerFollowing(s *State, cmd Command, user database.User) error {

	ctx := context.Background()
	userFeeds, err := s.Db.GetFeedFollowsForUser(ctx, user.ID)
	if err != nil {
		return fmt.Errorf("could not fetch feed: %v", err)
	}
	if len(userFeeds) == 0 {
		fmt.Printf("user %v is not following any feeds\n", s.Cfg.CurrentUserName)
	}
	for _, feed := range userFeeds {

		fmt.Printf("Feed names for current user: %v\n", feed.FeedName)

	}
	return nil
}
