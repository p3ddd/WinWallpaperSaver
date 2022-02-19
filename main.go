package main

import (
	"Win10WallpaperSaver/cmd"

	"github.com/spf13/cobra"
)

func main() {
	cobra.CheckErr(cmd.Execute())
}
