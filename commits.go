package opengit

import (
	"context"
	"fmt"
	"github.com/google/go-github/v45/github"
	"log"
)

// CompareCommits 比较Commits信息
func (rs RepositoriesService) CompareCommits(username string, repoName string, baseBranch string, compareBranch string) (*github.CommitsComparison, error) {
	compareResult, response, _ := rs.client.Repositories.CompareCommits(context.Background(), username, repoName, baseBranch, compareBranch, nil)
	log.Println(response.Status)
	switch response.Response.Status {
	case Unauthorized:
		return nil, fmt.Errorf("unauthorized")
	case Success:
		return compareResult, nil
	}
	return nil, nil
}
