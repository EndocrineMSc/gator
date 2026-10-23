# gator

A small command-line RSS feed aggregator written in Go. Register a user, add and follow RSS feeds, let the aggregator collect posts in the background, then browse the latest posts from the feeds you follow.

## Requirements

You'll need these installed to run gator:

- [Go](https://go.dev/doc/install) 1.27 or newer
- [PostgreSQL](https://www.postgresql.org/download/) 15 or newer
- [goose](https://github.com/pressly/goose) to run the database migrations:

  ```sh
  go install github.com/pressly/goose/v3/cmd/goose@latest
  ```

## Installation

Install the `gator` binary with `go install`:

```sh
go install github.com/EndocrineMSc/gator@latest
```

This puts `gator` in `$(go env GOPATH)/bin` (usually `~/go/bin`). Make sure that directory is on your `PATH`.

## Setup

### 1. Create the database

Start Postgres and create a database for gator:

```sh
psql -U postgres -c "CREATE DATABASE gator;"
```

### 2. Run the migrations

From a clone of this repository, apply the schema in `sql/schema` with goose:

```sh
goose -dir sql/schema postgres "postgres://postgres:postgres@localhost:5432/gator" up
```

Adjust the username, password, host and port to match your Postgres setup.

### 3. Create the config file

gator reads its settings from `~/.gatorconfig.json`. Create it with your connection string:

```json
{
  "db_url": "postgres://postgres:postgres@localhost:5432/gator?sslmode=disable"
}
```

gator adds `current_user_name` to this file itself when you register or log in.

## Usage

```sh
gator <command> [args...]
```

### Getting started

```sh
gator register alice                                   # create a user and log in as them
gator addfeed "Hacker News" https://news.ycombinator.com/rss   # add a feed and follow it
gator agg 1m                                           # fetch posts every minute (Ctrl+C to stop)
gator browse 5                                         # show the 5 newest posts from your feeds
```

### Commands

| Command | Description |
| --- | --- |
| `register <name>` | Create a new user and log in as them |
| `login <name>` | Switch to an existing user |
| `users` | List all users and mark the current one |
| `addfeed <name> <url>` | Add a feed and follow it as the current user |
| `feeds` | List all feeds in the database |
| `follow <url>` | Follow an existing feed |
| `following` | List the feeds the current user follows |
| `unfollow <url>` | Stop following a feed |
| `agg <interval>` | Fetch feeds continuously, e.g. `30s`, `1m`, `1h` |
| `browse [limit]` | Show the newest posts from followed feeds (default 2) |
| `reset` | Delete all users and their data — handy for testing |
