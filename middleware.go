package main

import (
	"context"
	"os"

	"github.com/EndocrineMSc/gator/internal/command"
	"github.com/EndocrineMSc/gator/internal/database"
)

func middlewareLoggedIn(handler func(s *command.State, cmd command.Command, user database.User) error) func(*command.State, command.Command) error {
	return func(s *command.State, cmd command.Command) error {
		user, err := s.Db.GetUserByName(context.Background(), s.Cfg.CurrentUserName)
		if err != nil {
			os.Exit(1)
			return err
		}

		return handler(s, cmd, user)
	}
}
