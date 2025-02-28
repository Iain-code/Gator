package state

import (
	"context"
	"fmt"
	"gator/internal/database"
	"strconv"
)

func HandlerBrowse(s *State, cmd Command, user database.User) error {

	limit := 2
	if len(cmd.Arg) == 1 {
		if specifiedLimit, err := strconv.Atoi(cmd.Arg[0]); err == nil {
			limit = specifiedLimit
		} else {
			return fmt.Errorf("invalid limit: %w", err)
		}
	}

	ctx := context.Background()
	getPosts := database.GetPostsForUserParams{}
	getPosts.UserID = user.ID
	getPosts.Limit = int32(limit) // limit is just an int and must be cast into correct format

	posts, err := s.Db.GetPostsForUser(ctx, getPosts)
	if err != nil {
		return fmt.Errorf("GetPostsForUser did not return correct values")
	}
	fmt.Println(posts)

	fmt.Printf("Found %d posts for user %s:\n", len(posts), user.Name)
	for _, post := range posts {
		fmt.Printf("%s from %s\n", post.PublishedAt.Format("Mon Jan 2"), post)
		fmt.Printf("--- %s ---\n", post.Title)
		fmt.Printf("    %v\n", post.Description)
		fmt.Printf("Link: %s\n", post.Url)
		fmt.Println("=====================================")
	}

	return nil
}
