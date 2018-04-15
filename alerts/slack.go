package alerts

import (
	"os"

	"github.com/bluele/slack"
)

var (
	token       = os.Getenv("SLACK_API_TOKEN")
	channelName = os.Getenv("SLACK_CHANNEL")
)

// Slack sends message to slack channel
func Slack(msg string) {

	api := slack.New(token)
	err := api.ChatPostMessage(channelName, msg, nil)
	if err != nil {
		panic(err)
	}
}
