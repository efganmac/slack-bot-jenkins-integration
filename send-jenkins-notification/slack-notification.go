package main

import (
	"fmt"
	"github.com/slack-go/slack"
	"os"
)

func main() {
	args := os.Args[1:]
	fmt.Println(args)

	os.Setenv("SLACK_BOT_TOKEN", "xoxb-2567467416019-6140180254420-H2HPpcBkLBPqVxiF1jiljSl8")
	api := slack.New("xoxb-2567467416019-6140180254420-H2HPpcBkLBPqVxiF1jiljSl8")
	preText := "*Hello, your Jenkins build FINISHED!*"
	jenkinsURL := "*Build URL:*  " + args[0]
	buildResult := "*" + args[1] + "*"
	buildNumber := "*" + args[2] + "*"
	jobName := "*" + args[3] + "*"

	if buildResult == "*SUCCESS*" {
		buildResult = buildResult + " :white_check_mark:"
	} else {
		buildResult = buildResult + ":x:"
	}

	dividerSection1 := slack.NewDividerBlock()
	jenkinsBuildDetails := jobName + "#" + buildNumber + "-" + buildResult + "\n" + jenkinsURL
	preTextField := slack.NewTextBlockObject("mrkdwn", preText+"\n\n", false, false)
	jenkinsBuildDetailsField := slack.NewTextBlockObject("mrkdwn", jenkinsBuildDetails, false, false)

	jenkinsBuildDetailSection := slack.NewSectionBlock(jenkinsBuildDetailsField, nil, nil)
	preTextSection := slack.NewSectionBlock(preTextField, nil, nil)

	msg := slack.MsgOptionBlocks(
		preTextSection,
		dividerSection1,
		jenkinsBuildDetailSection,
	)

	_, _, _, err := api.SendMessage(
		"C02HD32CVRN",
		msg,
	)
	if err != nil {
		fmt.Printf("%s\n", err)
		return
	}

}
