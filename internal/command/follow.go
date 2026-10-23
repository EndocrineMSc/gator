package command

import (
	"context"
	"fmt"
	"os"
	"time"

	"github.com/EndocrineMSc/gator/internal/database"
	"github.com/google/uuid"
)

func HandlerFollow(s *State, cmd Command, user database.User) error {
	if len(cmd.Args) == 0 {
		fmt.Println("feed url is needed")
		os.Exit(1)
	}

	feed, err := s.Db.GetFeedByUrl(context.Background(), cmd.Args[0])
	if err != nil {
		fmt.Println("feed not found")
		os.Exit(1)
	}

	feedFollowParams := database.CreateFeedFollowParams{
		ID:        uuid.New(),
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		UserID:    user.ID,
		FeedID:    feed.ID,
	}

	feedFollow, err := s.Db.CreateFeedFollow(context.Background(), feedFollowParams)
	if err != nil {
		fmt.Printf("create feed follow err:%s\n", err)
		os.Exit(1)
	}

	fmt.Printf("feed follow created:%v\n", feedFollow)
	os.Exit(0)
	return nil
}
