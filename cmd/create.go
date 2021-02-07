package cmd

import (
	"log"
	"strings"

	"github.com/spf13/cobra"
	"github.com/sugarraysam/archsugar-cli/scenario"
)

func init() {
	createCmd.Flags().StringVarP(&createDesc, "desc", "d", createDesc, "description of the scenario")
	rootCmd.AddCommand(createCmd)
}

var (
	createDesc     = ""
	createExamples = []string{
		"  # Create scenario magic with a description",
		"  archsugar create magic --desc 'double rainbow..all the way!'",
		"  # Create scenario trouble no description",
		"  archsugar create trouble",
	}
	createCmd = &cobra.Command{
		Use:     "create [-d DESCRIPTION] SCENARIO",
		Short:   "Create a scenario using predefined templates",
		Example: strings.Join(createExamples, "\n"),
		Args:    cobra.ExactArgs(1),
		Run:     createMain,
	}
)

func createMain(cmd *cobra.Command, args []string) {
	name := args[0]
	s, err := scenario.NewUncreatedScenario(name, createDesc)
	if err != nil {
		log.Fatalln("Could not initialize scneario:", err)
	}
	if err := s.Create(); err != nil {
		log.Fatalln("Could not create scenario:", err)
	}
}
