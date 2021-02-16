package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// Set by ldflags
var version = "unset"
var commit = "unset"
var date = "unset"

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Print version, build time and git commit.",
	Args:  cobra.NoArgs,
	Run:   versionMain,
}

func versionMain(cmd *cobra.Command, args []string) {
	format := "Version: %s\nDate: %s\nCommit: %s\n"
	fmt.Printf(format, version, date, commit)
}

func init() {
	rootCmd.AddCommand(versionCmd)
}
