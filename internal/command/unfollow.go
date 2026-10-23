package command

import (
	"context"
	"fmt"
	"os"

	"github.com/EndocrineMSc/gator/internal/database"
)

func HandlerUnfollow(s *State, cmd Command, user database.User) error {
	if len(cmd.Args) == 0 {
		fmt.Println("You need to provide the feed url")
		os.Exit(1)
	}

	feed, err := s.Db.GetFeedByUrl(context.Background(), cmd.Args[0])
	if err != nil {
		fmt.Printf("Get feed by url err (likely wrong url): %v\n", err)
		os.Exit(1)
	}

	deleteParams := database.DeleteByUserIdAndFeedIdParams{
		UserID: user.ID,
		FeedID: feed.ID,
	}

	err = s.Db.DeleteByUserIdAndFeedId(context.Background(), deleteParams)
	if err != nil {
		fmt.Printf("Delete feed by id err: %v\n", err)
		os.Exit(1)
	}

	os.Exit(0)
	return nil
}
