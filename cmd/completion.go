package cmd

import (
	"os"
	"strings"

	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(completionCmd)
}

var (
	completionExample = []string{
		"  # Load completion zsh current shell",
		"  source <(archsugar completion)",
		"  # Configure zsh shell to load completions for each session",
		"  archsugar completion > /usr/share/zsh/site-functions/_archsugar",
	}
	completionCmd = &cobra.Command{
		Use:     "completion",
		Short:   "Generates zsh completion script",
		Example: strings.Join(completionExample, "\n"),
		Run:     completionMain,
		Args:    cobra.NoArgs,
	}
)

func completionMain(cmd *cobra.Command, args []string) {
	_ = rootCmd.GenZshCompletion(os.Stdout)
}
