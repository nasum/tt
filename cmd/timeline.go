package cmd

import (
	"fmt"
	"net/url"

	"github.com/ChimeraCoder/anaconda"
	"github.com/nasum/tt/lib"
	"github.com/spf13/cobra"
)

type TimelineParams struct {
	Count   int
	SinceID int64
	MaxID   int64
	Reply   bool
}

func timelineCmd(api anaconda.TwitterApi) *cobra.Command {
	timelineParams := &TimelineParams{}

	cmd := &cobra.Command{
		Use:   "timeline",
		Short: "get your timeline",
		Run: func(cmd *cobra.Command, args []string) {
			if timelineParams.Reply == true {
				mentionTimeline(api, *timelineParams)
			} else {
				homeTimeline(api, *timelineParams)
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

func homeTimeline(api anaconda.TwitterApi, timelineParams TimelineParams) {
	v := buildParams(timelineParams)

	tweets, err := api.GetHomeTimeline(v)

	if err != nil {
		fmt.Println(err)
	}

	for _, v := range tweets {
		lib.ShowTweet(v)
	}
}

func mentionTimeline(api anaconda.TwitterApi, timelineParams TimelineParams) {
	v := buildParams(timelineParams)

	tweets, err := api.GetMentionsTimeline(v)

	if err != nil {
		fmt.Println(err)
	}

	for _, v := range tweets {
		lib.ShowTweet(v)
	}
}

func buildParams(timelineParams TimelineParams) url.Values {
	v := url.Values{}
	v.Set("count", fmt.Sprint(timelineParams.Count))
	if timelineParams.SinceID > 0 {
		v.Set("since_id", fmt.Sprint(timelineParams.SinceID))
	}
	if timelineParams.MaxID > 0 {
		v.Set("max_id", fmt.Sprint(timelineParams.MaxID))
	}
	return v
}
