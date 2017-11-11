package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

func timelineCmd() *cobra.Command {
	fmt.Println("[WIP] get timeline")
	cmd := &cobra.Command{
		Use:   "timeline",
		Short: "get your timeline",
		Run: func(cmd *cobra.Command, args []string) {
			cmd.Help()
		},
	}
	return cmd
}
