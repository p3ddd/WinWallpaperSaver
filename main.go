package main

import (
	"WinWallpaperSaver/cmd"

	"github.com/spf13/cobra"
)

func main() {
	cobra.CheckErr(cmd.Execute())
}
