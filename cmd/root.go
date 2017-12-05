package cmd

import (
	"fmt"

	"github.com/dghubble/go-twitter/twitter"
	"github.com/dghubble/oauth1"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var RootCmd = &cobra.Command{
	Use:           "tt",
	Short:         "Twitter Client",
	SilenceErrors: true,
	SilenceUsage:  true,
}

func init() {
	viper.SetConfigName("ttrc")
	viper.AddConfigPath("./")
	viper.AddConfigPath("$HOME/")

	err := viper.ReadInConfig()
	if err != nil {
		panic(fmt.Errorf("Fatal error config file: %s \n", err))
	}

	consumerKey := viper.GetString("CONSUMER_KEY")
	consumerSecret := viper.GetString("CONSUMER_SECRET")
	accessToken := viper.GetString("ACCESS_TOKEN")
	accessSecret := viper.GetString("ACCESS_SECRET")
	if consumerKey == "" || consumerSecret == "" || accessToken == "" || accessSecret == "" {
		panic(fmt.Errorf("Config file does not have authentication keys/secrets"))
	}

	config := oauth1.NewConfig(consumerKey, consumerSecret)
	token := oauth1.NewToken(accessToken, accessSecret)

	httpClient := config.Client(oauth1.NoContext, token)
	client := twitter.NewClient(httpClient)

	cobra.OnInitialize()
	RootCmd.AddCommand(
		timelineCmd(*client),
		tweetCmd(*client),
	)
}
