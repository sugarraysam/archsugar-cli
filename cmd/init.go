package cmd

import (
	"strings"

	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"github.com/sugarraysam/archsugar-cli/dotfiles"
)

func init() {
	initCmd.Flags().BoolVarP(&initForce, "force", "f", initForce, "overwrite existing dotfiles")
	initCmd.Flags().StringVarP(&initURL, "url", "u", initURL, "URL of dotfiles git repository")
	initCmd.Flags().StringVarP(&initBranch, "branch", "b", initBranch, "Branch of dotfiles repository to clone")
	rootCmd.AddCommand(initCmd)
}

var (
	// CLI flags
	initForce  = false
	initURL    = dotfiles.DefaultURL
	initBranch = dotfiles.DefaultBranch

	// Examples
	initExamples = []string{
		"  # Init archsugar using default dotfiles repo",
		"  archsugar init",
		"  # Run archsugar using dev branch of dotfiles repo",
		"  archsugar init -b dev",
		"  # Overwrite existing dotfiles with a custom repository",
		"  archsugar init -f true -b <branch> -u 'https://github.com/username/dotfiles'",
	}
	initCmd = &cobra.Command{
		Use:     "init [--url <url> --branch <branch> --force]",
		Short:   "Clone dotfiles repository to ~/.archsugar",
		Example: strings.Join(initExamples, "\n"),
		Run:     initMain,
		Args:    cobra.NoArgs,
	}
)

func initMain(cmd *cobra.Command, args []string) {
	repo, err := dotfiles.NewRepo(initURL, initBranch)
	if err != nil {
		log.Fatalln("Could not initialize dotfiles repo:", err)
	}
	if repo.Exists() {
		if !initForce {
			log.Fatalf("A repository already exists at %s, add --force to overwrite it.", repo.Dst)
		}
		if err = repo.Rm(); err != nil {
			log.Fatalln("Could not overwrite existing dotfiles repo:", err)
		}
	} else {
		if err = repo.Clone(); err != nil {
			log.Fatalln("Could not clone dotfiles repo:", err)
		}
	}
}
