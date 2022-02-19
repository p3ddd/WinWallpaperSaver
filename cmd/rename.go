package cmd

import (
	"Win10WallpaperSaver/utils"
	"path"

	"github.com/spf13/cobra"
)

var renameCmd = &cobra.Command{
	Use:     "rename",
	Aliases: []string{""},
	Short:   "create default template", // 初始化默认模板
	Run:     rename,
}

func rename(cmd *cobra.Command, args []string) {
	utils.RenameToPNG(utils.ClearPath(path.Join(utils.GetExecDir(), "wallpapers\\")))
}
