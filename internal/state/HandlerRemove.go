package state

import (
	"context"
	"fmt"
)

func HandlerRemove(s *State, cmd Command) error {

	ctx := context.Background()
	err := s.Db.DeleteFeed(ctx, cmd.Arg[0])
	if err != nil {
		fmt.Println("unable to remove feed")
		return err
	}
	fmt.Printf("feed url: %v has been removed\n", cmd.Arg[0])
	return nil
}
