package command

import (
	"context"
	"errors"
	"fmt"
	"os"
)

func HandlerLogin(s *State, cmd Command) error {
	if len(cmd.Args) == 0 {
		return errors.New("invalid arguments")
	}
	name := cmd.Args[0]
	_, err := s.Db.GetUserByName(context.Background(), name)
	if err != nil {
		fmt.Printf("user doesn't exist: %v\n", err)
		os.Exit(1)
	}
	err = s.Cfg.SetUser(name)
	if err != nil {
		return fmt.Errorf("set user: %w", err)
	}
	fmt.Println("set user:", name)
	return nil
}
