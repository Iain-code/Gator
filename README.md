# Gator
GOAggregator


Blog Aggregator
This is a blog aggregator application built using Go, with PostgreSQL as the database backend. It aggregates RSS feeds from various blogs and presents them in a user-friendly format.

Features
Collects and aggregates blog posts from multiple RSS feeds.
Stores blog post data in a PostgreSQL database.
Allows users to manage and update the list of subscribed feeds.
## Command words:
"login"
"register"
"users"
"feeds"
"reset"
"addfeed"
"scrape"
"follow"
"unfollow"
"following"
"remove"
"browse"

Installation
Prerequisites
Go 1.19+
PostgreSQL 13+
Git

Install dependencies:

 Make sure you have the current version of GO installed and have a local postgres database.

 Then use "go install https://github.com/Iain-code/Gator" in terminal. Make sure you are using WSL not windows terminal.

 ## Config

Create a `.gatorconfig.json` file in your home directory with the following structure:

```json
{
  "db_url": "postgres://username:@localhost:5432/database?sslmode=disable"
}
```

Replace the "username" value with your database connection string.
Find your home dir with "echo $HOME"

## Usage

Create a new user:

```bash
gator register <name>
```

Add a feed:

```bash
gator addfeed <name> <url>
```

Start the aggregator:

```bash
gator agg 30s
```

View the posts:

```bash
gator browse [limit]
```

