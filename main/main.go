package main

import (
	"github.com/summer-gonner/opengit"
	"log"
)

var (
	owner      = "summer-gonner"
	repoName   = "opengit"
	branchName = "main"
	token      = ""
)

func main() {
	g := &opengit.OpenGit{AccessToken: token}
	newClient, err := opengit.NewClient(g)
	if err != nil {
		log.Printf(err.Error())
	}
	//user, err := newClient.Users.GetUser(owner)
	//log.Printf("user%s\n", user)

	//_, err = newClient.Repositories.List(owner)
	//log.Printf("err%s\n", err)
	//organizations, err := newClient.Organizations.List(owner)
	//log.Printf("organizations%s\n", organizations)
	branches, err := newClient.Repositories.ListBranches(owner, repoName)
	for _, branch := range branches {
		log.Printf("branch%s\n", branch.GetCommit())
	}

}
