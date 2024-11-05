package opengit

import (
	"context"
	"fmt"
	"github.com/asaskevich/govalidator"
	"github.com/google/go-github/v45/github"
	"github.com/summer-gonner/opengit/lang"
	"golang.org/x/oauth2"
)

type OpenGit struct {
	AccessToken string
	Language    string
	GitType     string
}

var i18n *lang.Translation

type Client struct {
	Users         *UsersService
	Repositories  *RepositoriesService
	Organizations *OrganizationsService
}

type service struct {
	client *github.Client
}

// NewClient 初始化一个客户端可以进行配置
func NewClient(og *OpenGit) (*Client, error) {

	if !govalidator.IsNull(og.Language) {
		i18n = &lang.Translation{
			Language: lang.Chinese,
		}
	} else {
		i18n = &lang.Translation{
			Language: og.Language,
		}
	}

	if govalidator.IsNull(og.AccessToken) {
		return nil, fmt.Errorf(i18n.Translate("AccessToken不能为空"))
	}
	ts := oauth2.StaticTokenSource(
		&oauth2.Token{
			AccessToken: og.AccessToken,
		})
	tc := oauth2.NewClient(context.Background(), ts)
	client := github.NewClient(tc)

	return &Client{
		Users: &UsersService{
			client: client,
		},
		Repositories: &RepositoriesService{
			client: client,
		},
		Organizations: &OrganizationsService{
			client: client,
		},
	}, nil
}
