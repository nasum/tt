package lib

import (
	"fmt"

	"github.com/dghubble/go-twitter/twitter"
)

type TweetMethods struct {
	Client twitter.Client
	Text   string
}

func (t *TweetMethods) Update() *twitter.Tweet {
	tweet, res, err := t.Client.Statuses.Update(t.Text, nil)

	if err != nil {
		fmt.Println(res)
		return nil
	}
	return tweet
}
