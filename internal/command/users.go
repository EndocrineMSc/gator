package command

import (
	"context"
	"fmt"
	"os"
)

func HandlerUsers(s *State, _ Command) error {
	users, err := s.Db.GetAllUsers(context.Background())
	if err != nil {
		fmt.Printf("Users could not be retrieved: %v\n", err)
		os.Exit(1)
	}
	currentUser := s.Cfg.CurrentUserName
	if len(users) == 0 {
		fmt.Println("No users found")
		return nil
	}
	for _, user := range users {
		if user.Name == currentUser {
			fmt.Printf("* %s (current)\n", user.Name)
		} else {
			fmt.Printf("* %s", user.Name)
		}
	}
	return nil
}
