package gitzi

import (
	"io/ioutil"
	"gopkg.in/yaml.v2"
	"log"
	"strings"
	"regexp"
	"github.com/nlopes/slack"
	"fmt"
	"os"
)

type SlackUsers struct {
	Users map[string]string
}

var slackUsers SlackUsers
var api *slack.Client

func init() {
	api = slack.New(os.Getenv("SLACK_TOKEN"))
}

func ReadSlackUsers () {
	f, err := ioutil.ReadFile("./slack_users.yaml")
	if err != nil {
		panic(err)
	}
	err = yaml.Unmarshal(f, &slackUsers)
	if err != nil {
		panic(err)
	}
}

func SendIM(to_user string, url string, comment string) {
	params := slack.PostMessageParameters{}
	attachment := slack.Attachment{
		Pretext: "",
		Text:    comment,
	}
	params.Attachments = []slack.Attachment{attachment}
	params.AsUser = true
	_, _, err := api.PostMessage(to_user, fmt.Sprint("You were mentioned at ", url), params)
	if err != nil {
		fmt.Printf("%s\n", err)
		return
	}
}

func SendIMAssigned(to_user string, url string) {
	params := slack.PostMessageParameters{}
	params.AsUser = true
	_, _, err := api.PostMessage(to_user, fmt.Sprint("You were assigned to ", url), params)
	if err != nil {
		fmt.Printf("%s\n", err)
		return
	}
}

func CreateIssueCommentSlackMessage(url string, comment string) {
	users := getAllUsers(comment)
	for user := range users {
		u := users[user]
		log.Printf(fmt.Sprintf("@%s", u))
		SendIM(fmt.Sprintf("@%s", slackUsers.Users[u]), url, comment)
	}
}

func CreateIssueSlackMessage(to_user string, url string) {
	log.Println(to_user)
	SendIMAssigned(fmt.Sprintf("@%s", slackUsers.Users[to_user]), url)
}

func getAllUsers(comment string) []string {
	var githubUsers []string
	for key := range slackUsers.Users {
		githubUsers = append(githubUsers, key)
	}
	regexString := strings.Join(githubUsers, "|")
	r, _ := regexp.Compile(regexString)
	return removeDuplicates(r.FindAllString(comment, -1))
}

func removeDuplicates(elements []string) []string {
	// Use map to record duplicates as we find them.
	encountered := map[string]bool{}
	result := []string{}

	for v := range elements {
		if encountered[elements[v]] == true {
			// Do not add duplicate.
		} else {
			// Record this element as an encountered element.
			encountered[elements[v]] = true
			// Append to result slice.
			result = append(result, elements[v])
		}
	}
	// Return the new slice.
	return result
}