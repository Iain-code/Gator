package state

import (
	"context"
	"fmt"
	"gator/internal/database"
)

// wants to check all the feeds from a user using user.id -> user_id
// then check which feed it is using the feed.url
// delete that feed follows thingy

func HandlerUnfollow(s *State, cmd Command, user database.User) error {

	ctx := context.Background()
	deleteFollow := database.DeleteFollowParams{}
	deleteFollow.UserID = user.ID
	deleteFollow.Url = cmd.Arg[0]
	fmt.Printf("URL to unfollow: %v\n", deleteFollow.Url)

	feeds, err := s.Db.GetFeedFollowsForUser(ctx, user.ID)
	if err != nil {
		return err
	}
	for _, feed := range feeds {
		fmt.Printf("feed.FeedUrl: %v\n", feed.FeedUrl)
		if feed.FeedUrl == deleteFollow.Url {

			fmt.Printf("User: %v, has stopped following RSS feed: %v\n", user.Name, deleteFollow.Url)
			err = s.Db.DeleteFollow(ctx, deleteFollow)
			if err != nil {
				fmt.Printf("Unfollow failed: %v", err)
				return err
			}
		}
	}
	return nil
}
