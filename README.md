# Gator

## Requirements
This project requires both 

- [PostgreSQL](https://www.postgresql.org/)
- [Go](https://go.dev/doc/install)

On Linux: `sudo apt install postgresql postgresql-contrib`

For Go, on Linux-based systems you can find quick instructions also on https://webinstall.dev/golang/

## Installation
```bash
go install github.com/finfreezer/gator@latest
```

## Configuration
Set up a configuration file in your home directory called .gatorconfig.json that includes the database url you wish to use, the program
will also use this to track the logger-in user under the hood.
```json
{
  "db_url": "postgres://username:password@localhost:5432/gator?sslmode=disable"
}
```

## Usage

### Getting Started
To begin using the program, you can register a new user with `gator register {username}`.

Or change registered users with `gator login {username}`.

### Using the program
Once you've registered at least one user, you can add an RSS feed with `gator addfeed {name} {feed_url}`, which is automatically
added into the current logged-in user's followed feeds.

By using `gator agg {[number]{s|m|h}}`, the program will begin pulling articles from the feeds every time
the designated amount of time has passed, e.g. "gator agg 15s" for 15 seconds.

By using `gator browse {number}` you can see the latest `{number}` articles in your feed.
By omitting the number, the program displays the freshest 2 articles by default.

By using `gator follow {url}`, you can add an existing feed to be followed by the current user.
Do note that the program automatically does this for any user when they use "addfeed".
