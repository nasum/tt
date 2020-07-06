package lib

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/dghubble/go-twitter/twitter"
	"github.com/fatih/color"
)

const (
	// TwitterURL is twitter home url
	TwitterURL string = "https://twitter.com"
)

// DisplayConsole is struct
type DisplayConsole struct {
}

// TimeStamp is output colored timestamp text
func (d *DisplayConsole) TimeStamp(timestamp string) string {
	green := color.New(color.FgGreen).SprintFunc()
	return green(timestamp)
}

// TweetID is output colored tweet id text
func (d *DisplayConsole) TweetID(tweetID string) string {
	yellow := color.New(color.FgYellow).SprintFunc()
	return yellow(tweetID)
}

// ListID is output colored tweet id text
func (d *DisplayConsole) ListID(listID string) string {
	yellow := color.New(color.FgYellow).SprintFunc()
	return yellow(listID)
}

// ReplyTo is output colored reply text
func (d *DisplayConsole) ReplyTo(tweetID string) string {
	cyan := color.New(color.FgCyan).SprintFunc()
	return cyan(tweetID)
}

// URL is output colored url text
func (d *DisplayConsole) URL(url string) string {
	green := color.New(color.FgGreen).SprintFunc()
	return green(url)
}

// Name is output colored name text
func (d *DisplayConsole) Name(name string) string {
	yellow := color.New(color.FgYellow).SprintFunc()
	return yellow(name)
}

// ShowTweet is display tweet text
func (d *DisplayConsole) ShowTweet(tweet twitter.Tweet, showImage bool) error {
	createdAt, err := tweet.CreatedAtTime()

	if err != nil {
		return err
	}

	fmt.Fprintf(
		color.Output,
		"%s\t%s\t%s\t%s\n%s\n",
		d.TimeStamp(createdAt.Local().Format("2006/01/02 15:04:05")),
		d.TweetID(tweet.IDStr),
		d.ReplyTo("@"+tweet.User.ScreenName),
		d.URL(d.createTweetURL(tweet)),
		tweet.Text,
	)

	if showImage == false {
		return nil
	}

	media := &Media{}

	medias := tweet.Entities.Media
	for _, mediaEntity := range medias {
		if mediaEntity.Type == "photo" {
			media.ShowImage(mediaEntity.MediaURLHttps + ":thumb")
			fmt.Println()
		}
	}

	return nil
}

// ShowList is display list text
func (d *DisplayConsole) ShowList(title string, url string, id int64) {
	fmt.Fprintf(
		color.Output,
		"%s\t%s\t%s\n",
		d.URL("https://twitter.com/"+url),
		d.ListID(strconv.FormatInt(id, 10)),
		title,
	)
}

// ShowUser is display user
func (d *DisplayConsole) ShowUser(name string, screenName string, url string, friendsCount int, followersCount int, lastUpdatedAt string) {
	fmt.Fprintf(
		color.Output,
		"%s\t%s\t%s\t%s\t%s\t%s\n",
		d.URL("https://twitter.com/"+screenName),
		d.Name(name),
		d.Name("@"+screenName),
		strconv.Itoa(friendsCount),
		strconv.Itoa(followersCount),
		lastUpdatedAt,
	)
}

func (d *DisplayConsole) createTweetURL(tweet twitter.Tweet) string {
	array := []string{TwitterURL, tweet.User.ScreenName, "status", tweet.IDStr}
	return strings.Join(array, "/")
}
