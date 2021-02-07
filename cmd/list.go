package cmd

import (
	"log"
	"os"
	"strings"

	"github.com/spf13/cobra"
	"github.com/sugarraysam/archsugar-cli/scenario"
)

func init() {
	rootCmd.AddCommand(listCmd)
}

var (
	listExamples = []string{
		"  # List all scenarios",
		"  archsugar list",
	}
	listCmd = &cobra.Command{
		Use:     "list",
		Short:   "List all scenarios",
		Example: strings.Join(listExamples, "\n"),
		Args:    cobra.NoArgs,
		Run:     listMain,
	}
)

func listMain(cmd *cobra.Command, args []string) {
	xs, err := scenario.GetAllScenarios()
	if err != nil {
		log.Fatalln("Could not get all scenarios:", err)
	}
	if err := xs.List(os.Stdout); err != nil {
		log.Fatalln("Could not list all scenarios:", err)
	}
}
