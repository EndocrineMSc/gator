package command

import (
	"context"
	"encoding/xml"
	"fmt"
	"io"
	"net/http"
	"net/mail"
	"os"
	"strings"
	"time"

	"github.com/EndocrineMSc/gator/internal/database"
	"github.com/google/uuid"
)

func HandlerAgg(s *State, cmd Command) error {
	if len(cmd.Args) == 0 {
		fmt.Println("Usage: agg <time_between_reqs>")
		os.Exit(1)
	}

	dur, err := time.ParseDuration(cmd.Args[0])
	if err != nil {
		fmt.Println("Invalid time duration")
		os.Exit(1)
	}

	fmt.Printf("Collecting feeds every %v\n", dur)
	ticker := time.NewTicker(dur)
	for ; ; <-ticker.C {
		err = scrapeFeeds(context.Background(), s)
		if err != nil {
			fmt.Printf("Failed to scrape feeds: %v\n", err)
			os.Exit(1)
		}
		fmt.Println("--------------")
	}

	return nil
}

func fetchFeed(ctx context.Context, feedUrl string) (*RSSFeed, error) {
	req, err := http.NewRequestWithContext(ctx, "GET", feedUrl, nil)
	if err != nil {
		return nil, fmt.Errorf("could not create request: %w", err)
	}
	req.Header.Set("User-Agent", "gator")

	client := http.DefaultClient

	res, err := client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("could not fetch feed: %w", err)
	}
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			fmt.Printf("could not close body: %v\n", err)
		}
	}(res.Body)

	body, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, fmt.Errorf("could not read feed: %w", err)
	}

	var feed *RSSFeed
	err = xml.Unmarshal(body, &feed)
	if err != nil {
		return nil, fmt.Errorf("could not unmarshal feed: %w", err)
	}
	feed.Unescape()

	return feed, nil
}

func scrapeFeeds(ctx context.Context, s *State) error {
	nextFeed, err := s.Db.GetNextFeedToFetch(ctx)
	if err != nil {
		return fmt.Errorf("scrapeFeeds get next feed to fetch feeds: %w", err)
	}

	err = s.Db.MarkFeedFetchedById(ctx, nextFeed.ID)
	if err != nil {
		return fmt.Errorf("scrapeFeeds mark feed to fetch feeds: %w", err)
	}

	feed, err := fetchFeed(ctx, nextFeed.Url)
	if err != nil {
		return fmt.Errorf("scrapeFeeds fetch feed: %w", err)
	}

	for _, i := range feed.Channel.Item {
		publishedAt, err := mail.ParseDate(i.PubDate)
		if err != nil {
			fmt.Println(i.PubDate)
			fmt.Printf("scrapeFeeds fetch feed item pub date: %v\n", err)
			os.Exit(1)
		}
		postParams := database.CreatePostParams{
			ID:          uuid.New(),
			Title:       i.Title,
			Url:         i.Link,
			Description: i.Description,
			PublishedAt: publishedAt,
			FeedID:      nextFeed.ID,
		}
		post, err := s.Db.CreatePost(ctx, postParams)
		if err != nil {
			if strings.Contains(err.Error(), "(23505)") {
				continue
			}
			fmt.Printf("scrapeFeeds create post: %v\n", err)
			os.Exit(1)
		}
		fmt.Println(post.Title)
	}

	return nil
}
