package gitzi

import (
	"github.com/google/go-github/github"
	"encoding/json"
	"log"
	"os"
)

type Hook struct {
	Event string
	Payload []byte
}

var secretGithub = os.Getenv("GITHUB_SECRET")

func ghIssueComment(hook *Hook) {
	var issueGithub github.IssueCommentEvent
	var issueComment, issueUrl string
	err := json.Unmarshal(hook.Payload, &issueGithub)
	if err != nil {
		panic(err)
	}

	issueComment = *issueGithub.Comment.Body
	issueUrl = *issueGithub.Comment.HTMLURL
	log.Println(issueComment, issueUrl)
	CreateIssueCommentSlackMessage(issueUrl, issueComment)
}

func ghIssue(hook *Hook) {
	var issueEventGithub github.IssuesEvent
	var issueAssignee github.User
	var issueURL, issueAction string

	err := json.Unmarshal(hook.Payload, &issueEventGithub)
	if err != nil {
		panic(err)
	}

	issueAction = *issueEventGithub.Action
	if issueAction != "assigned" {
		return
	}

	issueAssignee = *issueEventGithub.Assignee
	issueURL = *issueEventGithub.Issue.HTMLURL
	CreateIssueSlackMessage(*issueAssignee.Login, issueURL)
}