package dotfiles_test

import (
	"os"
	"testing"

	"github.com/go-git/go-git/v5"
	"github.com/stretchr/testify/require"
	"github.com/sugarraysam/archsugar-cli/dotfiles"
	"github.com/sugarraysam/archsugar-cli/helpers"
)

func TestNewDefaultRepo(t *testing.T) {
	repo := dotfiles.NewDefaultRepo()
	require.Equal(t, repo.Dest, dotfiles.DefaultDest)
	require.Equal(t, repo.URL, dotfiles.DefaultURL)
	require.Equal(t, repo.Branch, dotfiles.DefaultBranch)
}

func TestArchsugarRepoCloneMasterAndDev(t *testing.T) {
	// cleanup
	tmpDir := helpers.TmpDir()
	defer func() {
		_ = os.RemoveAll(tmpDir)
	}()

	// clone master
	master, err := dotfiles.NewRepo(tmpDir, dotfiles.DefaultURL, "master")
	require.Nil(t, err)
	require.Nil(t, master.Clone())
	require.True(t, master.Exists())
	validateBranch(t, master.Dest, "master")

	// clone dev
	dev, err := dotfiles.NewRepo(tmpDir, dotfiles.DefaultURL, "dev")
	require.Nil(t, err)

	// repository already exists
	require.True(t, dev.Exists())
	require.NotNil(t, dev.Clone())

	// force clean before cloning
	require.Nil(t, master.Rm())
	require.Nil(t, dev.Clone())
	require.True(t, dev.Exists())
	validateBranch(t, dev.Dest, "dev")
}

func validateBranch(t *testing.T, path string, branch string) {
	r, err := git.PlainOpen(path)
	require.Nil(t, err)

	head, err := r.Head()
	require.Nil(t, err)

	require.Contains(t, head.Name(), branch)
	require.True(t, head.Name().IsBranch())
}
