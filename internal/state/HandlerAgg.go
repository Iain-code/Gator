package state

import (
	"context"
	"fmt"
	"gator/internal/rss"
)

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
