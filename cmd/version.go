package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

const version = "0.0.1_alpha"

var versionCmd = &cobra.Command{
	Use:     "version",
	Aliases: []string{"v"},
	Short:   "print version",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println(version)
	},
}
