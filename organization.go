package opengit

import (
	"context"
	"fmt"
	"github.com/google/go-github/v45/github"
	"log"
)

type Organization struct {
}

type OrganizationsService service

func (os OrganizationsService) List(username string) ([]*github.Organization, error) {
	list, response, _ := os.client.Organizations.List(context.Background(), username, nil)
	log.Println(response.Status)
	switch response.Response.Status {
	case Unauthorized:
		return nil, fmt.Errorf("unauthorized")
	case Success:
		return list, nil
	}
	return nil, nil
}
