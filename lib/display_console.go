package lib

import (
	"fmt"
	"strconv"

	"github.com/dghubble/go-twitter/twitter"
	"github.com/fatih/color"
)

func ShowTweet(tweet twitter.Tweet) error {
	green := color.New(color.FgGreen).SprintFunc()
	yellow := color.New(color.FgYellow).SprintfFunc()
	cyan := color.New(color.FgCyan).SprintfFunc()
	createdAt, err := tweet.CreatedAtTime()
	if err != nil {
		return err
	}
	fmt.Fprintf(color.Output, "%s\t%s\t%s\t%s\n", green(createdAt.Local().Format("2006/01/02 15:04:05")), yellow(strconv.FormatInt(tweet.ID, 10)), cyan("@"+tweet.User.ScreenName), tweet.Text)
	return nil
}
