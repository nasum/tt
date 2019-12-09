package cmd

import (
	"github.com/dghubble/go-twitter/twitter"
	"github.com/dghubble/oauth1"
	"github.com/nasum/tt/lib"
	"github.com/spf13/cobra"
)

func tweetCmd(config lib.Config) *cobra.Command {
	oauthConfig := oauth1.NewConfig(config.ConsumerKey, config.ConsumerSecret)
	token := oauth1.NewToken(config.AccessToken, config.AccessSecret)

	httpClient := oauthConfig.Client(oauth1.NoContext, token)
	client := twitter.NewClient(httpClient)

	displayConsole := &lib.DisplayConsole{}
	tm := lib.TweetMethods{Client: *client}
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

			createdAt, err := tweet.CreatedAtTime()
			if err != nil {
				return err
			}
			displayConsole.ShowTweet(createdAt, tweet.ID, tweet.User.ScreenName, tweet.FullText)
			return nil
		},
	}

	flags := cmd.Flags()
	flags.Int64VarP(&tm.ReplyTo, "mention", "m", 0, "Set mention tweet id")

	return cmd
}

func tweet(client twitter.Client, text string, reply_to string) error {
	return nil
}
