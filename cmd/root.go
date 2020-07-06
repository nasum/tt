package cmd

import (
	"fmt"
	"os"

	"github.com/nasum/tt/lib"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// RootCmd is root command
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

	viper.SetConfigType("json")
	err := viper.ReadInConfig()
	if err != nil {
		fmt.Fprintf(os.Stderr, "cannot read config file: %v", err)
		os.Exit(1)
	}

	consumerKey := viper.GetString("CONSUMER_KEY")
	consumerSecret := viper.GetString("CONSUMER_SECRET")
	accessToken := viper.GetString("ACCESS_TOKEN")
	accessSecret := viper.GetString("ACCESS_SECRET")
	if consumerKey == "" || consumerSecret == "" || accessToken == "" || accessSecret == "" {
		fmt.Fprintln(os.Stderr, "config file does not have authentication keys/secrets")
		os.Exit(1)
	}

	config := lib.Config{
		ConsumerKey:    consumerKey,
		ConsumerSecret: consumerSecret,
		AccessToken:    accessToken,
		AccessSecret:   accessSecret,
	}

	cobra.OnInitialize()
	RootCmd.AddCommand(
		timelineCmd(config),
		tweetCmd(config),
		listCmd(config),
	)
}
