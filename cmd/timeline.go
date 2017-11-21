package cmd

import (
	"fmt"

	"github.com/dghubble/go-twitter/twitter"
	"github.com/nasum/tt/lib"
	"github.com/spf13/cobra"
)

func timelineCmd(client twitter.Client) *cobra.Command {
	homeTimelineParams := &twitter.HomeTimelineParams{}

	cmd := &cobra.Command{
		Use:   "timeline",
		Short: "get your timeline",
		Run: func(cmd *cobra.Command, args []string) {
			tweets, res, err := client.Timelines.HomeTimeline(homeTimelineParams)

			if err != nil {
				fmt.Println(res)
			}

			for _, v := range tweets {
				lib.ShowTweet(v)
			}
		},
	}

	flags := cmd.Flags()
	flags.IntVarP(&homeTimelineParams.Count, "count", "c", 20, "Set tweet count")
	flags.Int64VarP(&homeTimelineParams.SinceID, "since-id", "s", 0, "Set since tweet id")
	flags.Int64VarP(&homeTimelineParams.MaxID, "max-id", "m", 0, "Set max tweet id")

	return cmd
}
