package rss

import (
	"context"
	"encoding/xml"
	"fmt"
	"html"
	"io"
	"net/http"
)

type RSSFeed struct {
	Channel struct {
		Title       string    `xml:"title"`
		Link        string    `xml:"link"`
		Description string    `xml:"description"`
		Item        []RSSItem `xml:"item"`
	} `xml:"channel"`
}

type RSSItem struct {
	Title       string `xml:"title"`
	Link        string `xml:"link"`
	Description string `xml:"description"`
	PubDate     string `xml:"pubDate"`
}

func FetchFeed(ctx context.Context, feedURL string) (*RSSFeed, error) {

	rss := &RSSFeed{}
	req, err := http.NewRequestWithContext(ctx, "GET", feedURL, nil) // CREATES the http:request so its ready to send
	req.Header.Set("User-Agent", "gator")
	if err != nil {
		return nil, err
	}
	fmt.Println("Get request made...")

	client := &http.Client{}    // make the GO initialized client struct
	resp, err := client.Do(req) // actually send the Request - and get the response
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	fmt.Println("Get request sent...")

	response, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	fmt.Println("Respose Received...")
	err = xml.Unmarshal(response, rss)
	fmt.Println("Respose Unmarshalled...")
	if err != nil {
		return nil, err
	}
	fmt.Printf("rss.Channel.Title: %v\n", rss.Channel.Title)

	rss.Channel.Title = html.UnescapeString(rss.Channel.Title) // removes all replacement char for things like < > ! etc
	rss.Channel.Description = html.UnescapeString(rss.Channel.Description)

	for i := range rss.Channel.Item {
		rss.Channel.Item[i].Title = html.UnescapeString(rss.Channel.Item[i].Title)
		rss.Channel.Item[i].Description = html.UnescapeString(rss.Channel.Item[i].Description)
	}
	fmt.Println("Returning RSS feed...")
	return rss, nil
}
