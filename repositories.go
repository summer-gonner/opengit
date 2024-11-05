package opengit

import (
	"context"
	"fmt"
	"github.com/google/go-github/v45/github"
)

type Repository struct {
	Branch *BranchesService
}

type RepositoriesService service

func (rs RepositoriesService) List(username string) ([]*github.Repository, error) {
	repositories, response, _ := rs.client.Repositories.List(context.Background(), username, nil)
	//if err != nil {
	//	return nil, err
	//}
	//log.Printf("1111111%s", response.Status)
	switch response.Response.Status {
	case Unauthorized:
		return nil, fmt.Errorf("unauthorized")
	case Success:
		return repositories, nil
	case NotFOUND:
		return nil, fmt.Errorf(i18n.Translate("仓库列表为空!"))
	}
	return nil, nil
}
func (rs RepositoriesService) Get(username string, repositoryName string) (*github.Repository, error) {
	repository, response, _ := rs.client.Repositories.Get(context.Background(), username, repositoryName)
	switch response.Response.Status {
	case Unauthorized:
		return nil, fmt.Errorf("unauthorized")
	case Success:
		return repository, nil
	case NotFOUND:
		return nil, fmt.Errorf(i18n.Translate("仓库不存在!"))
	}
	return nil, nil
}
func (rs RepositoriesService) Create(org string) (*github.Repository, error) {
	repository, response, _ := rs.client.Repositories.Create(context.Background(), org, &github.Repository{})
	switch response.Response.Status {
	case Unauthorized:
		return nil, fmt.Errorf("unauthorized")
	case Created:
		return repository, nil
	case NotFOUND:
		return nil, fmt.Errorf(i18n.Translate("仓库不存在!"))
	}
	return nil, nil
}
