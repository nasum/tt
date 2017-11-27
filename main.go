package main

import (
	"fmt"
	"os"

	"github.com/nasum/tt/cmd"
	"github.com/spf13/viper"
)

func main() {
	viper.SetConfigName("ttrc")
	viper.AddConfigPath("./")
	viper.AddConfigPath("$HOME/")

	err := viper.ReadInConfig()
	if err != nil {
		panic(fmt.Errorf("Fatal error config file: %s \n", err))
	}

	if err := cmd.RootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(-1)
	}
}
