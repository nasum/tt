package lib

import (
	"fmt"

	"github.com/dghubble/go-twitter/twitter"
)

// TweetMethods is Tweet functions
type TweetMethods struct {
	Client  twitter.Client
	Text    string
	ReplyTo int64
}

// Update is tweet post
func (t *TweetMethods) Update() (*twitter.Tweet, error) {
	status := twitter.StatusUpdateParams{
		InReplyToStatusID: t.ReplyTo,
	}
	tweet, res, err := t.Client.Statuses.Update(t.Text, &status)

	if err != nil {
		return nil, fmt.Errorf("cannot tweet: %v: %v", err, res.Status)
	}
	return tweet, nil
}
