package dotfiles

import (
	"fmt"
	"net/url"
	"os"

	"github.com/go-git/go-git/v5"
	"github.com/go-git/go-git/v5/plumbing"
	log "github.com/sirupsen/logrus"
	"github.com/sugarraysam/archsugar-cli/helpers"
)

const (
	DefaultURL    = "https://github.com/sugarraysam/archsugar"
	DefaultBranch = "master"
)

type Repo struct {
	URL    string
	Dst    string
	Branch string
}

func NewRepo(rawURL, branch string) (*Repo, error) {
	if rawURL == "" {
		rawURL = DefaultURL
	}
	u, err := url.Parse(rawURL)
	if err != nil {
		return nil, err
	}
	if branch == "" {
		branch = DefaultBranch
	}
	return &Repo{
		URL:    u.String(),
		Dst:    helpers.BaseDir,
		Branch: branch,
	}, nil
}

func (r *Repo) Clone() error {
	if r.Exists() {
		return fmt.Errorf("could not clone, repository already exists at: %s", r.Dst)
	}
	_, err := git.PlainClone(r.Dst, false, &git.CloneOptions{
		URL:           r.URL,
		ReferenceName: plumbing.NewBranchReferenceName(r.Branch),
		Depth:         1,
		Progress:      os.Stdout,
	})
	if err != nil {
		log.Errorf("Could not clone dotfiles repository, got %v", err)
	}
	return err
}

func (r *Repo) Exists() bool {
	_, err := os.Stat(r.Dst)
	return err == nil
}

func (r *Repo) Rm() error {
	return os.RemoveAll(helpers.BaseDir)
}
