package cmd

import (
	"WinWallpaperSaver/utils"

	"github.com/spf13/cobra"
)

var listCmd = &cobra.Command{
	Use:     "list",
	Aliases: []string{"l", "ls"},
	Short:   "list files in dstDir",
	Run:     list,
}

func list(cmd *cobra.Command, args []string) {
	utils.ListAll(utils.GetDstDir(), 0)
}
