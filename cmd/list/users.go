package list

import (
	"fmt"

	"github.com/dghubble/go-twitter/twitter"
	"github.com/dghubble/oauth1"
	"github.com/nasum/tt/lib"
	"github.com/spf13/cobra"
)

// UsersParams is users command paramater
type UsersParams struct {
	ListID int64
}

// UsersCmd is users command
func UsersCmd(config lib.Config) *cobra.Command {

	oauthConfig := oauth1.NewConfig(config.ConsumerKey, config.ConsumerSecret)
	token := oauth1.NewToken(config.AccessToken, config.AccessSecret)

	httpClient := oauthConfig.Client(oauth1.NoContext, token)
	client := twitter.NewClient(httpClient)

	displayConsole := &lib.DisplayConsole{}

	usersParams := &UsersParams{}

	cmd := &cobra.Command{
		Use:   "users",
		Short: "get your list users",
		RunE: func(cmd *cobra.Command, args []string) error {
			var cursor int64 = -1
			for cursor != 0 {

				members, res, err := client.Lists.Members(&twitter.ListsMembersParams{ListID: usersParams.ListID, Cursor: cursor})

				if err != nil {
					return fmt.Errorf("cannot get users: %v: %v", err, res.Status)
				}

				for _, user := range members.Users {
					userTimelineParams := &twitter.UserTimelineParams{
						UserID: user.ID,
					}
					tweets, res, err := client.Timelines.UserTimeline(userTimelineParams)

					if err != nil {
						return fmt.Errorf("cannot get user-timeline: %v: %v", err, res.Status)
					}

					if len(tweets) > 0 {
						displayConsole.ShowUser(user.Name, user.ScreenName, user.URL, user.FriendsCount, user.FollowersCount, tweets[0].CreatedAt)
					} else {
						displayConsole.ShowUser(user.Name, user.ScreenName, user.URL, user.FriendsCount, user.FollowersCount, "none")
					}
				}
				cursor = members.NextCursor
			}

			return nil
		},
	}
	flags := cmd.Flags()
	flags.Int64VarP(&usersParams.ListID, "list", "l", 0, "list id")
	return cmd
}
