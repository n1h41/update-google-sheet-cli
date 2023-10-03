package features

import (
	"fmt"
	"n1h41/update-google-sheet-cli/entities"
	"time"

	"github.com/go-git/go-git/v5"
	"github.com/go-git/go-git/v5/plumbing/object"
)

func GetCommitMessages(repoPath string, getTodaysOnly bool) []entities.CommitMessages {
	// WARN: For testing purposes only
	// repo_absolute_path := "/home/n1h41/development/Go/cli/update-google-sheet-cli"
	repo_absolute_path := "/mnt/d/nihal/Development/Flutter/works/raf-pharmacy/"

	var commitMessages []entities.CommitMessages

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

	// Iterate commits and get messages and time
	if err := logs.ForEach(func(c *object.Commit) error {
		if c.Author.When.Local().Format("02/01/2006") == time.Now().Format("02/01/2006") && getTodaysOnly {
			commitMessages = append(commitMessages, entities.CommitMessages{
				Message:  c.Message,
				DateTime: c.Author.When.Local(),
			})
			fmt.Println("Entry present for today")
			return nil
		}
		commitMessages = append(commitMessages, entities.CommitMessages{
			Message:  c.Message,
			DateTime: c.Author.When.Local(),
		})
		return nil
	}); err != nil {
		panic(err)
	}

	return commitMessages
}
