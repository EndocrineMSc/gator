package command

import (
	"context"
	"fmt"
	"os"
)

func HandlerFeeds(s *State, cmd Command) error {
	feeds, err := s.Db.GetAllFeeds(context.Background())
	if err != nil {
		fmt.Printf("get all feeds failed: %v\n", err)
		os.Exit(1)
	}

	if len(feeds) == 0 {
		fmt.Println("no feeds found")
		os.Exit(0)
	}

	for _, feed := range feeds {
		user, err := s.Db.GetUserById(context.Background(), feed.UserID)
		if err != nil {
			fmt.Printf("get user failed: %v\n", err)
			os.Exit(1)
		}
		fmt.Println(feed.Name)
		fmt.Println(feed.Url)
		fmt.Println(user.Name)
		fmt.Println("---------")
	}

	os.Exit(0)
	return nil
}
