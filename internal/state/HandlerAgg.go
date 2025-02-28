package state

import (
	"fmt"
	"time"
)

func HandlerAgg(s *State, cmd Command) error {

	if len(cmd.Arg) < 1 || len(cmd.Arg) > 2 {
		return fmt.Errorf("wrong amount of arguements")
	}
	duration, err := time.ParseDuration(cmd.Arg[0])
	if err != nil {
		return fmt.Errorf("invalid duration: %w", err)
	}

	fmt.Printf("Collecting feeds every %v seconds\n", duration)

	ticker := time.NewTicker(duration)
	for ; ; <-ticker.C {
		scrapeFeeds(s)
	}
}
