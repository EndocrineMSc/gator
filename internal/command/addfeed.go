package command

import (
	"context"
	"fmt"
	"os"
	"time"

	"github.com/EndocrineMSc/gator/internal/database"
	"github.com/google/uuid"
)

func HandlerAddfeed(s *State, cmd Command, user database.User) error {
	if len(cmd.Args) < 2 {
		fmt.Printf("add feed requires at least two arguments - name and url of the feed\n")
		os.Exit(1)
	}

	feedName := cmd.Args[0]
	feedUrl := cmd.Args[1]

	_, err := fetchFeed(context.Background(), feedUrl)
	if err != nil {
		fmt.Printf("error fetching feed %s: %v\n", feedUrl, err)
		os.Exit(1)
	}

	feedParams := database.CreateFeedParams{
		ID:        uuid.New(),
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		Name:      feedName,
		Url:       feedUrl,
		UserID:    user.ID,
	}
	feedRecord, err := s.Db.CreateFeed(context.Background(), feedParams)
	if err != nil {
		fmt.Printf("error creating feed %s: %v\n", feedUrl, err)
		os.Exit(1)
	}

	followParams := database.CreateFeedFollowParams{
		ID:        uuid.New(),
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		UserID:    user.ID,
		FeedID:    feedRecord.ID,
	}
	_, err = s.Db.CreateFeedFollow(context.Background(), followParams)
	if err != nil {
		fmt.Printf("error creating feed follow %s: %v\n", feedUrl, err)
		os.Exit(1)
	}

	fmt.Printf("created feed: %v\n", feedRecord)
	os.Exit(0)
	return nil
}
