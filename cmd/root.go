package cmd

import (
	"github.com/sirupsen/logrus"
	log "github.com/sirupsen/logrus"

	"github.com/spf13/cobra"
)

func init() {
	rootCmd.PersistentFlags().BoolVar(&rootVerbose, "v", rootVerbose, "verbose output")
	rootCmd.PersistentFlags().BoolVar(&rootExtraVerbose, "vvv", rootExtraVerbose, "extra verbose output")
	cobra.OnInitialize(setLogFormatter, setLogLevel)
}

var (
	rootVerbose      = false
	rootExtraVerbose = false
	rootCmd          = &cobra.Command{
		Use:   "archsugar",
		Short: "Ansible powered CLI to bootstrap and maintain a high-end archlinux workstation.",
		Run:   rootMain,
	}
)

func rootMain(cmd *cobra.Command, args []string) {
	_ = cmd.Help()
}

// Execute - called by main.main()
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		log.Fatalln(err)
	}
}

func setLogFormatter() {
	formatter := &logrus.TextFormatter{
		TimestampFormat:        "02-01-2006 15:04:05",
		FullTimestamp:          true,
		DisableLevelTruncation: true,
	}
	log.SetFormatter(formatter)
}

func setLogLevel() {
	log.SetLevel(log.InfoLevel)
	if rootVerbose {
		log.SetLevel(log.DebugLevel)
	}
	if rootExtraVerbose {
		log.SetLevel(log.TraceLevel)
	}
}
