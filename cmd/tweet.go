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
		Args:  cobra.RangeArgs(1, 1),
		RunE: func(cmd *cobra.Command, args []string) error {
			tm.Text = args[0]

			tweet, err := tm.Update()
			if err != nil {
				return err
			}

			return lib.ShowTweet(*tweet)
		},
	}

	flags := cmd.Flags()
	flags.Int64VarP(&tm.ReplyTo, "mention", "m", 0, "Set mention tweet id")

	return cmd
}

func tweet(client twitter.Client, text string, reply_to string) error {
	return nil
}
