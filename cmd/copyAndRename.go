package cmd

import "github.com/spf13/cobra"

var copyAndRenameCmd = &cobra.Command{
	Use:     "copyAndRename",
	Aliases: []string{""},
	Short:   "copy and rename",
	Run:     copyAndRename,
}

func copyAndRename(cmd *cobra.Command, args []string) {
	copy(cmd, args)
	rename(cmd, args)
}
