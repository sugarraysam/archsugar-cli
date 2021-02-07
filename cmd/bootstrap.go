package cmd

import (
	"os"
	"strings"

	log "github.com/sirupsen/logrus"

	"github.com/spf13/cobra"
	"github.com/sugarraysam/archsugar-cli/playbook"
)

func init() {
	bootstrapCmd.Flags().StringVar(&bootstrapDisk, "disk", bootstrapDisk, "root disk to use for partitioning")
	bootstrapCmd.Flags().StringVar(&bootstrapLuksPasswd, "luks", bootstrapLuksPasswd, "luksPasswd used to encrypt disk")
	bootstrapCmd.Flags().StringVar(&bootstrapRootPasswd, "root", bootstrapRootPasswd, "password of root user")
	bootstrapCmd.Flags().StringVar(&bootstrapUserPasswd, "user", bootstrapUserPasswd, "password of unpriviledged user")
	rootCmd.AddCommand(bootstrapCmd)
}

var (
	// CLI flags
	bootstrapDisk       = "/dev/sda"
	bootstrapLuksPasswd = "luks"
	bootstrapRootPasswd = "root"
	bootstrapUserPasswd = "sugar"

	// Examples
	bootstrapExamples = []string{
		"  # Run a bootstrap playbook with default args",
		"  archsugar bootstrap",
		"  # Run a bootstrap playbook by specifying all parameters",
		"  archsugar bootstrap --disk /dev/sda --luksPasswd luks --rootPasswd root --userPasswd user",
	}
	bootstrapCmd = &cobra.Command{
		Use:     "bootstrap [--disk <disk> --luksPasswd <luks> --rootPasswd <root> --userPasswd <user>]",
		Short:   "Run bootstrap and chroot stage to provision new bare machine or VM",
		Example: strings.Join(bootstrapExamples, "\n"),
		Run:     bootstrapMain,
		Args:    cobra.NoArgs,
	}
)

func bootstrapMain(cmd *cobra.Command, args []string) {
	// set args as env vars, will be added to Cmd Env
	setBoostrapExtraEnv()

	// Run bootstrap playbook
	bootstrap := playbook.NewBootstrapPlaybook()
	if err := bootstrap.Run(); err != nil {
		log.Fatalln("Error running bootstrap playbook:", err)
	}
	// Run chroot playbook
	chroot := playbook.NewChrootPlaybook()
	if err := chroot.Run(); err != nil {
		log.Fatalln("Error running chroot playbook:", err)
	}
}

func setBoostrapExtraEnv() {
	os.Setenv("SUGAR_DISK", bootstrapDisk)
	os.Setenv("SUGAR_LUKS_PASSWD", bootstrapLuksPasswd)
	os.Setenv("SUGAR_ROOT_PASSWD", bootstrapRootPasswd)
	os.Setenv("SUGAR_USER_PASSWD", bootstrapUserPasswd)
}
