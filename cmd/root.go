package cmd

import (
	"github.com/spf13/cobra"
)

var (
	rootCmd = &cobra.Command{
		Use:   "WinWallpaperSaver",
		Short: "A tool to save Win10/Win11 wallpapers",
		Run: copyAndRename,
	}
)

func init() {
	rootCmd.AddCommand(copyCmd)
	rootCmd.AddCommand(renameCmd)
	rootCmd.AddCommand(copyAndRenameCmd)
	rootCmd.AddCommand(versionCmd)
	rootCmd.AddCommand(listCmd)
}

func Execute() error {
	return rootCmd.Execute()
}
