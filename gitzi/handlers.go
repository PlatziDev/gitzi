package gitzi

import (
	"net/http"
	"github.com/google/go-github/github"
)

func GHWebhook(rw http.ResponseWriter, req *http.Request) {
	// X-GitHub-Delivery
	// X-Hub-Signature
	hook := Hook{}
	hook.Event = req.Header.Get("X-GitHub-Event")

	payload, err := github.ValidatePayload(req, []byte(secretGithub))

	if err != nil {
		http.Error(rw, err.Error(), http.StatusInternalServerError)
		return
	}

	hook.Payload = payload

	switch hook.Event {
	case "issue_comment":
		ghIssueComment(&hook)
	case "issues":
		ghIssue(&hook)
	default:
		panic("Event doesn't exist")
	}
}

