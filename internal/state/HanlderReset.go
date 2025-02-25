package state

import (
	"context"
	"fmt"
)

func HandlerReset(s *State, cmd Command) error {

	ctx := context.Background()
	err := s.Db.DeleteAllUsers(ctx)
	// s.Db is a struct that contains all the QUERIES made by "sqlc generate"

	if err != nil {
		return err
	}
	fmt.Println("All users were deleted")
	return nil
}
