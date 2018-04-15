package alerts

import (
	"os"

	"github.com/bluele/slack"
)

var (
	token       = os.Getenv("SLACK_API_TOKEN")
	channelName = os.Getenv("SLACK_CHANNEL")
)

func main() {

	api := slack.New(token)
	err := api.ChatPostMessage(channelName, "Hello, world!", nil)
	if err != nil {
		panic(err)
	}
}
