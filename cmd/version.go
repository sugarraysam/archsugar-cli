package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
	"github.com/sugarraysam/archsugar-cli/version"
)

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Print version, build time and git commit.",
	Args:  cobra.NoArgs,
	Run:   versionMain,
}

func versionMain(cmd *cobra.Command, args []string) {
	format := "Version: %s\nDate: %s\nCommit: %s\n"
	fmt.Printf(format, version.Version, version.Date, version.Commit)
}

func init() {
	rootCmd.AddCommand(versionCmd)
}
