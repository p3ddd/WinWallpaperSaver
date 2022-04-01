package cmd

import (
	"WinWallpaperSaver/utils"
	"fmt"

	"github.com/spf13/cobra"
)

var copyCmd = &cobra.Command{
	Use:     "copy",
	Aliases: []string{"c"},
	Short:   "copy",
	Run:     copy,
}

func copy(cmd *cobra.Command, args []string) {
	from := utils.GetAssetsPath()
	to := utils.GetDstDir()

	fmt.Println("AssetsDir: ", from)
	fmt.Println("ExecDir: ", to)
	fmt.Println("dstDir: ", to)

	if err := utils.Copy(from, to); err != nil {
		fmt.Println(err)
	}
}
