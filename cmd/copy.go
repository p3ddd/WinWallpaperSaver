package cmd

import (
	"Win10WallpaperSaver/utils"
	"fmt"
	"path"

	"github.com/spf13/cobra"
)

var copyCmd = &cobra.Command{
	Use:     "copy",
	Aliases: []string{""},
	Short:   "copy", // 初始化默认模板
	Run:     copy,
}

func copy(cmd *cobra.Command, args []string) {
	from := utils.GetAssetsPath()
	to := utils.ClearPath(path.Join(utils.GetExecDir(), "wallpapers\\"))

	fmt.Println("AssetsDir: ", from)
	fmt.Println("ExecDir: ", to)
	fmt.Println("dstDir: ", to)

	if err := utils.Copy(from, to); err != nil {
		fmt.Println(err)
	}
}
