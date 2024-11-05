package opengit

import (
	"context"
	"fmt"
	"github.com/google/go-github/v45/github"
	"log"
)

type User struct {
	Name string
}
type UsersService service

func (u UsersService) GetUser(name string) (*github.User, error) {
	user, response, err := u.client.Users.Get(context.Background(), name)

	if err != nil {
		return nil, err
	}

	log.Printf("Response status: %s\n", response.Response.Status)

	switch response.Response.Status {
	case Unauthorized:
		return nil, fmt.Errorf("unauthorized")
	case Created:
		return user, nil
	}
	return nil, nil

}
