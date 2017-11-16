package cmd

import (
	"fmt"

	"github.com/dghubble/go-twitter/twitter"
	"github.com/spf13/cobra"
)

func tweetCmd(client twitter.Client) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "tweet",
		Short: "tweet",
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println("[WIP] tweet")
			cmd.Help()
		},
	}
	return cmd
}
