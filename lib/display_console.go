package lib

import (
	"fmt"
	"strconv"

	"github.com/dghubble/go-twitter/twitter"
	"github.com/fatih/color"
)

func ShowTweet(tweet twitter.Tweet) {
	green := color.New(color.FgGreen).SprintFunc()
	yellow := color.New(color.FgYellow).SprintfFunc()
	cyan := color.New(color.FgCyan).SprintfFunc()
	fmt.Printf("%s\t%s\t%s\t%s\n", green(tweet.CreatedAt), yellow(strconv.FormatInt(tweet.ID, 10)), cyan("@"+tweet.User.ScreenName), tweet.Text)
}
