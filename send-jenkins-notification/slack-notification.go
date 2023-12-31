package main

import (
	"encoding/json"
	"fmt"
	"github.com/slack-go/slack"
	"net/http"
	"os"
)

type jenkinsBuild struct {
	BuildURL    string `json:"buildurl"`
	BuildResult string `json:"buildresult"`
	BuildNumber int    `json:"buildnumber"`
	JobName     string `json:"jobname"`
}

func sendSlackMessage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "<h1>Sent Slack Message!</h1>")
	build := jenkinsBuild{}
	err0 := json.NewDecoder(r.Body).Decode(&build)
	if err0 != nil {
		http.Error(w, err0.Error(), http.StatusBadRequest)
		return
	}

	api := slack.New(os.Getenv("SLACK_BOT_TOKEN"))
	fmt.Println(build)
	jenkinsURL := "*Build URL:* " + build.BuildURL
	buildResult := "*" + build.BuildResult + "*"
	buildNumber := "*" + fmt.Sprint(build.BuildNumber) + "*"
	jobName := "*" + build.JobName + "*"

	if buildResult == "*SUCCESS*" {
		buildResult = buildResult + " :white_check_mark:"
	} else {
		buildResult = buildResult + " :x:"
	}

	preText := "*Hello! Your Jenkins build has finished!*"
	dividerSection1 := slack.NewDividerBlock()
	jenkinsBuildDetails := jobName + " #" + buildNumber + " - " + buildResult + "\n" + jenkinsURL
	preTextField := slack.NewTextBlockObject("mrkdwn", preText+"\n\n", false, false)
	jenkinsBuildDetailsField := slack.NewTextBlockObject("mrkdwn", jenkinsBuildDetails, false, false)

	jenkinsBuildDetailsSection := slack.NewSectionBlock(jenkinsBuildDetailsField, nil, nil)
	preTextSection := slack.NewSectionBlock(preTextField, nil, nil)

	msg := slack.MsgOptionBlocks(
		preTextSection,
		dividerSection1,
		jenkinsBuildDetailsSection,
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

func main() {

	http.HandleFunc("/sendSlackMessage", sendSlackMessage)
	http.ListenAndServe(":8091", nil)

}
