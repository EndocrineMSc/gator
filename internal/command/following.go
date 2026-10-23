package command

import (
	"context"
	"fmt"
	"os"

	"github.com/EndocrineMSc/gator/internal/database"
)

func HandlerFollowing(s *State, _ Command, user database.User) error {
	feedFollows, err := s.Db.GetFeedFollowsForUser(context.Background(), user.ID)
	if err != nil {
		fmt.Printf("Couldn't get follows: %s %v\n", s.Cfg.CurrentUserName, err)
		os.Exit(1)
	}

	for _, follow := range feedFollows {
		fmt.Printf("Following: %s\n", follow.FeedName)
	}
	os.Exit(0)
	return nil
}
