package main

import (
	"fmt"
	"os"

	"github.com/slack-go/slack"
)

func main() {
	os.Setenv("SLACK_BOT_TOKEN", "xoxb-3428851110421-3444679023617-l2jLbEaeE1E0GuVxgSbjfM7r")
	os.Setenv("CHANNEL_ID", "C03C944AWBZ")
	api := slack.New(os.Getenv("SLACK_BOT_TOKEN"))
	channelArr := []string{os.Getenv("CHANNEL_ID")}
	fileArr := []string{"abc.pdf"}

	for i := 0; i < len(fileArr); i++ {
		params := slack.FileUploadParameters{
			Channels: channelArr, File: fileArr[i],
		}
		file, err := api.UploadFile(params)
		if err != nil {
			fmt.Printf("%s \n", err)
			return
		}
		fmt.Printf("Name: %s , URL : %s", file.Name, file.URLPrivate)
	}
}
