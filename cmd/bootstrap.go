package cmd

import (
	"os"
	"strconv"
	"strings"

	log "github.com/sirupsen/logrus"

	"github.com/spf13/cobra"
	"github.com/sugarraysam/archsugar-cli/playbook"
)

func init() {
	bootstrapCmd.Flags().StringVar(&bootstrapDisk, "disk", bootstrapDisk, "root disk to use for partitioning")
	bootstrapCmd.Flags().StringVar(&bootstrapLuksPasswd, "luksPasswd", bootstrapLuksPasswd, "LUKS Password used to encrypt disk")
	bootstrapCmd.Flags().StringVar(&bootstrapRootPasswd, "rootPasswd", bootstrapRootPasswd, "password of root user")
	bootstrapCmd.Flags().StringVar(&bootstrapUserPasswd, "userPasswd", bootstrapUserPasswd, "password of unpriviledged user")
	bootstrapCmd.Flags().BoolVar(&bootstrapVagrant, "vagrant", bootstrapVagrant, "Bootstrap VM as a vagrant box")
	rootCmd.AddCommand(bootstrapCmd)
}

var (
	// CLI flags
	bootstrapDisk       = "/dev/sda"
	bootstrapLuksPasswd = "luks"
	bootstrapRootPasswd = "root"
	bootstrapUserPasswd = "sugar"
	bootstrapVagrant    = false

	// Examples
	bootstrapExamples = []string{
		"  # Bootstrap a VM as a vagrant box",
		"  # luksPasswd, rootPasswd, userPasswd and username will all be set to 'vagrant'",
		"  archsugar bootstrap --disk /dev/sda --vagrant",
		"  # Bootstrap a bare-metal computer",
		"  archsugar bootstrap --disk /dev/sda --luksPasswd luks --rootPasswd root --userPasswd user",
	}
	bootstrapCmd = &cobra.Command{
		Use:     "bootstrap --disk <disk> [ --luksPasswd <luks> --rootPasswd <root> --userPasswd <user> | --vagrant ]",
		Short:   "Run bootstrap and chroot stage to provision a bare-metal machine or a VM",
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
	os.Setenv("SUGAR_VAGRANT", strconv.FormatBool(bootstrapVagrant))
	if bootstrapVagrant {
		os.Setenv("SUGAR_LUKS_PASSWD", "vagrant")
		os.Setenv("SUGAR_ROOT_PASSWD", "vagrant")
		os.Setenv("SUGAR_USER_PASSWD", "vagrant")
		os.Setenv("SUGAR_USER", "vagrant")
	} else {
		os.Setenv("SUGAR_LUKS_PASSWD", bootstrapLuksPasswd)
		os.Setenv("SUGAR_ROOT_PASSWD", bootstrapRootPasswd)
		os.Setenv("SUGAR_USER_PASSWD", bootstrapUserPasswd)
	}
}
