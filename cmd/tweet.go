package cmd

import (
	"github.com/ChimeraCoder/anaconda"
	"github.com/dghubble/go-twitter/twitter"
	"github.com/nasum/tt/lib"
	"github.com/spf13/cobra"
)

func tweetCmd(api anaconda.TwitterApi) *cobra.Command {
	tm := lib.TweetMethods{Api: api}
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
	flags.Int64VarP(&tm.ReplyTo, "mention", "m", 0, "Set mention tweet id")

	return cmd
}

func tweet(client twitter.Client, text string, reply_to string) {

}
