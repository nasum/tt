package lib

import (
	"fmt"
	"strconv"

	"github.com/ChimeraCoder/anaconda"
	"github.com/fatih/color"
)

func ShowTweet(tweet anaconda.Tweet) {
	green := color.New(color.FgGreen).SprintFunc()
	yellow := color.New(color.FgYellow).SprintfFunc()
	cyan := color.New(color.FgCyan).SprintfFunc()
	fmt.Printf("%s\t%s\t%s\t%s\n", green(tweet.CreatedAt), yellow(strconv.FormatInt(tweet.Id, 10)), cyan("@"+tweet.User.ScreenName), tweet.Text)
}
