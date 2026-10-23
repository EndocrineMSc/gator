package command

import (
	"context"
	"fmt"
	"os"
)

func HandlerReset(s *State, cmd Command) error {
	err := s.Db.ClearTable(context.Background())
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	fmt.Println("reset ok")
	os.Exit(0)
	return nil
}
