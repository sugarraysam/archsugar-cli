package cmd

import (
	"log"
	"strings"

	"github.com/spf13/cobra"
	"github.com/sugarraysam/archsugar-cli/helpers"
	"github.com/sugarraysam/archsugar-cli/scenario"
)

func init() {
	disableCmd.Flags().BoolVarP(&disableAll, "all", "a", disableAll, "disable all master scenarios")
	rootCmd.AddCommand(disableCmd)
}

var (
	disableAll      = false
	disableExamples = []string{
		"  # Disable scenarios foo and bar",
		"  archsugar disable foo bar",
		"  # Disable all scenarios",
		"  archsugar disable --all",
	}
	disableCmd = &cobra.Command{
		Use:     "disable [-a|--all] | [SCENARIO_1 ... SCENARIO_N]",
		Short:   "Disable one or more scenario",
		Example: strings.Join(disableExamples, "\n"),
		Run:     disableMain,
	}
)

func disableMain(cmd *cobra.Command, args []string) {
	// case 1: no flag and no args
	if !disableAll && len(args) == 0 {
		_ = cmd.Help()
		log.Fatalln("Must provide either --all or scenario names")
	}
	// case 2: flag and args
	if disableAll && len(args) > 0 {
		_ = cmd.Help()
		log.Fatalln("Must provide either --all or scenario names")
	}

	// case 3: --all flag
	if disableAll {
		disableAllScenarios()
	}

	// case 4: scenarios to disable provided as args
	if len(args) > 0 {
		disableAllFromArgs(args)
	}
}

func disableAllScenarios() {
	xs, err := scenario.ReadAll(helpers.BaseDir)
	if err != nil {
		log.Fatalln("Could not initialize all scenarios:", err)
	}
	if err := xs.DisableAll(); err != nil {
		log.Fatalln("Could not disable all scenarios:", err)
	}
}

func disableAllFromArgs(args []string) {
	for _, name := range args {
		s, err := scenario.Read(helpers.BaseDir, name)
		if err != nil {
			log.Fatalln("Could not initialize scenario", err)
		}
		if err := s.Disable(); err != nil {
			log.Fatalln("Could not disable a scenario", err)
		}
	}
}
