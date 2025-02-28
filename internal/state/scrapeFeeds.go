package state

import (
	"context"
	"fmt"
	"gator/internal/database"
	"gator/internal/rss"
	"time"

	"github.com/google/uuid"
)

func scrapeFeeds(s *State) error {

	ctx := context.Background()

	feed, err := s.Db.GetNextFeedToFetch(ctx)
	if err != nil {
		return err
	}
	fmt.Printf("Next feed URL for scraping: %v\n", feed.Url)
	err = s.Db.MarkFeedFetched(ctx, feed.ID)
	if err != nil {
		return err
	}
	fmt.Println("Feed marked as fetched...")

	gotFeed, err := rss.FetchFeed(ctx, feed.Url)
	if err != nil {
		return err
	}
	fmt.Println("Fetched feed...")

	for i := range gotFeed.Channel.Item {
		createPost := database.CreatePostParams{}
		createPost.ID = uuid.New()
		createPost.CreatedAt = time.Now()
		createPost.UpdatedAt = time.Now()
		createPost.Title = gotFeed.Channel.Item[i].Title
		createPost.Url = gotFeed.Channel.Item[i].Link
		createPost.Description = gotFeed.Channel.Item[i].Description
		createPost.FeedID = feed.ID
		fmt.Printf("FEED ID: %v\n", feed.ID)
		layout := "Mon, 02 Jan 2006 15:04:05 -0700"

		createPost.PublishedAt, err = time.Parse(layout, gotFeed.Channel.Item[i].PubDate)

		err := s.Db.CreatePost(ctx, createPost)
		fmt.Printf("Post Created: %v\n", createPost.Title)
		if err != nil {
			return err
		}

	}

	fmt.Println("")

	return nil
}
