package cmd

import (
	"os"

	"strings"

	log "github.com/sirupsen/logrus"

	"github.com/spf13/cobra"
	"github.com/sugarraysam/archsugar-cli/playbook"
)

func init() {
	rootCmd.AddCommand(runCmd)
}

var (
	runExamples = []string{
		"  # Run all enabled scenarios",
		"  archsugar run",
		"  # Run specified scenarios (they don't have to be enabled)",
		"  archsugar run foo bar baz",
	}
	runCmd = &cobra.Command{
		Use:     "run [SCENARIO_1 ... SCENARIO_N]",
		Short:   "Run scenarios",
		Example: strings.Join(runExamples, "\n"),
		Run:     runMain,
	}
)

func runMain(cmd *cobra.Command, args []string) {
	setRunExtraEnv(args)
	p := playbook.NewMasterPlaybook()
	if err := p.Run(); err != nil {
		log.Fatalln("Error running master scenario:", err)
	}
}

// export SUGAR_SCENARIOS to be picked up by ansible
// trim and comma separate scenarios
func setRunExtraEnv(scenarios []string) {
	if len(scenarios) == 0 {
		os.Setenv("SUGAR_SCENARIOS", "all")
		return
	}
	var v []string
	for _, s := range scenarios {
		v = append(v, strings.Trim(s, " ,.\n\t"))
	}
	os.Setenv("SUGAR_SCENARIOS", strings.Join(v, ","))
}
