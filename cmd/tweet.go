package cmd

import (
	"github.com/dghubble/go-twitter/twitter"
	"github.com/nasum/tt/lib"
	"github.com/spf13/cobra"
)

func tweetCmd(client twitter.Client) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "tweet",
		Short: "post your tweet",
		Run: func(cmd *cobra.Command, args []string) {

			if len(args) == 0 {
				cmd.Help()
				return
			}

			tm := lib.TweetMethods{Client: client, Text: args[0]}
			tweet := tm.Update()

			lib.ShowTweet(*tweet)
		},
	}
	return cmd
}
