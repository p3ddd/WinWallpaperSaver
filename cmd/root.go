package cmd

import (
	"github.com/spf13/cobra"
)

var (
	// cfgFile string
	rootCmd = &cobra.Command{
		Use:   "Win10WallpaperSaver",
		Short: "A tool to save Win10 wallpapers",
	}
)

func init() {
	rootCmd.AddCommand(copyCmd)
	rootCmd.AddCommand(renameCmd)
	rootCmd.AddCommand(copyAndRenameCmd)
	rootCmd.AddCommand(versionCmd)
}

func Execute() error {
	return rootCmd.Execute()
}
