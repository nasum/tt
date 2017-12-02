package lib

import (
	"fmt"
	"net/url"

	"github.com/ChimeraCoder/anaconda"
)

type TweetMethods struct {
	Api     anaconda.TwitterApi
	Text    string
	ReplyTo int64
}

func (t *TweetMethods) Update() *anaconda.Tweet {
	v := url.Values{}
	v.Set("in_reply_to_status_id", fmt.Sprint(t.ReplyTo))

	tweet, err := t.Api.PostTweet(t.Text, v)

	if err != nil {
		fmt.Println(err)
		return nil
	}
	return &tweet
}
