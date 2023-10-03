package features

import (
	"fmt"
	"n1h41/update-google-sheet-cli/entities"
	"time"

	"github.com/go-git/go-git/v5"
	"github.com/go-git/go-git/v5/plumbing/object"
)

// Pass path to the git repo into (repoPath) and pass a bool value into (getTodaysOnly)
// whether to get all commit history from head at current branch or only just todays.
func GetCommitMessages(repoPath string, getTodaysOnly bool) []entities.CommitMessages {

	var commitMessages []entities.CommitMessages

	r, err := git.PlainOpen(repoPath)
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
