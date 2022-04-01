package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

const version = "0.0.2"

var versionCmd = &cobra.Command{
	Use:     "version",
	Aliases: []string{"v"},
	Short:   "print version",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println(version)
	},
}
