package list

import (
	"fmt"

	"github.com/dghubble/go-twitter/twitter"
	"github.com/dghubble/oauth1"
	"github.com/nasum/tt/lib"
	"github.com/spf13/cobra"
)

// LSCmd is show my list
func LSCmd(config lib.Config) *cobra.Command {

	oauthConfig := oauth1.NewConfig(config.ConsumerKey, config.ConsumerSecret)
	token := oauth1.NewToken(config.AccessToken, config.AccessSecret)

	httpClient := oauthConfig.Client(oauth1.NoContext, token)
	client := twitter.NewClient(httpClient)

	displayConsole := &lib.DisplayConsole{}

	cmd := &cobra.Command{
		Use:   "ls",
		Short: "get your list",
		RunE: func(cmd *cobra.Command, args []string) error {
			user, res, err := client.Accounts.VerifyCredentials(&twitter.AccountVerifyParams{})

			if err != nil {
				return fmt.Errorf("cannot get VerifyCredentials: %v: %v", err, res.Status)
			}

			lists, res, err := client.Lists.List(&twitter.ListsListParams{
				UserID: user.ID,
			})

			if err != nil {
				return fmt.Errorf("cannot get list: %v: %v", err, res.Status)
			}

			for _, list := range lists {
				displayConsole.ShowList(list.Name, list.URI, list.ID)
			}

			return nil
		},
	}
	return cmd
}
