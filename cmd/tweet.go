package cmd

import (
	"strings"

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
		RunE: func(cmd *cobra.Command, args []string) error {
			tm.Text = strings.Join(args, " ")

			tweet, err := tm.Update()
			if err != nil {
				return err
			}

			err = displayConsole.ShowTweet(*tweet, true)
			if err != nil {
				return err
			}

			return nil
		},
	}

	flags := cmd.Flags()
	flags.Int64VarP(&tm.ReplyTo, "mention", "m", 0, "Set mention tweet id")

	return cmd
}
