package features

import (
	"fmt"
	"time"

	"github.com/go-git/go-git/v5"
	"github.com/go-git/go-git/v5/plumbing/object"
)

type CommitMessages struct {
  Message string
  DateTime time.Time
}

func GetCommitMessages() {
	// WARN: For testing purposes only
	// repo_absolute_path := "/home/n1h41/development/Go/cli/update-google-sheet-cli"
	repo_absolute_path := "/home/n1h41/development/Node/chat-api"

	r, err := git.PlainOpen(repo_absolute_path)
	if err != nil {
		panic(err)
	}

	head, err := r.Head()
	if err != nil {
		panic(err)
	}

	logs, err := r.Log(&git.LogOptions{
		From: head.Hash(),
		All:  true,
	})

	if err != nil {
		panic(err)
	}

	// Iterate commits
	logs.ForEach(func(c *object.Commit) error {
    fmt.Println(c.Author.When.Local().Format("02/01/2006"))
		return nil
	})
}
