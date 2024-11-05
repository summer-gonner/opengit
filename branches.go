package opengit

import (
	"context"
	"fmt"
	"github.com/google/go-github/v45/github"
	"log"
)

type Branch struct {
}

type BranchesService service

// ListBranches 获取分支列表
func (rs RepositoriesService) ListBranches(username string, repoName string) ([]*github.Branch, error) {
	branches, response, _ := rs.client.Repositories.ListBranches(context.Background(), username, repoName, nil)
	log.Println(response.Status)
	switch response.Response.Status {
	case Unauthorized:
		return nil, fmt.Errorf("unauthorized")
	case Success:
		return branches, nil
	}
	return nil, nil
}

// GetBranch 获取分支信息
func (rs RepositoriesService) GetBranch(username string, repoName string, branchName string) (*github.Branch, error) {
	branch, response, _ := rs.client.Repositories.GetBranch(context.Background(), username, repoName, branchName, true)
	log.Println(response.Status)
	switch response.Response.Status {
	case Unauthorized:
		return nil, fmt.Errorf("unauthorized")
	case Success:
		return branch, nil
	}
	return nil, nil
}

// RenameBranch 重命名分支信息
func (rs RepositoriesService) RenameBranch(username string, repoName string, branchName string, newName string) (*github.Branch, error) {
	branch, response, _ := rs.client.Repositories.RenameBranch(context.Background(), username, repoName, branchName, newName)
	log.Println(response.Status)
	switch response.Response.Status {
	case Unauthorized:
		return nil, fmt.Errorf("unauthorized")
	case Success:
		return branch, nil
	}
	return nil, nil
}
func (rs RepositoriesService) Merge(username string, repoName string, baseBranch string, headBranch string) (*github.RepositoryCommit, error) {
	mergeRequest := &github.RepositoryMergeRequest{
		Base:          github.String(baseBranch),                       // 目标分支
		Head:          github.String(headBranch),                       // 源分支
		CommitMessage: github.String("Merge feature branch into main"), // 合并的提交消息
	}
	repositoryCommitResult, response, _ := rs.client.Repositories.Merge(context.Background(), username, repoName, mergeRequest)
	log.Println(response.Status)
	switch response.Response.Status {
	case Unauthorized:
		return nil, fmt.Errorf("unauthorized")
	case Success:
		return repositoryCommitResult, nil
	}
	return nil, nil
}
