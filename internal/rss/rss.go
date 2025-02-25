package rss

import (
	"context"
	"encoding/xml"
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

	client := &http.Client{}    // make the GO initialized client struct
	resp, err := client.Do(req) // actually send the Request - and get the response
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	response, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	err = xml.Unmarshal(response, rss)

	if err != nil {
		return nil, err
	}

	html.UnescapeString(rss.Channel.Title) // removes all replacement char for things like < > ! etc
	html.UnescapeString(rss.Channel.Description)

	for _, item := range rss.Channel.Item {
		html.UnescapeString(item.Title)
		html.UnescapeString(item.Description)
	}

	return rss, nil
}
