package cmd

import (
	"log"
	"strings"

	"github.com/spf13/cobra"
	"github.com/sugarraysam/archsugar-cli/scenario"
)

func init() {
	rootCmd.AddCommand(rmCmd)
}

var (
	rmExamples = []string{
		"  # Remove scenario foo",
		"  archsugar rm foo",
	}
	rmCmd = &cobra.Command{
		Use:     "rm SCENARIO",
		Example: strings.Join(rmExamples, "\n"),
		Short:   "Remove a scenario",
		Args:    cobra.ExactArgs(1),
		Run:     rmMain,
	}
)

func rmMain(cmd *cobra.Command, args []string) {
	// only allow removing a single scenario at a time
	name := args[0]
	s, err := scenario.NewCreatedScenario(name)
	if err != nil {
		log.Fatalln("Could not initialize scenario:", err)
	}
	if err := s.Rm(); err != nil {
		log.Fatalln("Could not remove scenario:", err)
	}
}
