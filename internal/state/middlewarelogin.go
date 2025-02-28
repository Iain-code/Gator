package state

import (
	"context"
	"gator/internal/database"
)

func MiddlewareLoggedIn(handler func(s *State, cmd Command, user database.User) error) func(*State, Command) error {

	// takes the 3 handler functions that require a logged in user (addFeed, follow, following) and gives the func a user
	// so that each func doesnt need to get the user themselves.

	return func(s *State, cmd Command) error {
		ctx := context.Background()
		user, err := s.Db.GetUser(ctx, s.Cfg.CurrentUserName)
		if err != nil {
			return err
		}
		err = handler(s, cmd, user) // calls the handler function that was given to this function
		if err != nil {
			return err
		}
		return nil
	}
}
