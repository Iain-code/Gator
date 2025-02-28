package state

import (
	"context"
	"fmt"
)

func HandlerUsers(s *State, cmd Command) error {

	// this func will show all the users within the database and print it to console
	ctx := context.Background()
	userSlice, err := s.Db.GetUsers(ctx)
	if err != nil {
		return err
	}

	for _, user := range userSlice {
		if user.Name == s.Cfg.CurrentUserName {
			fmt.Printf("%s (current)\n", user.Name)
		} else {
			fmt.Println(user.Name)
		}
	}
	if len(userSlice) == 0 {
		fmt.Println("No users found")
	}
	return nil
}
