package lib

import (
	"fmt"
	"strconv"

	"github.com/dghubble/go-twitter/twitter"
	"github.com/fatih/color"
)

type DisplayConsole struct {
}

func (d *DisplayConsole) TimeStamp(timestamp string) string {
	green := color.New(color.FgGreen).SprintFunc()
	return green(timestamp)
}

func (d *DisplayConsole) TweetID(tweet_id string) string {
	yellow := color.New(color.FgYellow).SprintFunc()
	return yellow(tweet_id)
}

func (d *DisplayConsole) ReplyTo(tweet_id string) string {
	cyan := color.New(color.FgCyan).SprintFunc()
	return cyan(tweet_id)
}

func (d *DisplayConsole) ShowTweet(tweet twitter.Tweet) error {
	createdAt, err := tweet.CreatedAtTime()
	if err != nil {
		return err
	}
	fmt.Fprintf(color.Output, "%s\t%s\t%s\t%s\n", d.TimeStamp(createdAt.Local().Format("2006/01/02 15:04:05")), d.TweetID(strconv.FormatInt(tweet.ID, 10)), d.ReplyTo("@"+tweet.User.ScreenName), tweet.Text)
	return nil
}
