package main

import (
	"database/sql"
	"fmt"
	"os"

	"github.com/EndocrineMSc/gator/internal/command"
	"github.com/EndocrineMSc/gator/internal/config"
	"github.com/EndocrineMSc/gator/internal/database"
	_ "github.com/lib/pq"
)

func main() {
	state := setup()
	waitForAndRunCommand(registerCommands(), &state)
}

func setup() command.State {
	cfg, err := config.Read()
	if err != nil {
		fmt.Printf("could not read config: %v\n", err)
		os.Exit(1)
	}
	s := command.State{
		Cfg: cfg,
	}
	db, err := sql.Open("postgres", cfg.DbUrl)
	if err != nil {
		fmt.Printf("could not open database: %v\n", err)
		os.Exit(1)
	}
	dbQueries := database.New(db)
	s.Db = dbQueries

	return s
}

func registerCommands() command.Commands {
	commands := command.Commands{Cmap: make(map[string]func(*command.State, command.Command) error)}
	commands.Register("login", command.HandlerLogin)
	commands.Register("register", command.HandlerRegister)
	commands.Register("reset", command.HandlerReset)
	commands.Register("users", command.HandlerUsers)
	commands.Register("agg", command.HandlerAgg)
	commands.Register("addfeed", middlewareLoggedIn(command.HandlerAddfeed))
	commands.Register("feeds", command.HandlerFeeds)
	commands.Register("follow", middlewareLoggedIn(command.HandlerFollow))
	commands.Register("following", middlewareLoggedIn(command.HandlerFollowing))
	commands.Register("unfollow", middlewareLoggedIn(command.HandlerUnfollow))
	commands.Register("browse", middlewareLoggedIn(command.HandlerBrowse))
	return commands
}

func waitForAndRunCommand(commands command.Commands, s *command.State) {
	args := os.Args
	if len(args) < 2 {
		fmt.Println("Not enough args")
		os.Exit(1)
	}
	cmdName := args[1]
	args = args[2:]
	err := commands.Run(s, command.Command{Name: cmdName, Args: args})
	if err != nil {
		fmt.Printf("error running command: %v\n", err)
		os.Exit(1)
	}
	os.Exit(0)
}
