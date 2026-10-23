package command

import (
	"context"
	"fmt"
	"os"
	"strconv"

	"github.com/EndocrineMSc/gator/internal/database"
)

func HandlerBrowse(s *State, cmd Command, user database.User) error {
	limit := "2"
	if len(cmd.Args) > 0 {
		limit = cmd.Args[0]
	}

	int64Limit, err := strconv.ParseInt(limit, 10, 32)
	if err != nil {
		fmt.Printf("Invalid limit value: %s\n", limit)
		os.Exit(1)
	}
	intLimit := int32(int64Limit)
	params := database.GetPostsForUserByIdParams{
		UserID: user.ID,
		Limit:  intLimit,
	}

	posts, err := s.Db.GetPostsForUserById(context.Background(), params)
	if err != nil {
		fmt.Printf("Get posts failed: %s\n", err)
		os.Exit(1)
	}

	for _, post := range posts {
		fmt.Println(post)
	}

	os.Exit(0)
	return nil
}
