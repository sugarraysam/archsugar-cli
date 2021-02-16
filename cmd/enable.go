package cmd

import (
	"log"
	"strings"

	"github.com/spf13/cobra"
	"github.com/sugarraysam/archsugar-cli/helpers"
	"github.com/sugarraysam/archsugar-cli/scenario"
)

func init() {
	enableCmd.Flags().BoolVarP(&enableAll, "all", "a", enableAll, "enable all master scenarios")
	rootCmd.AddCommand(enableCmd)
}

var (
	enableAll      = false
	enableExamples = []string{
		"  # Enable scenario foo and bar",
		"  archsugar enable foo bar",
		"  # Enable all scenarios",
		"  archsugar enable --all",
	}
	enableCmd = &cobra.Command{
		Use:     "enable [-a|--all] | [SCENARIO_1 ... SCENARIO_N]",
		Short:   "Enable one or more scenarios",
		Example: strings.Join(enableExamples, "\n"),
		Run:     enableMain,
	}
)

func enableMain(cmd *cobra.Command, args []string) {
	// case 1: no flag and no args
	if !enableAll && len(args) == 0 {
		_ = cmd.Help()
		log.Fatalln("Must provide either --all or scenario names")
	}
	// case 2: flag and args
	if enableAll && len(args) > 0 {
		_ = cmd.Help()
		log.Fatalln("Must provide either --all or scenario names")
	}

	// case 3: --all flag
	if enableAll {
		enableAllScenarios()
	}

	// case 4: scenarios to enable provided as args
	if len(args) > 0 {
		enableAllFromArgs(args)
	}
}

func enableAllScenarios() {
	xs, err := scenario.ReadAll(helpers.BaseDir)
	if err != nil {
		log.Fatalln("Could not initialize all scenarios:", err)
	}
	if err := xs.EnableAll(); err != nil {
		log.Fatalln("Could not enable all scenarios:", err)
	}
}

func enableAllFromArgs(args []string) {
	for _, name := range args {
		s, err := scenario.Read(helpers.BaseDir, name)
		if err != nil {
			log.Fatalln("Could not initialize scenario:", err)
		}
		if err := s.Enable(); err != nil {
			log.Fatalln("Could not enable a scenario:", err)
		}
	}
}
