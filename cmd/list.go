package cmd

import (
	"github.com/nasum/tt/cmd/list"
	"github.com/nasum/tt/lib"
	"github.com/spf13/cobra"
)

func listCmd(config lib.Config) *cobra.Command {

	cmd := &cobra.Command{
		Use:           "list",
		Short:         "list command",
		SilenceErrors: true,
		SilenceUsage:  true,
	}

	cmd.AddCommand(
		list.LSCmd(config),
		list.UsersCmd(config),
	)

	return cmd
}
