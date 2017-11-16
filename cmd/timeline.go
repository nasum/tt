package cmd

import (
	"fmt"

	"github.com/dghubble/go-twitter/twitter"
	"github.com/spf13/cobra"
)

func timelineCmd(client twitter.Client) *cobra.Command {

	cmd := &cobra.Command{
		Use:   "timeline",
		Short: "get your timeline",
		Run: func(cmd *cobra.Command, args []string) {
			tweets, res, err := client.Timelines.HomeTimeline(&twitter.HomeTimelineParams{
				Count: 20,
			})

			if err != nil {
				fmt.Println(res)
			}

			for _, v := range tweets {
				fmt.Print(v)
			}
		},
	}
	return cmd
}
