package dotfiles_test

import (
	"fmt"
	"math/rand"
	"os"
	"path"
	"testing"
	"time"

	"github.com/go-git/go-git/v5"
	"github.com/stretchr/testify/require"
	"github.com/sugarraysam/archsugar-cli/dotfiles"
	"github.com/sugarraysam/archsugar-cli/helpers"
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

func TestMain(m *testing.M) {
	// Setup - create new directory structure under /tmp
	tmpDir := path.Join(os.TempDir(), fmt.Sprintf("archsugar-%d", helpers.GetRandomDigit()))
	helpers.BaseDir = tmpDir

	// Run
	rc := m.Run()

	// Teardown
	_ = os.RemoveAll(tmpDir)
	os.Exit(rc)
}

func TestArchsugarRepoCloneMasterAndDev(t *testing.T) {
	// clone master
	master, err := dotfiles.NewRepo(dotfiles.DefaultURL, dotfiles.DefaultBranch)
	require.Nil(t, err)
	require.Nil(t, master.Clone())
	require.True(t, master.Exists())
	validateBranch(t, master.Dst, "master")

	// clone dev
	dev, err := dotfiles.NewRepo(dotfiles.DefaultURL, "dev")
	require.Nil(t, err)

	// repository already exists
	require.True(t, dev.Exists())
	require.NotNil(t, dev.Clone())

	// force clean before cloning
	require.Nil(t, master.Rm())
	require.Nil(t, dev.Clone())
	require.True(t, dev.Exists())
	validateBranch(t, dev.Dst, "dev")
}

func validateBranch(t *testing.T, path string, branch string) {
	r, err := git.PlainOpen(path)
	require.Nil(t, err)

	head, err := r.Head()
	require.Nil(t, err)

	require.Contains(t, head.Name(), branch)
	require.True(t, head.Name().IsBranch())
}
