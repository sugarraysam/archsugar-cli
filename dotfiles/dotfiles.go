package dotfiles

import (
	"errors"
	"fmt"
	"net/url"
	"os"

	"github.com/go-git/go-git/v5"
	"github.com/go-git/go-git/v5/plumbing"
	log "github.com/sirupsen/logrus"
	"github.com/sugarraysam/archsugar-cli/helpers"
)

var (
	DefaultURL    = "https://github.com/sugarraysam/archsugar"
	DefaultDest   = helpers.BaseDir
	DefaultBranch = "master"
)

type Repo struct {
	Dest   string
	URL    string
	Branch string
}

func NewDefaultRepo() *Repo {
	return &Repo{
		Dest:   DefaultDest,
		URL:    DefaultURL,
		Branch: DefaultBranch,
	}
}

func NewRepo(dest, rawURL, branch string) (*Repo, error) {
	if rawURL == "" || dest == "" || branch == "" {
		return nil, errors.New("some or multiple invalid empty arguments")
	}
	u, err := url.Parse(rawURL)
	if err != nil {
		return nil, err
	}
	return &Repo{
		Dest:   dest,
		URL:    u.String(),
		Branch: branch,
	}, nil
}

func (r *Repo) Clone() error {
	if r.Exists() {
		return fmt.Errorf("could not clone, repository already exists at: %s", r.Dest)
	}
	_, err := git.PlainClone(r.Dest, false, &git.CloneOptions{
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
	_, err := os.Stat(r.Dest)
	return err == nil
}

func (r *Repo) Rm() error {
	return os.RemoveAll(r.Dest)
}
