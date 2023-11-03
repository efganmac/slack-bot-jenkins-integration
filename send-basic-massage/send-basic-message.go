package main

import (
	"fmt"
	"github.com/slack-go/slack"
	"os"
)

func main() {
	os.Setenv("SLACK_BOT_TOKEN", "xoxb-2567467416019-6140180254420-H2HPpcBkLBPqVxiF1jiljSl8")
	api := slack.New(os.Getenv("SLACK_BOT_TOKEN"))

	channelID, timeStamp, err := api.PostMessage(
		"C02HD32CVRN",
		slack.MsgOptionText("Test Message!", false),
	)
	if err != nil {
		fmt.Printf("%s\n", err)
		return
	}

	fmt.Printf("Message sent successfully to channel %s at %s ", channelID, timeStamp)
}
