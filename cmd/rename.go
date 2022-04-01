package cmd

import (
	"WinWallpaperSaver/utils"

	"github.com/spf13/cobra"
)

var renameCmd = &cobra.Command{
	Use:     "rename",
	Aliases: []string{"r"},
	Short:   "rename",
	Run:     rename,
}

func rename(cmd *cobra.Command, args []string) {
	utils.RenameToPNG(utils.GetDstDir())
}
