package cmd

import (
	"fmt"

	"github.com/dghubble/go-twitter/twitter"
	"github.com/nasum/tt/lib"
	"github.com/spf13/cobra"
)

type TimelineParams struct {
	Count   int
	SinceID int64
	MaxID   int64
	Reply   bool
}

func timelineCmd(client twitter.Client) *cobra.Command {
	timelineParams := &TimelineParams{}

	cmd := &cobra.Command{
		Use:   "timeline",
		Short: "get your timeline",
		Run: func(cmd *cobra.Command, args []string) {
			if timelineParams.Reply == true {
				mentionTimeline(client, *timelineParams)
			} else {
				homeTimeline(client, *timelineParams)
			}
		},
	}

	flags := cmd.Flags()
	flags.IntVarP(&timelineParams.Count, "count", "c", 20, "Set tweet count")
	flags.Int64VarP(&timelineParams.SinceID, "since-id", "S", 0, "Set since tweet id")
	flags.Int64VarP(&timelineParams.MaxID, "max-id", "M", 0, "Set max tweet id")
	flags.BoolVarP(&timelineParams.Reply, "mention", "m", false, "show mention timeline")

	return cmd
}

func homeTimeline(client twitter.Client, timelineParams TimelineParams) {
	homeTimelineParams := &twitter.HomeTimelineParams{
		Count:   timelineParams.Count,
		SinceID: timelineParams.SinceID,
		MaxID:   timelineParams.MaxID,
	}
	tweets, res, err := client.Timelines.HomeTimeline(homeTimelineParams)

	if err != nil {
		fmt.Println(res)
	}

	for _, v := range tweets {
		lib.ShowTweet(v)
	}
}

func mentionTimeline(client twitter.Client, timelineParams TimelineParams) {
	mentionTimelineParams := &twitter.MentionTimelineParams{
		Count:   timelineParams.Count,
		SinceID: timelineParams.SinceID,
		MaxID:   timelineParams.MaxID,
	}

	tweets, res, err := client.Timelines.MentionTimeline(mentionTimelineParams)

	if err != nil {
		fmt.Println(res)
	}

	for _, v := range tweets {
		lib.ShowTweet(v)
	}
}
