package cmd

import (
	"fmt"

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

			tweet, res, err := client.Statuses.Update(args[0], nil)

			if err != nil {
				fmt.Println(res)
				return
			}

			lib.ShowTweet(*tweet)
		},
	}
	return cmd
}
