package cmd

import (
	"github.com/dghubble/go-twitter/twitter"
	"github.com/nasum/tt/lib"
	"github.com/spf13/cobra"
)

func tweetCmd(client twitter.Client) *cobra.Command {
	tm := lib.TweetMethods{Client: client}
	cmd := &cobra.Command{
		Use:   "tweet",
		Short: "post your tweet",
		Run: func(cmd *cobra.Command, args []string) {

			if len(args) == 0 {
				cmd.Help()
				return
			}

			tm.Text = args[0]

			tweet := tm.Update()

			lib.ShowTweet(*tweet)
		},
	}

	flags := cmd.Flags()
	flags.Int64VarP(&tm.ReplyTo, "reply", "r", 0, "Set reply tweet id")

	return cmd
}

func tweet(client twitter.Client, text string, reply_to string) {

}
