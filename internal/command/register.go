package command

import (
	"context"
	"errors"
	"fmt"
	"os"
	"time"

	"github.com/EndocrineMSc/gator/internal/database"
	"github.com/google/uuid"
)

func HandlerRegister(s *State, cmd Command) error {
	if len(cmd.Args) == 0 {
		return errors.New("no arguments")
	}
	name := cmd.Args[0]
	_, err := s.Db.GetUserByName(
		context.Background(),
		name,
	)
	if err == nil {
		fmt.Println("user already exists")
		os.Exit(1)
	}

	params := database.CreateUserParams{
		ID:        uuid.New(),
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		Name:      name,
	}

	user, err := s.Db.CreateUser(
		context.Background(),
		params,
	)
	if err != nil {
		return fmt.Errorf("create user in db: %w", err)
	}
	err = s.Cfg.SetUser(name)
	if err != nil {
		return fmt.Errorf("set user to config: %w", err)
	}

	fmt.Printf(
		"User registered:\n  ID: %v\n  Created At: %v\n  Updated At: %v\n  Name: %s\n",
		user.ID,
		user.CreatedAt,
		user.UpdatedAt,
		user.Name,
	)
	return nil
}
