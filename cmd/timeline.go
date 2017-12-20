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
	displayConsole := &lib.DisplayConsole{}
	timelineParams := &TimelineParams{}

	cmd := &cobra.Command{
		Use:   "timeline",
		Short: "get your timeline",
		RunE: func(cmd *cobra.Command, args []string) error {
			if timelineParams.Reply == true {
				return mentionTimeline(client, *timelineParams, displayConsole)
			} else {
				return homeTimeline(client, *timelineParams, displayConsole)
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

func homeTimeline(client twitter.Client, timelineParams TimelineParams, displayConsole *lib.DisplayConsole) error {
	homeTimelineParams := &twitter.HomeTimelineParams{
		Count:   timelineParams.Count,
		SinceID: timelineParams.SinceID,
		MaxID:   timelineParams.MaxID,
	}
	tweets, res, err := client.Timelines.HomeTimeline(homeTimelineParams)

	if err != nil {
		return fmt.Errorf("cannot get home-timeline: %v: %v", err, res.Status)
	}

	for _, v := range tweets {
		displayConsole.ShowTweet(v)
	}
	return nil
}

func mentionTimeline(client twitter.Client, timelineParams TimelineParams, displayConsole *lib.DisplayConsole) error {
	mentionTimelineParams := &twitter.MentionTimelineParams{
		Count:   timelineParams.Count,
		SinceID: timelineParams.SinceID,
		MaxID:   timelineParams.MaxID,
	}

	tweets, res, err := client.Timelines.MentionTimeline(mentionTimelineParams)

	if err != nil {
		return fmt.Errorf("cannot get mention-timeline: %v: %v", err, res.Status)
	}

	for _, v := range tweets {
		if err = displayConsole.ShowTweet(v); err != nil {
			return err
		}
	}
	return nil
}
