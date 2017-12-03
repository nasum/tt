package cmd

import (
	"fmt"

	"github.com/ChimeraCoder/anaconda"
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

	anaconda.SetConsumerKey(consumerKey)
	anaconda.SetConsumerSecret(consumerSecret)
	api := anaconda.NewTwitterApi(accessToken, accessSecret)

	cobra.OnInitialize()
	RootCmd.AddCommand(
		timelineCmd(*api),
		tweetCmd(*api),
	)
}
