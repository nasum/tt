package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

func tweetCmd() *cobra.Command {
	fmt.Println("[WIP] tweet")
	cmd := &cobra.Command{
		Use:   "tweet",
		Short: "tweet",
		Run: func(cmd *cobra.Command, args []string) {
			cmd.Help()
		},
	}
	return cmd
}
